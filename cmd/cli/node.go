package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func listNodes(c *cli.Context) error {
	// Get values from flags
	target := "active"
	if c.Bool("all") {
		target = "all"
	}
	if c.Bool("inactive") {
		target = "inactive"
	}
	nodes, err := nodesmgr.Gets(target)
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Hostname",
		"UUID",
		"Platform",
		"Context",
		"Last Status",
		"Last Result",
	})
	if len(nodes) > 0 {
		data := [][]string{}
		fmt.Printf("Existing %s nodes (%d):\n", target, len(nodes))
		for _, n := range nodes {
			_n := []string{
				n.Hostname,
				n.UUID,
				n.Platform,
				n.Context,
				pastTimeAgo(n.LastStatus),
				pastTimeAgo(n.LastResult),
			}
			data = append(data, _n)
		}
		table.AppendBulk(data)
		table.Render()
	} else {
		fmt.Printf("No nodes\n")
	}
	return nil
}