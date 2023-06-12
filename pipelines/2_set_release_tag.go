package pipelines

import (
	"context"
	"os"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func SetReleaseTag(ctx context.Context, tagName, tagMessage string) error {
	// Create a new OAuth2 client using the GitHub access token
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Get the reference to the current commit of the branch
	ref, _, err := client.Git.GetRef(ctx, os.Getenv("GITHUB_OWNER"), os.Getenv("GITHUB_REPO"), "refs/heads/"+os.Getenv("GITHUB_BRANCH"))
	if err != nil {
		return err
	}

	// Create a new git tag and associate it with the current commit
	tag := &github.Tag{
		Tag:     github.String(tagName),
		Message: github.String(tagMessage),
		Object:  &github.GitObject{SHA: ref.Object.SHA, Type: github.String("commit")},
	}
	tag, _, err = client.Git.CreateTag(ctx, os.Getenv("GITHUB_OWNER"), os.Getenv("GITHUB_REPO"), tag)
	if err != nil {
		return err
	}

	// Create a reference to the new tag
	refObj := github.Reference{
		Ref:    github.String("refs/tags/" + *tag.Tag),
		Object: &github.GitObject{SHA: tag.SHA, Type: github.String("tag")},
	}
	_, _, err = client.Git.CreateRef(ctx, os.Getenv("GITHUB_OWNER"), os.Getenv("GITHUB_REPO"), &refObj)
	if err != nil {
		return err
	}

	// Update the branch to point to the new commit
	refObj.Ref = github.String("refs/heads/" + os.Getenv("GITHUB_BRANCH"))
	refObj.Object.SHA = tag.SHA
	_, _, err = client.Git.UpdateRef(ctx, os.Getenv("GITHUB_OWNER"), os.Getenv("GITHUB_REPO"), &refObj, false)
	if err != nil {
		return err
	}

	return nil
}
