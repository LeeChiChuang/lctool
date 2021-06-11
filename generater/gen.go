package generater

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

var generateParamErr = errors.New("name and number is necessary")

func GenerateCmd(c *cli.Context) error {
	var dir string
	if c.String("d") == "" {
		dir = "."
	} else {
		dir = c.String("d")
	}

	if c.String("m") == "" {
		return generateParamErr
	}

	var name string
	if c.String("m") != "" {
		name = c.String("m")
	}
	lc, err := NewLeetCode(name)
	if err != nil {
		return err
	}

	err = lc.WriteDesc(dir)
	if err != nil {
		return err
	}
	err = lc.WriteCode(dir)
	if err != nil {
		return err
	}
	fmt.Println("success")

	return nil
}
