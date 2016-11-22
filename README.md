# Background

This repo shows how to use [git2go](https://github.com/libgit2/git2go) with the c-based [libgit2](https://libgit2.github.com/) package to create a simple app that will clone a private repo using the https protocol.  

It's OSX specific (for now), but the end has a start at how to [xgo](https://github.com/karalabe/xgo) to cross-compile for other platforms.  

Here's some useful links to review:

* https://godoc.org/github.com/libgit2/git2go
* https://golog.co/blog/article/Git2Go
* https://help.github.com/articles/checking-for-existing-ssh-keys

## Install

Even if you're planning on using xgo, you need to have  [cmake](https://cmake.org/) so that you can build all the various c files into something go can use.

```
brew install cmake
```

## Get git2go

git2go is the Golang wrapper on top of libgit2:

`go get github.com/libgit2/git2go`

## Load the c file for libgit2

This step initially confused me, but basically here you just go through the steps listed in

```
cd $GOPATH/src/github.com/libgit2/git2go
git checkout next
git submodule update --init # get libgit2
make install
```

## Run `main.go`

At this point, you should be able to run main.go, like this:

```
go run main.go https://git.atlas.oreilly.com/landers/lora-using-arduino.git
```

# Building a release for equinox

Once you've got it working and compiled, you can use equinox to package and distribute it.

First, set your equinox-specific environment variables:

```
export EQUINOX_APP=<your app key>
export EQUINOX_KEY=<path to your equinox key>
export EQUINOX_TOKEN=<your token>
```

Then run the release tool:

```
equinox release \
  --version="0.0.1" \
  --platforms="darwin_amd64" \
  --signing-key=$EQUINOX_KEY \
  --app=$EQUINOX_APP \
  --token=$EQUINOX_TOKEN \
  .
```


# Cross compiling the binary with xgo:

[xgo](https://github.com/karalabe/xgo) is a Docker-based cross compiler for Go.  Basically, it puts a bunch of build environments as Docker images so that you can install their build toolchains more easily.

To install it (from the projects README):

```
docker pull karalabe/xgo-latest
```

To prevent having to remember a potentially complex Docker command every time, a lightweight Go wrapper was written on top of it.

```
go get github.com/karalabe/xgo
```

Once you have this going, you can build a binary for a specific platform like this:

```
xgo --targets=darwin/* .
```


# Running Godeps

Undoubtedly, at some point, you'll need to use Godeps to manage all the dependencies, so here are the commands.

```
$ godep save ./...       # saves your GOPATH to the Godep folder
$ godep go build ./...  # builds using the Godep vendored dependencies
```

For details, see:

*  https://blog.codeship.com/godep-dependency-management-in-golang/
