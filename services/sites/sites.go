package sites

import (
	"os"
)

// GetAll
// dir would refer to what dir to look into: available or enabled
// sites-available or sites-enabled
func GetAll(dir string) ([]string, error) {
	dirContent, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	files := make([]string, len(dirContent))
	for i, file := range dirContent {
		files[i] = file.Name()
	}

	return files, nil
}

// ReadSingle
// reads the vhosts of a site and returns the []byte array to be sent on the response to the client
func ReadSingle(dir string, name string) ([]byte, error) {
	vhost, err := os.ReadFile(dir + name)
	if err != nil {
		return nil, err
	}

	return vhost, nil
}

func WriteSingle(dir string, name string) {
	// not implemented yet
}

func DeleteSingle(dir string, name string) {
	// not implemented yet
}
