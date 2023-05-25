package bilibili

import (
	"net/http"

	"github.com/eric2788/common-utils/request"
	"github.com/sirupsen/logrus"
)

var (
	
	apiRequester  = generateRequester("https://api.bilibili.com/x", request.WithDefaultDecoder(XRespDecoder))

	liveRequester = generateRequester("https://api.live.bilibili.com",
		request.WithDefaultDecoder(V1RespDecoder),
		request.WithHeaders(map[string]string{
			"Referer": "https://live.bilibili.com/",
			"Origin":  "https://live.bilibili.com/",
		}),
	)

	logger = logrus.WithField("service", "bilibili")
)

func generateRequester(baseUrl string, factories ...request.RequesterFactory) *request.Requester {
	fixed := []request.RequesterFactory{
		request.WithBaseUrl(baseUrl),
		request.AddHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64"),
		request.AddHeader("Accept", "application/json"),
		request.AddHeader("Content-Type", "application/json"),

		request.AddRequestIntercepter(func(r *http.Request) error {
			logrus.Debug("prepare to request ", r.URL.String())
			return nil
		}),
	}
	return request.New(append(fixed, factories...)...)
}
