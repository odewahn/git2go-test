package main

import (
	"fmt"
	"log"
	"os"

	"github.com/libgit2/git2go"
)

func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
	//ret, cred := git.NewCredSshKey("git", "/Users/odewahn/.ssh/id_rsa.pub", "/Users/odewahn/.ssh/id_rsa", "")
	fmt.Print("Enter your username: ")
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

	if len(os.Args) < 2 {
		log.Fatal("You must supply a git repo to download!")
	}

	cloneOptions := &git.CloneOptions{}
	cloneOptions.FetchOptions = &git.FetchOptions{
		RemoteCallbacks: git.RemoteCallbacks{
			CredentialsCallback:      credentialsCallback,
			CertificateCheckCallback: certificateCheckCallback,
		},
	}

	fmt.Println("Cloning repo to a directory named web.  Yup, web.  It's hardcoded, all right, since this is just a test..")
	_, err := git.Clone(os.Args[1], "web", cloneOptions)
	if err != nil {
		log.Fatal(err)
	}
}
