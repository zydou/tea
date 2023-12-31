// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package milestones

import (
	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"

	"github.com/urfave/cli/v2"
)

// CmdMilestonesDelete represents a sub command of milestones to delete an milestone
var CmdMilestonesDelete = cli.Command{
	Name:        "delete",
	Aliases:     []string{"rm"},
	Usage:       "delete a milestone",
	Description: "delete a milestone",
	ArgsUsage:   "<milestone name>",
	Action:      deleteMilestone,
	Flags:       flags.AllDefaultFlags,
}

func deleteMilestone(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Ensure(context.CtxRequirement{RemoteRepo: true})
	client := ctx.Login.Client()

	_, err := client.DeleteMilestoneByName(ctx.Owner, ctx.Repo, ctx.Args().First())
	return err
}
