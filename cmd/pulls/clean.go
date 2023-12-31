// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package pulls

import (
	"fmt"

	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"
	"code.gitea.io/tea/modules/interact"
	"code.gitea.io/tea/modules/task"
	"code.gitea.io/tea/modules/utils"

	"github.com/urfave/cli/v2"
)

// CmdPullsClean removes the remote and local feature branches, if a PR is merged.
var CmdPullsClean = cli.Command{
	Name:        "clean",
	Usage:       "Deletes local & remote feature-branches for a closed pull request",
	Description: `Deletes local & remote feature-branches for a closed pull request`,
	ArgsUsage:   "<pull index>",
	Action:      runPullsClean,
	Flags: append([]cli.Flag{
		&cli.BoolFlag{
			Name:  "ignore-sha",
			Usage: "Find the local branch by name instead of commit hash (less precise)",
		},
	}, flags.AllDefaultFlags...),
}

func runPullsClean(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Ensure(context.CtxRequirement{LocalRepo: true})
	if ctx.Args().Len() != 1 {
		return fmt.Errorf("Must specify a PR index")
	}

	idx, err := utils.ArgToIndex(ctx.Args().First())
	if err != nil {
		return err
	}

	return task.PullClean(ctx.Login, ctx.Owner, ctx.Repo, idx, ctx.Bool("ignore-sha"), interact.PromptPassword)
}
