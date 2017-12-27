// Copyright 2015, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package getaway defines functionality for creating HTTP server endpoints which return
// HTML that includes go-get redirects for "vanity" urls.
package getaway

import (
	"net/http"
	"path"
	"strings"
	"text/template"
)

var page = `<html>
	<head>
		<title>Nothing to see here</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<meta name="go-import" content="{{.Pkg}} {{.RepoType}} {{.RepoPath}}">
	</head>

	<body>
		<p>Nothing to see here, <a href="{{.RepoPath}}">move along</a>.</p>
	</body>
</html>`

var pageTmpl = template.Must(template.New("page").Parse(page))

// Static creates a handler which creates static redirects for
// specific packages.
type Static struct {
	// Package URL to redirect.
	Pkg string

	// Path to repository containing package.
	RepoPath string

	// Type of repository (git, hg, svn, etc...).
	RepoType string
}

// ServeHTTP implements http.Handler.
func (s Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := pageTmpl.Execute(w, s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Dynamic creates a handler which takes the first path
// component as the repo name, and constructs pkgs and repo paths
// from this.
//
// For example for RootPkgPath go.yourdomain.com and RootRepoPath
// https://github.com/youraccount this will setup:
// go.yourdomain.com/reponame/subpkg -> https://github.com/youraccount/reponame
type Dynamic struct {
	// Root package URL to redirect.
	RootPkgPath string

	// Root path to repos where package is located.
	RootRepoPath string

	// Type of repository (git, hg, svn, etc...)
	RepoType string
}

// ServeHTTP implements http.Handler.
func (h Dynamic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fullPath := r.URL.Path
	if fullPath == "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if fullPath[0] == '/' {
		fullPath = fullPath[1:]
	}

	cs := strings.Split(fullPath, "/")
	name := cs[0]

	Static{
		Pkg:      path.Join(h.RootPkgPath, name),
		RepoPath: path.Join(h.RootRepoPath, name),
		RepoType: h.RepoType,
	}.ServeHTTP(w, r)
}
