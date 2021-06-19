package main

import (
	"fmt"
	"github.com/leechichuang/lctool/exec"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
)

const (
	dirFlag = "dir"
)

var (
	commands = []*cli.Command{
		{
			Name:  "init",
			Usage: "init a leetcode note project",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "d",
					Usage: "the dir",
				},
			},
			Action: exec.InitProject,
		},
		{
			Name:  "generate",
			Usage: "generate code files",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "d",
					Usage: "the dir",
				},
				&cli.StringFlag{
					Name:  "q",
					Usage: "the name of question",
				},
				&cli.BoolFlag{
					Name:  "n",
					Usage: "the note of question",
				},
				&cli.BoolFlag{
					Name:  "r",
					Usage: "repeat",
				},
			},
			Action: exec.GenerateCmd,
		},
	}
	BuildVersion = "1.1"
)

func main() {
	app := cli.NewApp()
	app.Usage = "a cli tool for leetcode"
	app.Version = fmt.Sprintf("%s %s/%s", BuildVersion, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	// cli already print error messages
	if err := app.Run(os.Args); err != nil {
		fmt.Println("error:", err)
	}
}
