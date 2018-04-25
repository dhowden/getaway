// Copyright 2016, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
getaway is a simple tool for providing redirects for go-get 'vanity' URLs.
*/
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dhowden/getaway"
)

var (
	listen = flag.String("listen", "localhost:8080", "bind `address` for HTTP server")
	pkg    = flag.String("pkg", "", "`package` URL to redirect")
	repo   = flag.String("repo", "", "path to `repo`")
	vcs    = flag.String("vcs", "git", "`type` of repo (git, hg, svn, etc...)")
)

func main() {
	flag.Parse()

	h := getaway.Static{
		Pkg:      *pkg,
		RepoPath: *repo,
		RepoType: *vcs,
	}
	http.Handle("/", h)

	log.Printf("Starting HTTP server on %v...", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
