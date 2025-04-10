package phpfpm

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"net"
	"path/filepath"
	"strings"
	"time"

	"github.com/tomasen/fcgi_client"
)

// PHPFpmHandlerConfig holds the configuration for the PHP-FPM handler.
type PHPFpmHandlerConfig struct {
	Network string // Network type (tcp, unix)
	Addr    string // Address of the FastCGI server
	Root    string // Document root for PHP files
	Static  string // Serve static files directly
}

// PHPFpmHandlerFactory creates a handler function for PHP-FPM.
func PHPFpmHandlerFactory(config PHPFpmHandlerConfig) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		// Get the requested file path
		filePath := "/" + r.Get("any").String()

		// Check if the file path is empty
		if filePath == "/" {
			filePath = "/index.php"
		}

		// Serve static files directly
		if !strings.HasSuffix(filePath, ".php") {
			r.Response.ServeFile(filepath.Join(config.Static, filePath))
			return
		}

		https := "off"
		if r.GetSchema() == "https" {
			https = "on"
		}

		// Get the remote address and port
		remoteAddr, port, _ := net.SplitHostPort(r.RemoteAddr)

		// Create environment variables for FastCGI
		env := map[string]string{
			"SCRIPT_FILENAME":   filepath.Join(config.Root, filePath),
			"REQUEST_METHOD":    r.Method,
			"SCRIPT_NAME":       filePath,
			"REQUEST_URI":       r.RequestURI,
			"QUERY_STRING":      r.URL.RawQuery,
			"CONTENT_TYPE":      r.Header.Get("Content-Type"),
			"CONTENT_LENGTH":    r.Header.Get("Content-Length"),
			"REMOTE_ADDR":       remoteAddr,
			"SERVER_NAME":       r.Host,
			"SERVER_PORT":       port,
			"SERVER_PROTOCOL":   r.Proto,
			"HTTPS":             https,
			"REQUEST_TIME":      gconv.String(time.Now().Unix()),
			"PATH_INFO":         filePath,
			"DOCUMENT_ROOT":     config.Root,
			"GATEWAY_INTERFACE": "CGI/1.1",
			"SERVER_SOFTWARE":   "gf",
		}

		// Add all HTTP headers to environment with HTTP_ prefix
		for headerName, headerValues := range r.Header {
			if len(headerValues) > 0 {
				headerName = strings.ReplaceAll(strings.ToUpper(headerName), "-", "_")
				if headerName != "CONTENT_TYPE" && headerName != "CONTENT_LENGTH" {
					env["HTTP_"+headerName] = headerValues[0]
				}
			}
		}

		// Connect to FastCGI server
		fc, err := fcgiclient.Dial(config.Network, config.Addr)
		if err != nil {
			g.Log().Error(context.Background(), "FastCGI connection failed:", err)
			return
		}
		defer fc.Close()

		// Send request to FastCGI server
		resp, err := fc.Request(env, r.Body)
		if err != nil {
			g.Log().Error(context.Background(), "FastCGI request failed:", err)
			return
		}
		defer resp.Body.Close()

		// Set the response status code
		r.Response.WriteHeader(resp.StatusCode)

		// Copy headers from PHP response to our response
		for key, values := range resp.Header {
			for _, value := range values {
				r.Response.Header().Add(key, value)
			}
		}

		// Copy response body
		_, err = io.Copy(r.Response.BufferWriter, resp.Body)
		if err != nil {
			g.Log().Error(context.Background(), "FastCGI response failed:", err)
			return
		}
	}
}
