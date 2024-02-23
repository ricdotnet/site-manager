package services

import (
	"fmt"
	"github.com/ricdotnet/goenvironmental"
	"os"
	"path/filepath"
	"ricr.dev/site-manager/utils"
)

type SitesService struct {
	//logger *logging.Logger
}

func NewSitesService() *SitesService {
	return &SitesService{}
}

// GetAll
// dir would refer to what dir to look into: available or enabled
// sites-available or sites-enabled
func (ss *SitesService) GetAll(dir string) ([]string, error) {
	//LOGGER ss.logger.Info("Entered SitesService getAll")

	dirContent, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	files := make([]string, len(dirContent))
	for i, file := range dirContent {
		files[i] = file.Name()
	}

	//LOGGER ss.logger.Info("Exited SitesService getAll")
	return files, nil
}

// ReadSingle
// reads the vhosts of a site and returns the []byte array to be sent on the response to the client
func (ss *SitesService) ReadSingle(dir string, name string) ([]byte, error) {
	//LOGGER ss.logger.Info("Entered SitesService readSingle")

	vhost, err := os.ReadFile(dir + name)
	if err != nil {
		//LOGGER ss.logger.Errorf("Failed to read a vhosts: %s", name)
		return nil, err
	}

	//LOGGER ss.logger.Info("Exited SitesService readSingle")
	return vhost, nil
}

// TODO: Update this to write based on site_data
func (ss *SitesService) WriteSingle(name string, content string) error {
	apacheDir, _ := goenvironmental.Get("APACHE_DIR")
	apacheDir += "sites-available/"

	err := os.WriteFile(apacheDir+name, []byte(content), 0666)
	if err != nil {
		return err
	}

	//LOGGER ss.logger.Infof("Finished writing the new config for %s", name)
	return nil
}

func (ss *SitesService) DeleteSingle(dir string, name string) {
	// not implemented yet
}

// UpdateName
// update the name of a .conf file
func (ss *SitesService) UpdateName(curr string, new string) error {

	if !utils.IsValidFilename(new) {
		//LOGGER ss.logger.Errorf("The filename used %s is not a valid filename", new)
		return fmt.Errorf("the filename used %s is not a valid filename", new)
	}

	apachePath := utils.BuildApachePath("sites-available/")

	oldPath := filepath.Join(apachePath, curr)
	newPath := filepath.Join(apachePath, new)

	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}

	return nil
}
