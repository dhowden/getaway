# getaway
[![Build Status](https://travis-ci.org/dhowden/getaway.svg?branch=master)](https://travis-ci.org/dhowden/getaway)

Getaway is a simple HTTP server for creating `go get` vanity URLs.

## Usage

Assuming that you have `GOPATH` configured (see [https://golang.org/doc/code.html#GOPATH](https://golang.org/doc/code.html#GOPATH)), and `GOPATH/bin` is contained in your `PATH`:

    $ go get github.com/dhowden/getaway
    $ getaway -pkg tchaik.com -repo https://github.com/tchaik/tchaik
