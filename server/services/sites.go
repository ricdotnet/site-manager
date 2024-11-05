package services

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/ricdotnet/goenvironmental"
	"os"
	"path/filepath"
	"ricr.dev/site-manager/utils"
)

type SitesService struct {
}

func NewSitesService() *SitesService {
	return &SitesService{}
}

func (ss *SitesService) GetAll(dir string) ([]string, error) {
	log.Info("Entered SitesService getAll")

	dirContent, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	files := make([]string, len(dirContent))
	for i, file := range dirContent {
		files[i] = file.Name()
	}

	log.Info("Exited SitesService getAll")
	return files, nil
}

func (ss *SitesService) ReadSingle(dir string, name string) ([]byte, error) {
	log.Info("Entered SitesService readSingle")

	conf, err := os.ReadFile(filepath.Join(dir, name))
	if err != nil {
		log.Errorf("Failed to read an nginx config: %s", name)
		return nil, err
	}

	log.Info("Exited SitesService readSingle")
	return conf, nil
}

func (ss *SitesService) WriteSingle(name string, content string) error {
	nginxDir, _ := goenvironmental.Get("NGINX_PATH")

	if !utils.IsValidFilename(name) {
		log.Errorf("The filename used %s is not valid", name)
		return fmt.Errorf("the filename used %s is not a valid filename", name)
	}

	path := filepath.Join(nginxDir, name)

	log.Infof("Will write the new config in %s", path)

	err := os.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return err
	}

	log.Infof("Finished writing the new config for %s", name)
	return nil
}

func (ss *SitesService) DeleteSingle(dir string, name string) {
	// not implemented yet
}

// UpdateName
// update the name of a .conf file
func (ss *SitesService) UpdateName(curr string, new string) error {
	log.Infof("Updating the name of %s to %s", curr, new)

	if !utils.IsValidFilename(new) {
		log.Errorf("The filename used %s is not valid", new)
		return fmt.Errorf("the filename used %s is not a valid filename", new)
	}

	nginxDir, _ := goenvironmental.Get("NGINX_PATH")

	oldPath := filepath.Join(nginxDir, curr)
	newPath := filepath.Join(nginxDir, new)

	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}

	log.Infof("Finished updating the name of %s to %s", curr, new)
	return nil
}
