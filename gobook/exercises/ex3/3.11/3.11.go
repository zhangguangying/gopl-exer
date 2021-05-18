package __11

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	matissaStart := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		matissaStart = 1
	}
	matissaEnd := strings.Index(s, ".")
	if matissaEnd == -1 {
		matissaEnd = len(s)
	}
	mantissa := s[matissaStart:matissaEnd]
	pre := len(mantissa) % 3
	if pre > 0 {
		buf.WriteString(mantissa[:pre])
		if len(mantissa) > pre {
			buf.WriteString(",")
		}
	}
	for i, c := range mantissa[pre:] {
		if i%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(c)
	}
	buf.WriteString(s[matissaEnd:])
	return buf.String()
}
