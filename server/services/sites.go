package services

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/ricdotnet/goenvironmental"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/utils"
	"slices"
	"strings"
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

func (ss *SitesService) DeleteSingle(name string) error {
	nginxDir, _ := goenvironmental.Get("NGINX_PATH")

	return os.Remove(filepath.Join(nginxDir, name))
}

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

func (ss *SitesService) FindFileOnly(sites *[]models.Site) error {
	nginxPath, _ := goenvironmental.Get("NGINX_PATH")

	files, err := os.ReadDir(nginxPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		contains := slices.ContainsFunc(*sites, func(site models.Site) bool {
			return site.ConfigName == file.Name()
		})
		if !contains && strings.HasSuffix(file.Name(), ".conf") {
			*sites = append(*sites, models.Site{
				ConfigName: file.Name(),
				ConfigOnly: true,
			})
		}
	}

	return nil
}

type CertificatesResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Failed  bool   `json:"failed"`
}

func (ss *SitesService) GetCertificates() error {
	commandServiceUrl, _ := goenvironmental.Get("APP_COMMAND_SERVICE_URL")

	req, _ := http.NewRequest("POST", commandServiceUrl+"certificates", nil)
	req.Header.Set("Authorization", "some-secret-key")

	client := &http.Client{}
	response, _ := client.Do(req)

	responseBuf, _ := io.ReadAll(response.Body)
	//responseStr := string(responseBuf)

	var result CertificatesResponse
	_ = json.Unmarshal(responseBuf, &result)

	certificateParts := strings.Split(result.Message, "\n")

	for _, certificatePart := range certificateParts {
		println(certificatePart)
	}

	return nil
}
