package providers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/bitly/oauth2_proxy/api"
)

type DigidentityProvider struct {
	*ProviderData
}

func NewDigidentityProvider(p *ProviderData) *DigidentityProvider {
	p.ProviderName = "Digidentity"
	if p.LoginURL == nil || p.LoginURL.String() == "" {
		p.LoginURL = &url.URL{
			Scheme: "https",
			Host:   "auth.digidentity.eu",
			Path:   "/oauth2/authorize",
		}
	}
	if p.RedeemURL == nil || p.RedeemURL.String() == "" {
		p.RedeemURL = &url.URL{
			Scheme: "https",
			Host:   "auth.digidentity.eu",
			Path:   "/oauth2/token",
		}
	}
	if p.ValidateURL == nil || p.ValidateURL.String() == "" {
		p.ValidateURL = &url.URL{
			Scheme: "https",
			Host:   "gate.digidentity.eu",
			Path:   "/v1/profile",
		}
	}
	if p.Scope == "" {
		p.Scope = "api"
	}
	return &DigidentityProvider{ProviderData: p}
}

func (p *DigidentityProvider) GetEmailAddress(s *SessionState) (string, error) {

	req, err := http.NewRequest("GET",
		p.ValidateURL.String(), nil)

	if err != nil {
		log.Printf("failed building request %s", err)
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+s.AccessToken)
	json, err := api.Request(req)
	if err != nil {
		log.Printf("failed making request %s", err)
		return "", err
	}
	return json.Get("email_address").String()
}
