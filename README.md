# getaway

Getaway is a simple tool for creating `go get` vanity URLs.

## Usage

Assuming that you have `GOPATH` configured (see [https://golang.org/doc/code.html#GOPATH](https://golang.org/doc/code.html#GOPATH)), and `GOPATH/bin` is contained in your `PATH`:

    $ go get github.com/dhowden/getaway
    $ getaway -pkg tchaik.com -repo https://github.com/tchaik/tchaik
