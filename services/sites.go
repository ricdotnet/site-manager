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
		log.Errorf("Failed to read a vhosts: %s", name)
		return nil, err
	}

	log.Info("Exited SitesService readSingle")
	return conf, nil
}

// TODO: Update this to write based on site_data
func (ss *SitesService) WriteSingle(name string, content string) error {
	apacheDir, _ := goenvironmental.Get("APACHE_DIR")
	apacheDir += "sites-available/"

	err := os.WriteFile(apacheDir+name, []byte(content), 0666)
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

	if !utils.IsValidFilename(new) {
		log.Errorf("The filename used %s is not valid", new)
		return fmt.Errorf("the filename used %s is not a valid filename", new)
	}

	nginxDir, _ := goenvironmental.Get("SITES_AVAILABLE_PATH")

	oldPath := filepath.Join(nginxDir, curr)
	newPath := filepath.Join(nginxDir, new)

	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}

	return nil
}
