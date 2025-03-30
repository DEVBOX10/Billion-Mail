package rbac

import (
	"billionmail-core/internal/service/public"
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/util/guid"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

// JWTCustomClaims defines the custom claims for JWT
type JWTCustomClaims struct {
	AccountId int64    `json:"accountId"`
	Username  string   `json:"username"`
	Roles     []string `json:"roles"`
	jwt.RegisteredClaims
}

// JWTService provides methods for JWT operations
type JWTService struct {
	Secret        string        // Secret key for signing JWT
	AccessExpiry  time.Duration // Duration for access token expiry
	RefreshExpiry time.Duration // Duration for refresh token expiry
}

// newJWTService creates a new JWTService instance
func newJWTService() *JWTService {
	return &JWTService{
		Secret:        g.Cfg().MustGet(context.Background(), "jwt.secret", "billion-mail-jwt-secret").String(),
		AccessExpiry:  time.Duration(g.Cfg().MustGet(context.Background(), "jwt.accessExpiry", 3600).Int()) * time.Second,
		RefreshExpiry: time.Duration(g.Cfg().MustGet(context.Background(), "jwt.refreshExpiry", 86400*7).Int()) * time.Second,
	}
}

// GenerateToken generates a new JWT token
func (s *JWTService) GenerateToken(accountId int64, username string, roles []string) (string, int64, error) {
	expiryTime := time.Now().Add(s.AccessExpiry)
	claims := &JWTCustomClaims{
		AccountId: accountId,
		Username:  username,
		Roles:     roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "billion-mail",
			Subject:   "access_token",
			ID:        guid.S(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", 0, err
	}

	return signedToken, expiryTime.Unix(), nil
}

// GenerateRefreshToken generates a refresh token
func (s *JWTService) GenerateRefreshToken(accountId int64, username string) (string, error) {
	expiryTime := time.Now().Add(s.RefreshExpiry)
	claims := &JWTCustomClaims{
		AccountId: accountId,
		Username:  username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "billion-mail",
			Subject:   "refresh_token",
			ID:        guid.S(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ParseToken parses and validates a JWT token
func (s *JWTService) ParseToken(tokenString string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, gerror.New("invalid token")
}

// JWTAuthMiddleware provides JWT authentication middleware
func (s *JWTService) JWTAuthMiddleware(r *ghttp.Request) {
	// Skip authentication for login and refresh token endpoints
	if r.URL.Path == "/api/v1/login" || r.URL.Path == "/api/v1/refresh-token" {
		r.Middleware.Next()
		return
	}

	// Get token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteJson(public.CodeMap[401])
		r.Exit()
		return
	}

	// Parse token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := s.ParseToken(tokenString)
	if err != nil {
		r.Response.WriteJson(public.CodeMap[401])
		r.Exit()
		return
	}

	// Set account info in context
	r.SetCtxVar("accountId", claims.AccountId)
	r.SetCtxVar("username", claims.Username)
	r.SetCtxVar("roles", claims.Roles)

	r.Middleware.Next()
}
