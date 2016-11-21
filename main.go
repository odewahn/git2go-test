package main

import "github.com/libgit2/git2go"

// https://golog.co/blog/article/Git2Go

func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
	ret, cred := git.NewCredSshKey("git", "/Users/odewahn/.ssh/id_rsa.pub", "/Users/odewahn/.ssh/id_rsa", "")
	return git.ErrorCode(ret), &cred
}

func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
	return 0
}

func main() {

	cloneOptions := &git.CloneOptions{}
	cloneOptions.FetchOptions = &git.FetchOptions{
		RemoteCallbacks: git.RemoteCallbacks{
			CredentialsCallback:      credentialsCallback,
			CertificateCheckCallback: certificateCheckCallback,
		},
	}

	_, err := git.Clone(
		"git@git.atlas.oreilly.com:landers/lora-using-arduino.git",
		"web",
		cloneOptions,
	)
	if err != nil {
		panic(err)
	}
}
