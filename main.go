package main

import (
	"flag"
	"path"
	"runtime"

	"github.com/zhs007/jarvistelebot/base"
)

func main() {
	runtime.GOMAXPROCS(1)

	var rundir string
	var resdir string
	flag.StringVar(&rundir, "run", "./", "run path")
	flag.StringVar(&resdir, "res", "./", "res path")
	flag.Parse()

	base.LoadConfig(path.Join(rundir, "./config.yaml"), rundir, resdir)
	base.InitLogger()
	base.Info("restart...")
	cfg := base.GetConfig()

	StartServ(cfg.WebServAddr)
}
