package utils

import (
	"net/http"
	"net"
)

func ParseIP(r *http.Request) string {
	//authFastRegisterRequest.Header.Set("X-FORWARDED-FOR", "125.6.5.21,125.6.5.22")
	if ipProxy := r.Header.Get("X-FORWARDED-FOR"); len(ipProxy) > 0 {
		return ipProxy
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
