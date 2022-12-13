package ssh

import (
	"os"
	"testing"
)

func TestVerifyKeyPair(t *testing.T) {
	type args struct {
		private string
		public  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "a",
			args: args{
				private: DefaultPrivateKey,
				public:  DefaultPublicKey,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyKeyPair(tt.args.private, tt.args.public); (err != nil) != tt.wantErr {
				t.Errorf("VerifyKeyPair() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenSSHConfigScript(t *testing.T) {
	type args struct {
		insName       string
		insIP         string
		insPort       string
		socks5Addr    string
		socks5Auth    string
		sshPrivateKey string
		hostPubKey    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "111",
			args: args{
				insName:       "train-11-chief-0",
				insIP:         "10.14.128.249",
				insPort:       "23322",
				socks5Addr:    "47.117.66.151:23333",
				socks5Auth:    "seelie-job-proxy:c2VlbGllLXBhc3N3b3Jk",
				sshPrivateKey: "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn\nNhAAAAAwEAAQAAAQEAqUmohVLsFE60OobUtGRw+vckMt3js20uo/76eyj5frN2dYw7etId\nieA6BU2AcXbP1D96ikVTa6uycgelxtZ1bxMXnSKzzXm2iq/tuhLeFys8xTT/JiAmAuX4UP\nvDRhWBh9+SaCaKq5KHTfZQGXHyPnpoY0JOpZwBayKgoEKyhUFohK5WrPc87JS2pzBv8ynt\nVJ80lz/qhdcBgk9gK039OLW0gbh4d+2ZtkT+jD1NVQFN16NIRZ2rYlnVxRa6E9LjnH6Avx\nGVKuBYkp7wwQIUuLsnniSMMJ0oQKZOnu/HzIipyHo8/ae84BsHFnSfnG8JBvlIZCIRDxcm\nXt8JspSZ4QAAA9i+5XpGvuV6RgAAAAdzc2gtcnNhAAABAQCpSaiFUuwUTrQ6htS0ZHD69y\nQy3eOzbS6j/vp7KPl+s3Z1jDt60h2J4DoFTYBxds/UP3qKRVNrq7JyB6XG1nVvExedIrPN\nebaKr+26Et4XKzzFNP8mICYC5fhQ+8NGFYGH35JoJoqrkodN9lAZcfI+emhjQk6lnAFrIq\nCgQrKFQWiErlas9zzslLanMG/zKe1UnzSXP+qF1wGCT2ArTf04tbSBuHh37Zm2RP6MPU1V\nAU3Xo0hFnatiWdXFFroT0uOcfoC/EZUq4FiSnvDBAhS4uyeeJIwwnShApk6e78fMiKnIej\nz9p7zgGwcWdJ+cbwkG+UhkIhEPFyZe3wmylJnhAAAAAwEAAQAAAQBfpYJVcbh48M7bknpz\nQZyj7ybApqWUJsgHWHTlSQ1ODM+NMqIYjsaps8qUXGmJsftSjFsL7IdpeiTkUHXVli3biA\nn7ejPkkDQWv7etPiPFK2S2d28Bd4CCerSF7Pkzi8sXnbX2qnAG6E8SOWygM8UOj9KS4k/V\nOgODKJlgcs6ygorozqKAxh2FVTc3TDBFmwtPbqw792Jqhrzb7u9bjuIQHNo66XLUi3ElWh\nBZZm/9lx8LSC2jMO+3RaTdiI7jLHpfWxMFo/0CkO7XJ/CBgN0ouGP9AkAUjheYqchBnZBj\nvWpX8ch5vUrSUxP1QtRxurR7fod7rit2ZWhH8a6VU2QBAAAAgFVEE/dE+nIcS3CDanZFhk\nWY2ROggQFgCyZvGK5/hOs7wMWfGRqZ78XrAl+mN+tKc2ICfBG5OD4BhDQA0KXRrx1C7WRT\nZlXzqDPB17Qt8EQq3baHE4of49pMgEeJDhk6IMIgdudUxs3lnb/qGMyNKn7e80VUn7HAE1\nOOP8f+Tx5AAAAAgQDTtnccvM2LBOrERkeehxcz3VlIQGcXSoY6zYSAJVkMdIhzdiIa2Ol0\nGr9g5zrzmYKCiRWGhGHoJdmRF/xnP6+d5Dyy66tIHTe687xxegw0tIjmkICrqpEwWtemrB\nKFphMb/qZRZgn+bqC+dmoIy+BbJIFsnqaYQp9rpdxgoSst8QAAAIEAzLNE8qX1/grSX84K\nCM1eBaM7AHs3Mb4PJ93VY4gx9pvAxqWgZHAnQboV/U5Y+7TA8J/H9rfQI063U/KBo4buG/\nZU+t7T7wmGUGhw5NDBLKMBvidX9Q08XPTsxAW9afbMLok5tunLO2rwIef60GpoFegWxQ01\nEZ1rw5cs0QIh+vEAAAAdYmlsaWJpbGlAYmlsaWJpbC10ZXN0LTIubG9jYWwBAgMEBQY=\n-----END OPENSSH PRIVATE KEY-----",
				hostPubKey:    "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAzfg/8Z0VmJfK2n8ZNvzFZElaRzyEEgfq+nQjLPxryqdfEb0OKVIF5BbTvZAIJEa0P0CW02b/DCafHc+MqJzeU= root@debian",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenSSHConfigScript(tt.args.insName, tt.args.insIP, tt.args.insPort, tt.args.socks5Addr, tt.args.socks5Auth, tt.args.sshPrivateKey, tt.args.hostPubKey)
			t.Log(got)
			if err := os.WriteFile("/tmp/ssh-config.sh", []byte(got), 0644); err != nil {
				t.Fatal(err)
			}
			t.Log("bash /tmp/ssh-config.sh")
		})
	}
}
