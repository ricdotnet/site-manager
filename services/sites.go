package services

import (
	"github.com/op/go-logging"
	"os"
	"ricr.dev/site-manager/config"
)

type SitesService struct {
	logger *logging.Logger
}

func NewSitesService(cfg *config.Config) *SitesService {
	return &SitesService{
		logger: cfg.Logger,
	}
}

// GetAll
// dir would refer to what dir to look into: available or enabled
// sites-available or sites-enabled
func (ss *SitesService) GetAll(dir string) ([]string, error) {
	ss.logger.Info("Entered SitesService getAll")

	dirContent, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	files := make([]string, len(dirContent))
	for i, file := range dirContent {
		files[i] = file.Name()
	}

	ss.logger.Info("Exited SitesService getAll")
	return files, nil
}

// ReadSingle
// reads the vhosts of a site and returns the []byte array to be sent on the response to the client
func (ss *SitesService) ReadSingle(dir string, name string) ([]byte, error) {
	ss.logger.Info("Entered SitesService readSingle")

	vhost, err := os.ReadFile(dir + name)
	if err != nil {
		ss.logger.Errorf("Failed to read a vhosts: %s", name)
		return nil, err
	}

	ss.logger.Info("Exited SitesService readSingle")
	return vhost, nil
}

func (ss *SitesService) WriteSingle(dir string, name string) {
	// not implemented yet
}

func (ss *SitesService) DeleteSingle(dir string, name string) {
	// not implemented yet
}
