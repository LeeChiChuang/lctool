package exec

import (
	"errors"
	"fmt"
	"github.com/leechichuang/lctool/question"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
)

var generateParamErr = errors.New("name and number is necessary")

func GenerateCmd(c *cli.Context) error {
	var dir = c.String("d")
	if len(dir) == 0 {
		dir = "."
	}

	var q = c.String("q")
	if len(q) == 0 {
		return generateParamErr
	}

	var note = c.Bool("n")
	var repeat = c.Bool("r")

	err := doGenerateCmd(dir, q, note, repeat)
	if err != nil {
		return err
	}
	fmt.Println(aurora.Green("Done."))

	return nil
}

func doGenerateCmd(dir string, q string, note bool, repeat bool) error {
	fmt.Println(dir, q, note, repeat)
	lc, err := question.NewLeetCode(q)
	if err != nil {
		return err
	}

	if note {
		err = GenerateNote(lc)
		if err != nil {
			return err
		}

		err = GenerateRepeat(lc)
		if err != nil {
			return err
		}

		err = GenerateReadme(lc)
		if err != nil {
			return err
		}
	}

	err = lc.WriteDesc(dir)
	if err != nil {
		return err
	}

	err = lc.WriteCode(dir, repeat)
	if err != nil {
		return err
	}

	return nil
}
