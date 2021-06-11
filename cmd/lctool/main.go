package main

import (
	"fmt"
	"github.com/lcc321/lctools/generater"
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
			Name:  "generate",
			Usage: "generate code files",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "d",
					Usage: "the dir",
				},
				&cli.StringFlag{
					Name:  "m",
					Usage: "the name of question",
				},
				&cli.StringFlag{
					Name:  "n",
					Usage: "the number of question",
				},
			},
			Action: generater.GenerateCmd,
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
