// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"

	"code.gitea.io/tea/cmd/labels"
	"github.com/urfave/cli/v2"
)

// CmdLabels represents to operate repositories' labels.
var CmdLabels = cli.Command{
	Name:        "labels",
	Aliases:     []string{"label"},
	Category:    catEntities,
	Usage:       "Manage issue labels",
	Description: `Manage issue labels`,
	ArgsUsage:   " ", // command does not accept arguments
	Action:      runLabels,
	Subcommands: []*cli.Command{
		&labels.CmdLabelsList,
		&labels.CmdLabelCreate,
		&labels.CmdLabelUpdate,
		&labels.CmdLabelDelete,
	},
	Flags: labels.CmdLabelsList.Flags,
}

func runLabels(ctx *cli.Context) error {
	if ctx.Args().Len() == 1 {
		return runLabelsDetails(ctx)
	}
	return labels.RunLabelsList(ctx)
}

func runLabelsDetails(ctx *cli.Context) error {
	return fmt.Errorf("Not yet implemented")
}
