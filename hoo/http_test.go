package hoo

import "testing"

func TestHost(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test0",args{url: "https://apd-69d07362356a6bd5f7eb063f411a3ebb.v.smtcdns.com/moviets.tc.qq.com/AAPpXrjOQGrBTYnU0yHy6cFbW6eWf2H_u-Rg8DiA4ufk/uwMROfz2r5xgoaQXGdGnC2df64gVTKzl5C_X6A3JOVT0QIb-/M1ue3ZPek8jJdXEFUrnLVcKCQTXxGfk6Nne3ouqL1t-xtS__FrarJ1xbP7Ats0W0y7ulqXvns2AnEISQt2Z56IZOIeEdcA4lHAAhjnIRV5jpl9wCIhZqkjfidSAZvR7NC4t_90fQO-192XVa_AGUNcbAKRPu1hMjQWfmsC8AenM/00_d0030amdf50.321004.1.ts?index=0&start=0&end=7200&brs=0&bre=980607&ver=4"},"apd-69d07362356a6bd5f7eb063f411a3ebb.v.smtcdns.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Host(tt.args.url); got != tt.want {
				t.Errorf("Host() = %v, want %v", got, tt.want)
			}
		})
	}
}
