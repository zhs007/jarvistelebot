package base

import (
	"io/ioutil"
	"os"
	"path"
	"sync"

	"gopkg.in/yaml.v2"
)

// Config - config
type Config struct {
	WebServAddr  string
	LogPath      string
	ErrPath      string
	LogLevel     string
	DTAPI        string
	DTAPPKEY     string
	DTBusinessid string
	runpath      string
	respath      string
}

var cfg Config
var onceCfg sync.Once

// LoadConfig - load config
func load(filename string) error {
	fi, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer fi.Close()
	fd, err1 := ioutil.ReadAll(fi)
	if err1 != nil {
		return err1
	}

	err2 := yaml.Unmarshal(fd, &cfg)
	if err2 != nil {
		return err2
	}

	return nil
}

func procBaseDir(rundir string, resdir string) {
	cfg.LogPath = path.Join(rundir, cfg.LogPath)
	cfg.ErrPath = path.Join(rundir, cfg.ErrPath)

	cfg.runpath = rundir
	cfg.respath = resdir
}

// LoadConfig - load config
func LoadConfig(filename string, rundir string, resdir string) (err error) {
	onceCfg.Do(func() {
		err = load(filename)

		procBaseDir(rundir, resdir)
	})

	return
}

// GetConfig - get Config
func GetConfig() *Config {
	return &cfg
}

// BuildRunPath -
func BuildRunPath(dir string) string {
	return path.Join(cfg.runpath, dir)
}

// BuildResPath -
func BuildResPath(dir string) string {
	return path.Join(cfg.respath, dir)
}
