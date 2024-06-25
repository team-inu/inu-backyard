package captcha

import (
	"encoding/json"
	"net/http"
	"net/url"

	errs "github.com/team-inu/inu-backyard/entity/error"
)

type Turnstile struct {
	secret string
}

type cloudflareToken struct {
	Success bool `json:"success"`
}

func NewTurnstile(secret string) *Turnstile {
	return &Turnstile{
		secret: secret,
	}
}

func (t Turnstile) Validate(token string, ip string) (bool, error) {
	resp, err := http.PostForm("https://challenges.cloudflare.com/turnstile/v0/siteverify", url.Values{
		"secret":   {t.secret},
		"response": {token},
		"remoteip": {ip},
	})
	if err != nil {
		return false, errs.New(0, "unexpected error when validate token", err)
	}
	defer resp.Body.Close()

	var cfRes cloudflareToken
	err = json.NewDecoder(resp.Body).Decode(&cfRes)
	if err != nil {
		return false, errs.New(0, "unexpected error when validate token", err)
	}

	return cfRes.Success, nil
}
