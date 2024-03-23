package scripts

import (
	"bufio"
	"github.com/ricdotnet/goenvironmental"
	"os"
	"strings"
)

type sitesA map[string]string

type SitesAvailable struct {
}

func Init() *SitesAvailable {
	return &SitesAvailable{}
}

func (sa *SitesAvailable) MapSites() sitesA {
	//log.Info("Entering mapSites")

	apacheDir, err := goenvironmental.Get("APACHE_DIR")
	if err != nil {
		//sa.logger.Fatal("Could not read the APACHE_DIR env variable")
		return nil
	}

	sitesAvailable := make(sitesA)

	files, _ := os.ReadDir(apacheDir + "/sites-available")
	for _, file := range files {
		if file.Name() != ".DS_Store" {
			sa.parseFile(apacheDir+"/sites-available/", file.Name(), sitesAvailable)
		}
	}

	//sa.logger.Infof("Finished parsing and mapping a total of %d files", len(files))
	return sitesAvailable
}

func (sa *SitesAvailable) parseFile(path string, fileName string, sitesA sitesA) {
	//sa.logger.Infof("Started parsing %s", fileName)

	lines, _ := os.Open(path + fileName)
	defer lines.Close()

	reader := bufio.NewScanner(lines)
	reader.Split(bufio.ScanLines)
	for reader.Scan() {
		pair := strings.Split(strings.TrimSpace(reader.Text()), " ")
		if pair[0] == "ServerName" {
			//sa.logger.Infof("Mapping %s to %s\n", fileName, pair[1])
			sitesA[fileName] = pair[1]
			break
		}
	}

	//sa.logger.Infof("Finished parsing %s", fileName)
}
