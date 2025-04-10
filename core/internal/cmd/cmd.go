package cmd

import (
	"billionmail-core/internal/consts"
	"billionmail-core/internal/controller/dockerapi"
	"billionmail-core/internal/controller/domains"
	"billionmail-core/internal/controller/mail_boxes"
	"billionmail-core/internal/controller/overview"
	"billionmail-core/internal/controller/rbac"
	"billionmail-core/internal/service/database_initialization"
	docker "billionmail-core/internal/service/dockerapi"
	"billionmail-core/internal/service/middlewares"
	"billionmail-core/internal/service/phpfpm"
	"billionmail-core/internal/service/public"
	rbac2 "billionmail-core/internal/service/rbac"
	"billionmail-core/internal/service/redis_initialization"
	"billionmail-core/internal/service/timers"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"
)

var (
	Main = gcmd.Command{
		Name:  consts.DEFAULT_SERVER_NAME,
		Usage: consts.DEFAULT_SERVER_NAME,
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// Init Database
			err = database_initialization.InitDatabase()

			if err != nil {
				g.Log().Error(ctx, "initialize databases failed ", err)
				return err
			}

			// Init Redis
			err = redis_initialization.InitRedis()

			if err != nil {
				g.Log().Error(ctx, "initialize redis failed ", err)
				return err
			}

			// Start timers
			err = timers.Start(ctx)

			if err != nil {
				g.Log().Error(ctx, "start timers failed ", err)
				return err
			}

			// Connect to Docker
			dk, err := docker.NewDockerAPI()

			if err != nil {
				g.Log().Error(ctx, "failed to connect to docker-api ", err)
				return err
			}

			defer dk.Close()

			// Create a new server instance
			s := g.Server(consts.DEFAULT_SERVER_NAME)

			s.Group("/api", func(group *ghttp.RouterGroup) {
				// Add CORS middleware
				group.Middleware(ghttp.MiddlewareCORS)

				// Add docker client middleware
				group.Middleware(func(r *ghttp.Request) {
					r.SetCtxVar(consts.DEFAULT_DOCKER_CLIENT_CTX_KEY, dk)
					r.Middleware.Next()
				})

				// Add JWT middleware
				group.Middleware(rbac2.JWT().JWTAuthMiddleware)

				// Add RBAC middleware
				// group.Middleware(middlewares.NewRBACMiddleware().PermissionCheck)

				// group.Middleware(ghttp.MiddlewareHandlerResponse)

				// Add response
				group.Middleware(middlewares.HandleApiResponse)

				group.Bind(
					rbac.NewV1(),
					domains.NewV1(),
					mail_boxes.NewV1(),
					overview.NewV1(),
					dockerapi.NewV1(),
				)
			})

			// Add PHP-FPM middleware
			s.BindMiddleware("/roundcube/*any", func(r *ghttp.Request) {
				if r.Method == "POST" && strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
					// Get and store the request body
					r.GetBody()
					r.Middleware.Next()
					return
				}
				r.Middleware.Next()
			})

			// Binding PHP-FPM handler
			s.BindHandler("/roundcube/*any", phpfpm.PHPFpmHandlerFactory(phpfpm.PHPFpmHandlerConfig{
				Network: "unix",
				Addr:    consts.PHP_FPM_SOCK_PATH,
				Root:    consts.ROUNDCUBE_ROOT_PATH_IN_CONTAINER,
				Static:  consts.ROUNDCUBE_ROOT_PATH,
			}))

			// Proxy 60880 port for ACME challenge
			s.BindHandler("/.well-known/acme-challenge/*any", func(r *ghttp.Request) {
				proxy := httputil.NewSingleHostReverseProxy(&url.URL{
					Scheme: "http",
					Host:   "127.0.0.1:60880",
				})

				proxy.ServeHTTP(r.Response.BufferWriter, r.Request)
			})

			// Add static file handler
			s.BindHandler("/*", func(r *ghttp.Request) {
				r.Response.ServeFile("public/dist/index.html")
			})

			// Generate self-signed certificate if not exists
			public.SelfSignedCert().Generate()

			// Enable HTTPS
			s.EnableHTTPS(public.AbsPath(filepath.Join(consts.SSL_PATH, "cert.pem")), public.AbsPath(filepath.Join(consts.SSL_PATH, "key.pem")))
			s.SetHTTPSPort(g.Cfg().MustGet(ctx, "server.httpsPort", 443).Int())

			s.Run()

			return nil
		},
	}
)
