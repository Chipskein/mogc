package repositories

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

type File struct {
	Name         string
	Size         int64
	FullPath     string
	Extension    string
	IsADirectory bool
}
type LocalRepository struct {
	DEFAULT_DIRECTORY string
	CURRENT_DIRECTORY string
	Files             map[string]File
}

func (r *LocalRepository) ListFiles() error {
	files, err := r.ReadDirectoryOrDefault()
	if err != nil {
		return err
	}
	filelist := make(map[string]File)
	for _, file := range files {
		var listedfile File
		listedfile.Name = file.Name()
		listedfile.FullPath = fmt.Sprintf("%s/%s", r.CURRENT_DIRECTORY, listedfile.Name)
		listedfile.Size = file.Size()
		listedfile.IsADirectory = file.IsDir()
		listedfile.Extension = filepath.Ext(listedfile.FullPath)

		filelist[listedfile.Name] = listedfile
	}
	r.Files = filelist
	return nil
}
func (r *LocalRepository) ReadDirectoryOrDefault() ([]fs.FileInfo, error) {
	absolute, _ := filepath.Abs(r.CURRENT_DIRECTORY)
	r.CURRENT_DIRECTORY = absolute
	default_absolute, _ := filepath.Abs(r.DEFAULT_DIRECTORY)
	r.DEFAULT_DIRECTORY = default_absolute

	files, err := ioutil.ReadDir(r.CURRENT_DIRECTORY)
	if err != nil {
		r.CURRENT_DIRECTORY = r.DEFAULT_DIRECTORY
	}
	files, err = ioutil.ReadDir(r.CURRENT_DIRECTORY)
	if err != nil {
		return nil, err
	}

	return files, nil
}
func IsExtesionSupported(extension string) bool {
	switch extension {
	case ".ogg":
		return true
	case ".wav":
		return true
	case ".mp3":
		return true
	case ".flac":
		return true
	default:
		return false
	}
}
