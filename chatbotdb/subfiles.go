package chatbotdb

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Subfiles - subfiles
type Subfiles struct {
	Subfiles []string
}

// LoadSubfiles - load subfiles
func LoadSubfiles(filename string) (*Subfiles, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	sf := &Subfiles{}
	err = yaml.Unmarshal(fd, sf)
	if err != nil {
		return nil, err
	}

	return sf, nil
}

// LoadSubfilesFromBuff - load subfiles from buffer
func LoadSubfilesFromBuff(fd []byte) (*Subfiles, error) {
	sf := &Subfiles{}
	err := yaml.Unmarshal(fd, sf)
	if err != nil {
		return nil, err
	}

	return sf, nil
}
