// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package organizations

import (
	"fmt"

	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"
	"github.com/urfave/cli/v2"
)

// CmdOrganizationDelete represents a sub command of organizations to delete a given user organization
var CmdOrganizationDelete = cli.Command{
	Name:        "delete",
	Aliases:     []string{"rm"},
	Usage:       "Delete users Organizations",
	Description: "Delete users organizations",
	ArgsUsage:   "<organization name>",
	Action:      RunOrganizationDelete,
	Flags: []cli.Flag{
		&flags.LoginFlag,
		&flags.RemoteFlag,
	},
}

// RunOrganizationDelete delete user organization
func RunOrganizationDelete(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)

	client := ctx.Login.Client()

	if ctx.Args().Len() < 1 {
		return fmt.Errorf("You have to specify the organization name you want to delete")
	}

	response, err := client.DeleteOrg(ctx.Args().First())
	if response != nil && response.StatusCode == 404 {
		return fmt.Errorf("The given organization does not exist")
	}

	return err
}
