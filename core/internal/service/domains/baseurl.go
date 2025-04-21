package domains

import (
	v1 "billionmail-core/api/domains/v1"
	"billionmail-core/internal/consts"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	baseurl = ""
)

// GetBaseURL get baseurl of console panel
func GetBaseURL() string {
	return baseurl
}

func UpdateBaseURL() {
	g.Log().Debug(context.Background(), "UpdateBaseURL --> Starting")
	defer func() {
		g.Log().Debug(context.Background(), "UpdateBaseURL --> Ending")
	}()

	scheme := "https"
	serverPort := g.Server(consts.DEFAULT_SERVER_NAME).GetListenedHTTPSPort()

	if serverPort == -1 {
		scheme = "http"
		serverPort = g.Server(consts.DEFAULT_SERVER_NAME).GetListenedPort()
	}

	withPort := true

	if serverPort == -1 || serverPort == 80 || serverPort == 443 {
		withPort = false
	}

	serverIP, localIP, err := public.GetServerIPAndLocalIP()

	if err != nil {
		g.Log().Error(context.Background(), "failed to get server IP and local IP", err)
		return
	}

	if appEnv, err := public.DockerEnv("APP_ENV"); err == nil && appEnv != "" {
		switch appEnv {
		case "development":
			baseurl = scheme + "://" + localIP
			if withPort {
				baseurl += ":" + gconv.String(serverPort)
			}
			return
		}
	}

	hostname, err := public.DockerEnv("BILLIONMAIL_HOSTNAME")

	if err != nil || !ValidateARecord(v1.DNSRecord{
		Type:  "A",
		Host:  hostname,
		Value: serverIP,
		Valid: false,
	}) {
		baseurl = scheme + "://" + serverIP
		if withPort {
			baseurl += ":" + gconv.String(serverPort)
		}
		return
	}

	baseurl = scheme + "://" + hostname
	if withPort {
		baseurl += ":" + gconv.String(serverPort)
	}

	return
}
