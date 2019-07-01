package hoo

import (
	"testing"
)

func TestMinify(t *testing.T) {
	type args struct {
		html string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"<html></html>"}, "<html></html>"},
		{"case1", args{"<html>   </html>"}, "<html></html>"},
		{"case1", args{"<html>\n    </html>"}, "<html></html>"},
		{"case1", args{"<html>\n    \t</html>"}, "<html></html>"},
		{"case1", args{"<html\\>   \\</html>"}, "<html\\>\\</html>"},
		{"case1", args{"<html\\>   </html>"}, "<html\\></html>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Minify(tt.args.html); got != tt.want {
				t.Errorf("Minify() = %v, want %v", got, tt.want)
			}
		})
	}
}
