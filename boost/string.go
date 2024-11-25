package boost

import "strings"

func Join(ss ...string) string {
	var b strings.Builder
	b.Grow(64)
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}
