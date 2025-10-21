package utils

import (
	"net"
	"net/http"
	"strings"
)

// ClientIP 尽可能返回客户端真实 IP。
// 必须保证“最外层”Nginx 类似下面这样配置：
//   proxy_set_header X-Real-IP       $remote_addr;
//   proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

func ParseIP(r *http.Request) string {
	// 1. 先取 X-Real-IP，简单可靠
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	// 2. 取 X-Forwarded-For 第一段
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// 去掉空格并按 "," 拆分
		for _, seg := range strings.Split(xff, ",") {
			seg = strings.TrimSpace(seg)
			if seg != "" && net.ParseIP(seg) != nil {
				return seg
			}
		}
	}
	// 3. 兜底：直接连接的地址（会是 Nginx 的地址）
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}
