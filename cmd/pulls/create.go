// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pulls

import (
	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"
	"code.gitea.io/tea/modules/interact"
	"code.gitea.io/tea/modules/task"

	"github.com/urfave/cli/v2"
)

// CmdPullsCreate creates a pull request
var CmdPullsCreate = cli.Command{
	Name:        "create",
	Aliases:     []string{"c"},
	Usage:       "Create a pull-request",
	Description: "Create a pull-request in the current repo",
	Action:      runPullsCreate,
	Flags: append([]cli.Flag{
		&cli.StringFlag{
			Name:  "head",
			Usage: "Branch name of the PR source (default is current one). To specify a different head repo, use <user>:<branch>",
		},
		&cli.StringFlag{
			Name:    "base",
			Aliases: []string{"b"},
			Usage:   "Branch name of the PR target (default is repos default branch)",
		},
	}, flags.IssuePREditFlags...),
}

func runPullsCreate(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)

	// no args -> interactive mode
	if ctx.NumFlags() == 0 {
		return interact.CreatePull(ctx)
	}

	// else use args to create PR
	opts, err := flags.GetIssuePREditFlags(ctx)
	if err != nil {
		return err
	}

	return task.CreatePull(
		ctx,
		ctx.String("base"),
		ctx.String("head"),
		opts,
	)
}
