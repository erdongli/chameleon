package ip

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func Get() (string, error) {
	resp, err := http.Get("https://toolbox.erdongli.com/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ip := string(body)

	if net.ParseIP(ip) == nil {
		return "", fmt.Errorf("invalid ip %s", ip)
	}

	return ip, nil
}
