package main

import "strings"

type cmdAfterSplit struct {
	cmd []string
}

func newResultsAfterSplit() *cmdAfterSplit {
	return &cmdAfterSplit{
		cmd: nil,
	}
}

func (c *cmdAfterSplit) getCommand(command string) (cmd []string) {
	c.cmd = strings.Fields(command)
	return c.cmd
}

func (c *cmdAfterSplit) clearCommand() {
	c.cmd = nil
}
