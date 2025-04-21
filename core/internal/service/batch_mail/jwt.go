package batch_mail

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang-jwt/jwt/v5"
	"sync"
	"time"
)

const (
	JWT_SECRET_KEY = "unsubscribe_jwt_secret"
)

// UnsubscribeClaims 定义退订JWT的声明结构
type UnsubscribeClaims struct {
	Email      string `json:"email"`
	TemplateId int    `json:"template_id"`
	TaskId     int    `json:"task_id"`
	jwt.RegisteredClaims
}

// jwtConfig JWT配置结构
type jwtConfig struct {
	secret string
	expiry time.Duration
}

var (
	config *jwtConfig
	once   sync.Once
)

// getConfig 获取JWT配置单例
func getConfig() *jwtConfig {
	once.Do(func() {
		config = loadConfig()
	})
	return config
}

// loadConfig 加载JWT配置
func loadConfig() *jwtConfig {
	ctx := gctx.New()
	return &jwtConfig{
		secret: getOrGenerateSecret(ctx),
		expiry: 30 * 24 * time.Hour, // 30 days
	}
}

// getOrGenerateSecret 获取或生成密钥
func getOrGenerateSecret(ctx context.Context) string {
	// 1. 从数据库获取密钥
	val, err := g.DB().Model("bm_options").
		Where("name", JWT_SECRET_KEY).
		Value("value")

	if err == nil && val != nil && val.String() != "" {
		return val.String()
	}

	// 2. 生成新的密钥
	newSecret := generateSecret(ctx)

	// 3. 保存到数据库
	_, err = g.DB().Model("bm_options").
		Data(g.Map{
			"name":  JWT_SECRET_KEY,
			"value": newSecret,
		}).
		Where("name", JWT_SECRET_KEY).
		WhereNull("value").
		Insert()

	if err != nil {
		g.Log().Warning(ctx, "Failed to persist JWT secret to database:", err)

		val, err := g.DB().Model("bm_options").
			Where("name", JWT_SECRET_KEY).
			Value("value")

		if err == nil && val != nil && val.String() != "" {
			return val.String()
		}
	} else {
		g.Log().Info(ctx, "Generated and persisted new JWT unsubscribe secret")
	}

	return newSecret
}

// generateSecret 生成密钥
func generateSecret(ctx context.Context) string {
	// 使用应用程序的固定特征作为密钥组件
	components := []string{
		"BILLION_MAIL_UNSUBSCRIBE",                            // 应用标识
		g.Cfg().MustGet(ctx, "server.address").String(),       // 服务器地址
		g.Cfg().MustGet(ctx, "server.sessionIdName").String(), // session名称
	}

	// 如果有数据库配置，添加数据库名作为组件
	if dbName := g.Cfg().MustGet(ctx, "database.name").String(); dbName != "" {
		components = append(components, dbName)
	}

	// 添加服务器特定信息
	if serverAgent := g.Cfg().MustGet(ctx, "server.serverAgent").String(); serverAgent != "" {
		components = append(components, serverAgent)
	}

	// 使用 SHA256 哈希所有组件
	h := sha256.New()
	for _, component := range components {
		h.Write([]byte(component))
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

// GenerateUnsubscribeJWT 生成退订JWT
func GenerateUnsubscribeJWT(email string, templateId, taskId int) (string, error) {
	cfg := getConfig()

	claims := UnsubscribeClaims{
		Email:      email,
		TemplateId: templateId,
		TaskId:     taskId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.secret))
}

// ParseUnsubscribeJWT 解析退订JWT
func ParseUnsubscribeJWT(tokenString string) (*UnsubscribeClaims, error) {
	cfg := getConfig()

	token, err := jwt.ParseWithClaims(tokenString, &UnsubscribeClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UnsubscribeClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
