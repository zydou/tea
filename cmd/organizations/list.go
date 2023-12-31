// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package organizations

import (
	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"
	"code.gitea.io/tea/modules/print"

	"code.gitea.io/sdk/gitea"
	"github.com/urfave/cli/v2"
)

// CmdOrganizationList represents a sub command of organizations to list users organizations
var CmdOrganizationList = cli.Command{
	Name:        "list",
	Aliases:     []string{"ls"},
	Usage:       "List Organizations",
	Description: "List users organizations",
	ArgsUsage:   " ", // command does not accept arguments
	Action:      RunOrganizationList,
	Flags: append([]cli.Flag{
		&flags.PaginationPageFlag,
		&flags.PaginationLimitFlag,
	}, flags.AllDefaultFlags...),
}

// RunOrganizationList list user organizations
func RunOrganizationList(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	client := ctx.Login.Client()

	userOrganizations, _, err := client.ListUserOrgs(ctx.Login.User, gitea.ListOrgsOptions{
		ListOptions: ctx.GetListOptions(),
	})
	if err != nil {
		return err
	}

	print.OrganizationsList(userOrganizations, ctx.Output)

	return nil
}
