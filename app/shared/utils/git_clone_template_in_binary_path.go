package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func GitCloneTemplateInBinaryPath(repositoryUrl, userCreds, tag string) (string, error) {
	targetPath, err := GetTemplateFolderPath(repositoryUrl)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var auth *http.BasicAuth
	if userCreds != "no-auth" {
		user, token, err := SplitCredentials(userCreds)
		if err != nil {
			fmt.Println("Failed to parse user credentials:", err)
			return "", err
		}
		auth = &http.BasicAuth{Username: user, Password: token}
	}

	// Clonar en un directorio temporal
	tmpDir, err := ioutil.TempDir("", "git-")
	if err != nil {
		return "", fmt.Errorf("error creating temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Limpia el directorio temporal después

	_, err = git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL:      repositoryUrl,
		Progress: os.Stdout,
		Auth:     auth,
	})
	if err != nil {
		fmt.Println("Failed to clone repository into temp folder:", err)
		return "", err
	}

	// Abrir el repositorio clonado
	repo, err := git.PlainOpen(tmpDir)
	if err != nil {
		fmt.Println("Failed to open repository:", err)
		return "", err
	}

	effectiveTag := tag
	if tag == "" {
		// Obtén el tag más reciente si no se proporciona uno
		tagRefs, err := repo.Tags()
		if err != nil {
			fmt.Println("Failed to list tags:", err)
			return "", err
		}

		err = tagRefs.ForEach(func(ref *plumbing.Reference) error {
			effectiveTag = ref.Name().Short()
			return nil
		})
		if err != nil {
			fmt.Println("Failed to iterate over tags:", err)
			return "", err
		}
	}

	w, err := repo.Worktree()
	if err != nil {
		fmt.Println("Failed to get worktree:", err)
		return "", err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", effectiveTag)),
	})
	if err != nil {
		fmt.Println("Failed to checkout tag:", err)
		return "", err
	}

	// Mover contenido del directorio temporal al directorio final
	tagFolderPath := filepath.Join(targetPath, effectiveTag)
	if err := os.MkdirAll(tagFolderPath, os.ModePerm); err != nil {
		fmt.Println("Failed to create tag folder:", err)
		return "", err
	}
	if err := moveDirectoryContents(tmpDir, tagFolderPath); err != nil {
		fmt.Println("Failed to move repository content:", err)
		return "", err
	}

	fmt.Println("Repository cloned to:", tagFolderPath)
	return tagFolderPath, nil
}

func moveDirectoryContents(srcDir, destDir string) error {
	entries, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {

		srcPath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if entry.IsDir() {
			if err := os.Rename(srcPath, destPath); err != nil {
				if err := copyDir(srcPath, destPath); err != nil {
					return err
				}
				if err := os.RemoveAll(srcPath); err != nil {
					return err
				}
			}
		} else {
			if err := moveFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func moveFile(src, dest string) error {
	err := os.Rename(src, dest)
	if err != nil {
		if os.IsExist(err) {
			if copyErr := copyFile(src, dest); copyErr != nil {
				return copyErr
			}
			if delErr := os.Remove(src); delErr != nil {
				return delErr
			}
		} else {
			return err
		}
	}
	return nil
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func copyDir(srcDir, destDir string) error {
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}
