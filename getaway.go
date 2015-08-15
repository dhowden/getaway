// Copyright 2015, David Howden
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
	"text/template"
)

var page = `<html>
	<head>
		<title>Nothing to see here</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="{{.Pkg}} {{.VCS}} {{.Repo}}">
	</head>

	<body>
		<p>Nothing to see here, <a href="{{.Repo}}">move along</a>.</p>
	</body>
</html>`

var pageTmpl = template.Must(template.New("page").Parse(page))

var listen string
var pkg string
var repo, vcs string

func init() {
	flag.StringVar(&listen, "listen", "localhost:8080", "bind address for HTTP server")
	flag.StringVar(&pkg, "pkg", "", "package URL to redirect")
	flag.StringVar(&repo, "repo", "", "path to repo")
	flag.StringVar(&vcs, "vcs", "git", "type of repo (git, hg, svn, etc...)")
}

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := pageTmpl.Execute(w, struct {
			Pkg, VCS, Repo string
		}{
			Pkg:  pkg,
			VCS:  vcs,
			Repo: repo,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(listen, nil))
}
