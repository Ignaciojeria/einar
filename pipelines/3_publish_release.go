package pipelines

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

const changelogPath = "dist/CHANGELOG.md" // Specify the constant changelog path
const distFolder = "dist" // Specify the constant dist folder path

func PublishRelease(ctx context.Context, tagName string) error {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Read changelog file
	changelogBytes, err := ioutil.ReadFile(changelogPath) // Use the constant changelog path
	if err != nil {
		return err
	}
	changelog := string(changelogBytes)

	// Upload release assets from the dist folder
	files, err := ioutil.ReadDir(distFolder)
	if err != nil {
		return err
	}

	// Create the release on GitHub
	release := &github.RepositoryRelease{
		TagName:         github.String(tagName),
		TargetCommitish: github.String(os.Getenv("GITHUB_BRANCH")),
		Name:            github.String(tagName),
		Body:            github.String(changelog), // Use the changelog as the release body
		Draft:           github.Bool(false),
		Prerelease:      github.Bool(false),
	}

	repoOwner := os.Getenv("GITHUB_OWNER")
	repoName := os.Getenv("GITHUB_REPO")
	release, _, err = client.Repositories.CreateRelease(ctx, repoOwner, repoName, release)
	if err != nil {
		return err
	}

	// Upload the release assets
	for _, file := range files {
		// Skip directories
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(distFolder, file.Name())
		fileName := file.Name()

		// Open the release asset file
		fileContent, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer fileContent.Close()

		// Create the release asset
		asset, _, err := client.Repositories.UploadReleaseAsset(ctx, repoOwner, repoName, *release.ID, &github.UploadOptions{
			Name:      filepath.Base(fileName),
			Label:     strings.TrimSuffix(fileName, filepath.Ext(fileName)),
			MediaType: "application/octet-stream",
		}, fileContent)
		if err != nil {
			return err
		}

		fmt.Printf("Uploaded release asset: %s\n", *asset.BrowserDownloadURL)
	}

	return nil
}
