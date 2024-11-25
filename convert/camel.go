package convert

import (
	"bytes"
	"strings"
)

var commonInitialisms = []string{"ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "VIP", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "UID", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS"}
var commonInitialismsReplacer *strings.Replacer
var uncommonInitialismsReplacer *strings.Replacer

func init() {
	var commonInitialismsForReplacer = make([]string, 0, len(commonInitialisms))
	var uncommonInitialismsForReplacer = make([]string, 0, len(commonInitialisms))
	for _, initialism := range commonInitialisms {
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
		uncommonInitialismsForReplacer = append(uncommonInitialismsForReplacer, strings.Title(strings.ToLower(initialism)), initialism)
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)
	uncommonInitialismsReplacer = strings.NewReplacer(uncommonInitialismsForReplacer...)
}

func ToCamelCase(name string) string {
	if name == "" {
		return ""
	}

	temp := strings.Split(name, "_")
	var s string
	for _, v := range temp {
		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z') {
				vv[0] -= 32
			}
			s += string(vv)
		}
	}

	s = uncommonInitialismsReplacer.Replace(s)
	return s
}

func FromCamelCase(name string) string {
	const (
		lower = 0
		upper = 1
	)

	if name == "" {
		return ""
	}

	var (
		value = commonInitialismsReplacer.Replace(name)
		buf   = bytes.NewBufferString("")
	)
	var lastCase, currCase, nextCase, nextNumber int

	for i, v := range value[:len(value)-1] {
		if value[i+1] >= 'A' && value[i+1] <= 'Z' {
			nextCase = upper
		} else {
			nextCase = lower
		}
		if value[i+1] >= '0' && value[i+1] <= '9' {
			nextNumber = upper
		} else {
			nextNumber = lower
		}

		if i > 0 {
			if currCase == upper {
				if lastCase == upper && (nextCase == upper || nextNumber == upper) {
					buf.WriteRune(v)
				} else {
					if value[i-1] != '_' && value[i+1] != '_' {
						buf.WriteRune('_')
					}
					buf.WriteRune(v)
				}
			} else {
				buf.WriteRune(v)
				if i == len(value)-2 && (nextCase == upper && nextNumber == lower) {
					buf.WriteRune('_')
				}
			}
		} else {
			currCase = upper
			buf.WriteRune(v)
		}
		lastCase = currCase
		currCase = nextCase
	}

	buf.WriteByte(value[len(value)-1])

	s := strings.ToLower(buf.String())
	return s
}
