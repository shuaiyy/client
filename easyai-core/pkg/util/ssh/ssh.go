package ssh

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

// VerifyKeyPair key pair valid
func VerifyKeyPair(private, public string) error {
	// https://stackoverflow.com/questions/62236441/getting-ssh-short-read-error-when-trying-to-parse-a-public-key-in-golang
	pk, _, _, _, err := ssh.ParseAuthorizedKey([]byte(public))
	if err != nil {
		return err
	}
	publicKey, err := ssh.ParsePublicKey(pk.Marshal())
	if err != nil {
		return err
	}
	privateKey, err := ssh.ParsePrivateKey([]byte(private))
	if err != nil {
		return err
	}
	if string(privateKey.PublicKey().Marshal()) == string(publicKey.Marshal()) &&
		privateKey.PublicKey().Type() == publicKey.Type() {
		return nil
	}
	return fmt.Errorf("public and private keys do not match")
}

// GenSSHConfigScript  a shell script to config ssh connection
func GenSSHConfigScript(insName, insIP, insPort, socks5Addr, socks5Auth, sshPrivateKey, hostPubKey string) string {
	return fmt.Sprintf(scriptTemplate, insName, insIP, insPort, socks5Addr, socks5Auth, sshPrivateKey, hostPubKey)
}
