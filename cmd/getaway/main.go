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

var listen string
var pkg string
var repo, vcs string

func init() {
	flag.StringVar(&listen, "listen", "localhost:8080", "bind `address` for HTTP server")
	flag.StringVar(&pkg, "pkg", "", "`package` URL to redirect")
	flag.StringVar(&repo, "repo", "", "path to `repo`")
	flag.StringVar(&vcs, "vcs", "git", "`type` of repo (git, hg, svn, etc...)")
}

func main() {
	flag.Parse()

	h := getaway.Static{
		Pkg:      pkg,
		RepoPath: repo,
		RepoType: vcs,
	}
	http.Handle("/", h)
	log.Fatal(http.ListenAndServe(listen, nil))
}
