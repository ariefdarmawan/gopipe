package gopipe

import (
	"github.com/eaciit/toolkit"
)

type CommandEnum int

const (
	CommandDirect CommandEnum = 0
	CommandFile               = 1
	CommandLine               = 2
	CommandUri                = 3
)

type ICommand interface {
	Exec(toolkit.M) *toolkit.Result
}

type Command struct {
	CommandType CommandEnum
	Txt         string

	Libs []string
}

func NewCommand(cmdtype CommandEnum, txt string) *Command {
	cmd := new(Command)
	cmd.CommandType = cmdtype
	cmd.Txt = txt
	return cmd
}

func (c *Command) AddLib(libs ...string) *Command {
	for _, lib := range libs {
		c.Libs = append(c.Libs, lib)
	}
	return c
}

func (c *Command) RemoveLib(libs ...string) *Command {
	tmpLib := []string{}
	for _, lib := range libs {
		for _, currentLib := range c.Libs {
			if currentLib != lib {
				tmpLib = append(tmpLib, currentLib)
			}
		}
	}
	c.Libs = tmpLib
	return c
}

func (c *Command) Exec(in toolkit.M) *toolkit.Result {
	r := toolkit.NewResult()
	return r
}
