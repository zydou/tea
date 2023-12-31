// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package milestones

import (
	"code.gitea.io/tea/cmd/flags"

	"github.com/urfave/cli/v2"
)

// CmdMilestonesClose represents a sub command of milestones to close an milestone
var CmdMilestonesClose = cli.Command{
	Name:        "close",
	Usage:       "Change state of one or more milestones to 'closed'",
	Description: `Change state of one or more milestones to 'closed'`,
	ArgsUsage:   "<milestone name> [<milestone name>...]",
	Action: func(ctx *cli.Context) error {
		if ctx.Bool("force") {
			return deleteMilestone(ctx)
		}
		return editMilestoneStatus(ctx, true)
	},
	Flags: append([]cli.Flag{
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "delete milestone",
		},
	}, flags.AllDefaultFlags...),
}
