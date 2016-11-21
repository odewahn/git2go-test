package main

import (
	"fmt"
	"log"
	"os"

	"github.com/libgit2/git2go"
)

// Useful articles:
//   git2go godocs: https://godoc.org/github.com/libgit2/git2go
//   private git access for libgit2: https://golog.co/blog/article/Git2Go
//   default ssh keyfile names: https://help.github.com/articles/checking-for-existing-ssh-keys

func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
	//ret, cred := git.NewCredSshKey("git", "/Users/odewahn/.ssh/id_rsa.pub", "/Users/odewahn/.ssh/id_rsa", "")
	fmt.Print("Enter username: ")
	var user string
	fmt.Scanln(&user)

	fmt.Print("Enter password: ")
	var password string
	fmt.Scanln(&password)

	ret, cred := git.NewCredUserpassPlaintext(user, password)
	return git.ErrorCode(ret), &cred
}

func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
	return 0
}

func main() {

	//url := "git@git.atlas.oreilly.com:landers/lora-using-arduino.git"
	//url := "https://git.atlas.oreilly.com/landers/lora-using-arduino.git"

	cloneOptions := &git.CloneOptions{}
	cloneOptions.FetchOptions = &git.FetchOptions{
		RemoteCallbacks: git.RemoteCallbacks{
			CredentialsCallback:      credentialsCallback,
			CertificateCheckCallback: certificateCheckCallback,
		},
	}

	_, err := git.Clone(os.Args[1], "web", cloneOptions)
	if err != nil {
		log.Fatal(err)
	}
}
