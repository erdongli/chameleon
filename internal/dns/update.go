package dns

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	urlFormat = "https://%s:%s@domains.google.com/nic/update"

	paramHostname = "hostname"
	paramIP       = "myip"

	headerUserAgent   = "user-agent"
	userAgent         = "chameleon"
	headerContentType = "content-type"
	contentType       = "application/x-www-form-urlencoded"

	respGood     = "good"
	respNoChange = "nochg"
)

type Updater struct {
	url, hostname, ip string
	client            *http.Client
}

func New(username, password, hostname string) *Updater {
	return &Updater{
		url:      fmt.Sprintf(urlFormat, url.QueryEscape(username), url.QueryEscape(password)),
		hostname: hostname,
		client:   &http.Client{},
	}
}

func (u *Updater) Update(ip string) (bool, error) {
	if ip == u.ip {
		return false, nil
	}

	v := url.Values{}
	v.Add(paramHostname, u.hostname)
	v.Add(paramIP, ip)

	req, err := http.NewRequest(http.MethodPost, u.url, strings.NewReader(v.Encode()))
	if err != nil {
		return false, err
	}

	req.Header.Add(headerUserAgent, userAgent)
	req.Header.Add(headerContentType, contentType)

	resp, err := u.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	r := string(b)
	if !strings.HasPrefix(r, respGood) && !strings.HasPrefix(r, respNoChange) {
		return false, fmt.Errorf("error response: %s", r)
	}
	u.ip = ip

	return !strings.HasPrefix(r, respNoChange), nil
}
