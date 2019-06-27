package htmlutil

import "testing"

func TestMinify(t *testing.T) {
	tests := []struct {
		html string
		want string
	}{
		{"<html></html>", "<html></html>"},
		{"<html>   </html>", "<html></html>"},
		{"<html>\n    </html>", "<html></html>"},
		{"<html>\n    \t</html>", "<html></html>"},
		{"<html\\>   \\</html>", "<html\\>   \\</html>"},
		{"<html\\>   </html>", "<html\\>   </html>"},
	}
	for _, tt := range tests {
		if got := Minify(tt.html); got != tt.want {
			t.Errorf("Minify(%v) = %v, want %v", tt.html, got, tt.want)
		}
	}
}
