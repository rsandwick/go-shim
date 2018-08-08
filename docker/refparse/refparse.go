package refparse

import (
	"github.com/docker/distribution/reference"
	"strings"
)

var (
	defaultDomain    = "docker.io"
	officialRepoName = "library"
)

type ImageReference struct {
	Domain string
	Path   string
}

func SplitDomainPath(s string) (*ImageReference, error) {
	fullref, err := reference.ParseNormalizedNamed(s)
	if err != nil {
		return nil, err
	}
	parts := strings.SplitN(fullref.String(), "/", 2)
	ref := &ImageReference{parts[0], parts[1]}
	if ref.Domain == defaultDomain {
		ref.Path = strings.TrimPrefix(ref.Path, officialRepoName+"/")
	}
	return ref, nil
}
