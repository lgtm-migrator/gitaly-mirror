package objectpool

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"gitlab.com/gitlab-org/gitaly/internal/git"
	"gitlab.com/gitlab-org/gitaly/internal/helper"
	"gitlab.com/gitlab-org/gitaly/proto/go/gitalypb"
)

// ObjectPool are a way to dedup objects between repositories, where the objects
// live in a pool in a distinct repository which is used as an alternate object
// store for other repositories.
type ObjectPool struct {
	storageName string
	storagePath string

	relativePath string
}

// NewObjectPool will initialize the object with the required data on the storage
// shard. If the shard cannot be found, this function returns an error
func NewObjectPool(storageName, relativePath string) (pool *ObjectPool, err error) {
	storagePath, err := helper.GetStorageByName(storageName)
	if err != nil {
		return nil, err
	}

	return &ObjectPool{storageName, storagePath, relativePath}, nil
}

// GetGitAlternateObjectDirectories for object pools are empty, given pools are
// never a member of another pool, nor do they share Alternate objects with other
// repositories which the pool doesn't contain itself
func (o *ObjectPool) GetGitAlternateObjectDirectories() []string {
	return []string{}
}

// GetGitObjectDirectory satisfies the repository.GitRepo interface, but is not
// used for ObjectPools
func (o *ObjectPool) GetGitObjectDirectory() string {
	return ""
}

// Exists will return true if the pool path exists and is a directory
func (o *ObjectPool) Exists() bool {
	fi, err := os.Stat(o.FullPath())
	if os.IsNotExist(err) || err != nil {
		return false
	}

	return fi.IsDir()
}

// IsValid checks if a repository exists, and if its valid.
func (o *ObjectPool) IsValid() bool {
	if !o.Exists() {
		return false
	}

	return helper.IsGitDirectory(o.FullPath())
}

// Create will create a pool for a repository and pull the required data to this
// pool. `repo` that is passed also joins the repository.
func (o *ObjectPool) Create(ctx context.Context, repo *gitalypb.Repository) (err error) {
	if err := os.MkdirAll(path.Dir(o.FullPath()), 0755); err != nil {
		return err
	}

	if err := o.clone(ctx, repo); err != nil {
		return fmt.Errorf("clone: %v", err)
	}

	if err := o.removeHooksDir(); err != nil {
		return fmt.Errorf("remove hooks: %v", err)
	}

	if err := o.setConfig(ctx, "gc.auto", "0"); err != nil {
		return fmt.Errorf("config gc.auto: %v", err)
	}

	return nil
}

// Remove will remove the pool, and all its contents without preparing and/or
// updating the repositories depending on this object pool
// Subdirectories will remain to exist, and will never be cleaned up, even when
// these are empty.
func (o *ObjectPool) Remove(ctx context.Context) (err error) {
	return os.RemoveAll(o.FullPath())
}

// Init will intiailize an empty pool repository
// if one already exists, it will do nothing
func (o *ObjectPool) Init(ctx context.Context) (err error) {
	targetDir := o.FullPath()

	if helper.IsGitDirectory(targetDir) {
		return nil
	}

	initArgs := []string{"init", "--bare", targetDir}
	cmd, err := git.CommandWithoutRepo(ctx, initArgs...)
	if err != nil {
		return err
	}

	return cmd.Wait()
}

// FromRepo returns an instance of ObjectPool that the repository points to
func FromRepo(repo *gitalypb.Repository) (*ObjectPool, error) {
	dir, err := getAlternateObjectDir(repo)
	if err != nil {
		return nil, err
	}

	if dir == "" {
		return nil, nil
	}

	altPathRelativeToStorage, err := objectPathRelativeToStorage(repo, dir)
	if err != nil {
		return nil, err
	}

	return NewObjectPool(repo.GetStorageName(), filepath.Dir(altPathRelativeToStorage))
}

var (
	// ErrInvalidPoolRepository indicates the directory the alternates file points to is not a valid git repository
	ErrInvalidPoolRepository = errors.New("object pool is not a valid git repository")

	// ErrAlternateObjectDirNotExist indicates a repository does not have an alternates file
	ErrAlternateObjectDirNotExist = errors.New("no alternates directory exists")
)

// getAlternateObjectDir returns the entry in the objects/info/attributes file if it exists
// it will only return the first line of the file if there are multiple lines.
func getAlternateObjectDir(repo *gitalypb.Repository) (string, error) {
	altPath, err := git.InfoAlternatesPath(repo)
	if err != nil {
		return "", err
	}

	if _, err = os.Stat(altPath); err != nil {
		if os.IsNotExist(err) {
			return "", ErrAlternateObjectDirNotExist
		}
		return "", err
	}

	altFile, err := os.Open(altPath)
	if err != nil {
		return "", err
	}
	defer altFile.Close()

	r := bufio.NewReader(altFile)
	b, err := r.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("reading alternates file: %v", err)
	}

	if err == nil {
		b = b[:len(b)-1]
	}

	if bytes.HasPrefix(b, []byte("#")) {
		return "", ErrAlternateObjectDirNotExist
	}

	return string(b), nil
}

// objectPathRelativeToStorage takes an object path that's relative to a repository's object directory
// and returns the path relative to the storage path of the repository.
func objectPathRelativeToStorage(repo *gitalypb.Repository, path string) (string, error) {
	repoPath, err := helper.GetPath(repo)
	if err != nil {
		return "", err
	}

	storagePath, err := helper.GetStorageByName(repo.GetStorageName())
	if err != nil {
		return "", err
	}
	objectDirPath := filepath.Join(repoPath, "objects")

	poolObjectDirFullPath := filepath.Join(objectDirPath, path)

	if !helper.IsGitDirectory(filepath.Dir(poolObjectDirFullPath)) {
		return "", ErrInvalidPoolRepository
	}

	poolRelPath, err := filepath.Rel(storagePath, poolObjectDirFullPath)
	if err != nil {
		return "", err
	}

	return poolRelPath, nil
}
