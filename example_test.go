package ssh_test

import (
	"io"
	"os"

	"github.com/jsthtlf/ssh"
)

func ExampleListenAndServe() {
	ssh.ListenAndServe(":2222", func(s ssh.Session) {
		io.WriteString(s, "Hello world\n")
	})
}

func ExamplePasswordAuth() {
	ssh.ListenAndServe(":2222", nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) (bool, ssh.AuthHandlers) {
			return pass == "secret", ssh.AuthHandlers{}
		}),
	)
}

func ExampleNoPty() {
	ssh.ListenAndServe(":2222", nil, ssh.NoPty())
}

func ExamplePublicKeyAuth() {
	ssh.ListenAndServe(":2222", nil,
		ssh.PublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) (bool, ssh.AuthHandlers) {
			data, _ := os.ReadFile("/path/to/allowed/key.pub")
			allowed, _, _, _, _ := ssh.ParseAuthorizedKey(data)
			return ssh.KeysEqual(key, allowed), ssh.AuthHandlers{}
		}),
	)
}

func ExampleHostKeyFile() {
	ssh.ListenAndServe(":2222", nil, ssh.HostKeyFile("/path/to/host/key"))
}
