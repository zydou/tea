// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package notifications

import (
	"log"

	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"
	"code.gitea.io/tea/modules/print"

	"code.gitea.io/sdk/gitea"
	"github.com/urfave/cli/v2"
)

var notifyFieldsFlag = flags.FieldsFlag(print.NotificationFields, []string{
	"id", "status", "index", "type", "state", "title",
})

var notifyTypeFlag = flags.NewCsvFlag("types", "subject types to filter by", []string{"t"},
	[]string{"issue", "pull", "repository", "commit"}, nil)

// CmdNotificationsList represents a sub command of notifications to list notifications
var CmdNotificationsList = cli.Command{
	Name:        "ls",
	Aliases:     []string{"list"},
	Usage:       "List notifications",
	Description: `List notifications`,
	ArgsUsage:   " ", // command does not accept arguments
	Action:      RunNotificationsList,
	Flags: append([]cli.Flag{
		notifyFieldsFlag,
		notifyTypeFlag,
	}, flags.NotificationFlags...),
}

// RunNotificationsList list notifications
func RunNotificationsList(ctx *cli.Context) error {
	var states []gitea.NotifyStatus
	statesStr, err := flags.NotificationStateFlag.GetValues(ctx)
	if err != nil {
		return err
	}
	for _, s := range statesStr {
		states = append(states, gitea.NotifyStatus(s))
	}

	var types []gitea.NotifySubjectType
	typesStr, err := notifyTypeFlag.GetValues(ctx)
	if err != nil {
		return err
	}
	for _, t := range typesStr {
		types = append(types, gitea.NotifySubjectType(t))
	}

	return listNotifications(ctx, states, types)
}

// listNotifications will get the notifications based on status and subject type
func listNotifications(cmd *cli.Context, status []gitea.NotifyStatus, subjects []gitea.NotifySubjectType) error {
	var news []*gitea.NotificationThread
	var err error

	ctx := context.InitCommand(cmd)
	client := ctx.Login.Client()
	all := ctx.Bool("mine")

	// This enforces pagination (see https://github.com/go-gitea/gitea/issues/16733)
	listOpts := ctx.GetListOptions()
	if listOpts.Page == 0 {
		listOpts.Page = 1
	}

	fields, err := notifyFieldsFlag.GetValues(cmd)
	if err != nil {
		return err
	}

	if all {
		// add repository to the default fields
		if !cmd.IsSet("fields") {
			fields = append(fields, "repository")
		}

		news, _, err = client.ListNotifications(gitea.ListNotificationOptions{
			ListOptions:  listOpts,
			Status:       status,
			SubjectTypes: subjects,
		})
	} else {
		ctx.Ensure(context.CtxRequirement{RemoteRepo: true})
		news, _, err = client.ListRepoNotifications(ctx.Owner, ctx.Repo, gitea.ListNotificationOptions{
			ListOptions:  listOpts,
			Status:       status,
			SubjectTypes: subjects,
		})
	}
	if err != nil {
		log.Fatal(err)
	}

	print.NotificationsList(news, ctx.Output, fields)
	return nil
}
