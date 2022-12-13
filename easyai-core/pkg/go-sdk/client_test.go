package sdk

import (
	"testing"
)

func TestNewClient(t *testing.T) {

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
			port      int64 = 10088
			apiPrefix       = "/api/v1"
			token           = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoic2h1YWkueWFuZyIsInJlYWxfbmFtZSI6InNodWFpLn" +
				"lhbmciLCJwaG9uZSI6IjEyMzQiLCJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJleHAiOjE3NzgxNTg4NDgsImlhdCI6MTY1MjAxNDg0O" +
				"CwibmJmIjoxNjUyMDE0ODQ4LCJzdWIiOiJzaHVhaS55YW5nIn0.EsFh5O2mtQr3QujWSSDbJeWPyv6g7mxdUvorxR538HbDATepEEHAYy" +
				"jmGrWYqYyFt_0TtjdJKePp1PvvA3mJpA"
			debugHTTP = true
		)
		t.Run(tt.name, func(t *testing.T) {
			clt, err := NewClient(host, port, protocol, apiPrefix, token, debugHTTP)
			if err != nil {
				t.Fatalf("fail to new client: %+v", err)
			}
			input := NewJobGetInput().WithJobID(16)
			job, err := clt.GetJob(input)
			if err != nil {
				t.Fatalf("fail to get job: %+v", err)
			}
			t.Logf("get job: %+v", job)
			for _, task := range job.Tasks {
				t.Logf("get job task: %+v", task)
			}
		})
	}
}
