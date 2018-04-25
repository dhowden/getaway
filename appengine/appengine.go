package appengine

import (
	"net/http"
	"os"

	"github.com/dhowden/getaway"
)

const (
	PkgEnv      = "PKG"
	RepoPathEnv = "REPO_PATH"
	RepoTypeEnv = "REPO_TYPE"
)

func init() {
	h := getaway.Static{
		Pkg:      os.Getenv(PkgEnv),
		RepoPath: os.Getenv(RepoPathEnv),
		RepoType: os.Getenv(RepoTypeEnv),
	}
	http.Handle("/", h)
}
