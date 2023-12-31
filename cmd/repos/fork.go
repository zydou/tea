// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repos

import (
	"fmt"

	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"
	"code.gitea.io/tea/modules/print"

	"code.gitea.io/sdk/gitea"
	"github.com/urfave/cli/v2"
)

// CmdRepoFork represents a sub command of repos to fork an existing repo
var CmdRepoFork = cli.Command{
	Name:        "fork",
	Aliases:     []string{"f"},
	Usage:       "Fork an existing repository",
	Description: "Create a repository from an existing repo",
	ArgsUsage:   " ", // command does not accept arguments
	Action:      runRepoFork,
	Flags: append([]cli.Flag{
		&cli.StringFlag{
			Name:    "owner",
			Aliases: []string{"O"},
			Usage:   "name of fork's owner, defaults to current user",
		},
	}, flags.LoginRepoFlags...),
}

func runRepoFork(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Ensure(context.CtxRequirement{RemoteRepo: true})
	client := ctx.Login.Client()

	opts := gitea.CreateForkOption{}
	if ctx.IsSet("owner") {
		owner := ctx.String("owner")
		opts.Organization = &owner
	}

	repo, _, err := client.CreateFork(ctx.Owner, ctx.Repo, opts)
	if err != nil {
		return err
	}

	topics, _, err := client.ListRepoTopics(repo.Owner.UserName, repo.Name, gitea.ListRepoTopicsOptions{})
	if err != nil {
		return err
	}
	print.RepoDetails(repo, topics)

	fmt.Printf("%s\n", repo.HTMLURL)
	return nil
}
