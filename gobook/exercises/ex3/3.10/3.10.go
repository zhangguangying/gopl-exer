package __10

import "bytes"

func comma(s string) string {
	var buf bytes.Buffer
	pre := len(s) % 3
	if pre == 0 {
		pre = 3
	}
	buf.WriteString(s[:pre])
	for i := pre; i < len(s); i+=3 {
		buf.WriteString(",")
		buf.WriteString(s[i:i+3])
	}
	return buf.String()
}
