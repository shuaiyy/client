package sdk

import (
	"testing"
)

func TestClient_GetClusterConfigs(t *testing.T) {

	tests := []struct {
		name    string
		want    *Client
		wantErr bool
	}{
		{
			name:    "client",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		var (
			protocol        = "http"
			host            = "127.0.0.1"
			port      int64 = 80
			apiPrefix       = "/api/v1"
			token           = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoic2h1YWkueWFuZyIsInJlYWxfbmFtZSI6" +
				"IuadqOW4hSIsInBob25lIjoiIiwiZW1haWwiOiJzaHVhaS55YW5nQG1paG95by5jb20iLCJleHAiOjE3ODk0NjM0NDAsImlhdCI6" +
				"MTY2MzMxOTQ0MCwibmJmIjoxNjYzMzE5NDQwLCJzdWIiOiJzaHVhaS55YW5nIn0.hFKDbeLcJgPmNcsGTeUOrbI1cGKtvRQIrhqsTP" +
				"l7wamB9ccLm-37WtHev9jg5Q-Sna049dvMz6BUf49EczPdtQ"
			debugHTTP = true
		)
		t.Run(tt.name, func(t *testing.T) {
			clt, err := NewClient(host, port, protocol, apiPrefix, token, debugHTTP)
			if err != nil {
				t.Fatalf("fail to new client: %+v", err)
			}
			input := NewClusterConfigsInput()
			cfgs, err := clt.GetClusterConfigs(input)
			if err != nil {
				t.Fatalf("fail to get cluste configs: %+v", err)
			}
			t.Logf("get cluste cfgs: %+v", cfgs)
			for cName, cfg := range cfgs {
				t.Logf("get cluster<%s> config: %+v", cName, cfg)
			}
		})
	}
}
