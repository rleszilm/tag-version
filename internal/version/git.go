package version

import (
	"errors"
	"io"
	"regexp"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

const (
	gitBranchMatchPattern = `^refs\/heads\/(?P<branch>[\S]*)$`
)

var (
	gitBranchMatch = regexp.MustCompile(gitBranchMatchPattern)
)

// Git is a struct that interacts with a git repo.
type Git struct {
	repo *git.Repository
	tags map[plumbing.Hash]*plumbing.Reference
}

// Branches returns the current branch.
func (g *Git) Branches() ([]string, error) {
	h, err := g.repo.Head()
	if err != nil {
		return nil, err
	}

	refs, err := g.repo.Storer.IterReferences()
	if err != nil {
		return nil, err
	}

	branches := []string{}
	refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Hash() == h.Hash() {
			if !ref.Name().IsBranch() {
				return nil
			}

			matches := gitBranchMatch.FindAllStringSubmatch(ref.Name().String(), -1)
			if matches != nil {
				for _, match := range matches {
					for i, name := range gitBranchMatch.SubexpNames() {
						if name == "branch" {
							branches = append(branches, match[i])
						}
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return branches, nil
}

// Commit returns the commit of the HEAD.
func (g *Git) Commit() (string, error) {
	h, err := g.repo.Head()
	if err != nil {
		return "", err
	}
	return h.String(), nil
}

// Committish returns the shortened committish of the HEAD.
func (g *Git) Committish() (string, error) {
	h, err := g.repo.Head()
	if err != nil {
		return "", err
	}
	return h.String()[:7], nil
}

// Tag returns the closest tag to the HEAD.
func (g *Git) Tag() (string, error) {
	h, err := g.repo.Head()
	if err != nil {
		return "", err
	}

	c, err := g.repo.CommitObject(h.Hash())
	if err != nil {
		return "", err
	}

	ci, err := g.repo.Log(&git.LogOptions{From: c.Hash})
	if err != nil {
		return "", err
	}

	var t string
	for {
		c, err := ci.Next()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return "", err
		}

		if ref, ok := g.tags[c.Hash]; ok {
			t = ref.Name().Short()
			break
		}
	}

	return t, nil
}

// NewGit returns a new Git reader.
func NewGit(path string) (*Git, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}

	tags := map[plumbing.Hash]*plumbing.Reference{}
	ti, err := repo.Tags()
	if err != nil {
		return nil, err
	}

	if err := ti.ForEach(func(r *plumbing.Reference) error {
		tags[r.Hash()] = r
		return nil
	}); err != nil {
		return nil, err
	}

	return &Git{
		repo: repo,
		tags: tags,
	}, nil
}
