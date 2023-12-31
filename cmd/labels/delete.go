// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package labels

import (
	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"

	"github.com/urfave/cli/v2"
)

// CmdLabelDelete represents a sub command of labels to delete label.
var CmdLabelDelete = cli.Command{
	Name:        "delete",
	Aliases:     []string{"rm"},
	Usage:       "Delete a label",
	Description: `Delete a label`,
	ArgsUsage:   " ", // command does not accept arguments
	Action:      runLabelDelete,
	Flags: append([]cli.Flag{
		&cli.IntFlag{
			Name:  "id",
			Usage: "label id",
		},
	}, flags.AllDefaultFlags...),
}

func runLabelDelete(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Ensure(context.CtxRequirement{RemoteRepo: true})

	_, err := ctx.Login.Client().DeleteLabel(ctx.Owner, ctx.Repo, ctx.Int64("id"))
	return err
}
