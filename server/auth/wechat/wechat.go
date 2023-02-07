package wechat

import (
	"crypto/tls"
	"fmt"
	"github.com/medivhzhan/weapp/v3"
	"net/http"
	"time"
)

// Service implements a WeChat auth service.
type Service struct {
	AppID     string
	AppSecret string
}

// Resolve resolves authorization code to return WeChat openid.
func (s *Service) Resolve(code string) (string, error) {
	cli := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 10 * time.Second,
	}
	sdk := weapp.NewClient(
		s.AppID,
		s.AppSecret,
		weapp.WithHttpClient(cli),
	)
	resp, err := sdk.Login(code)
	if err != nil {
		return "", fmt.Errorf("weapp.Login failed: %v", err)
	}
	return resp.OpenID, nil
}
