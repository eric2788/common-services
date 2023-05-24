package bilibili

import (
	"net/http"

	"github.com/eric2788/common-utils/request"
	"github.com/sirupsen/logrus"
)



var (

	requester = request.New(
		request.WithBaseUrl(""),
		request.AddHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64"),
		request.AddHeader("Accept", "application/json"),
		request.AddHeader("Content-Type", "application/json"),


		request.AddRequestIntercepter(func(r *http.Request) error {
			logrus.Debug("prepare to request ", r.URL.String())
			return nil
		}),
	)

	logger = logrus.WithField("service", "bilibili")
)
