// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package task

import (
	"fmt"
	"regexp"
	"strings"

	"code.gitea.io/sdk/gitea"
	"code.gitea.io/tea/modules/config"
	"code.gitea.io/tea/modules/context"
	local_git "code.gitea.io/tea/modules/git"
	"code.gitea.io/tea/modules/print"
	"code.gitea.io/tea/modules/utils"
)

var (
	spaceRegex  = regexp.MustCompile(`[\s_-]+`)
	noSpace     = regexp.MustCompile(`^[^a-zA-Z\s]*`)
	consecutive = regexp.MustCompile(`[\s]{2,}`)
)

// CreatePull creates a PR in the given repo and prints the result
func CreatePull(ctx *context.TeaContext, base, head string, allowMaintainerEdits bool, opts *gitea.CreateIssueOption) (err error) {
	// default is default branch
	if len(base) == 0 {
		base, err = GetDefaultPRBase(ctx.Login, ctx.Owner, ctx.Repo)
		if err != nil {
			return err
		}
	}

	// default is current one
	if len(head) == 0 {
		if ctx.LocalRepo == nil {
			return fmt.Errorf("no local git repo detected, please specify head branch")
		}
		headOwner, headBranch, err := GetDefaultPRHead(ctx.LocalRepo)
		if err != nil {
			return err
		}

		head = GetHeadSpec(headOwner, headBranch, ctx.Owner)
	}

	// head & base may not be the same
	if head == base {
		return fmt.Errorf("can't create PR from %s to %s", head, base)
	}

	// default is head branch name
	if len(opts.Title) == 0 {
		opts.Title = GetDefaultPRTitle(head)
	}
	// title is required
	if len(opts.Title) == 0 {
		return fmt.Errorf("title is required")
	}

	client := ctx.Login.Client()

	pr, _, err := client.CreatePullRequest(ctx.Owner, ctx.Repo, gitea.CreatePullRequestOption{
		Head:      head,
		Base:      base,
		Title:     opts.Title,
		Body:      opts.Body,
		Assignees: opts.Assignees,
		Labels:    opts.Labels,
		Milestone: opts.Milestone,
		Deadline:  opts.Deadline,
	})
	if err != nil {
		return fmt.Errorf("could not create PR from %s to %s:%s: %s", head, ctx.Owner, base, err)
	}

	if pr.AllowMaintainerEdit != allowMaintainerEdits {
		pr, _, err = client.EditPullRequest(ctx.Owner, ctx.Repo, pr.Index, gitea.EditPullRequestOption{
			AllowMaintainerEdit: gitea.OptionalBool(allowMaintainerEdits),
		})
		if err != nil {
			return fmt.Errorf("could not enable maintainer edit on pull: %v", err)
		}
	}

	print.PullDetails(pr, nil, nil)

	fmt.Println(pr.HTMLURL)

	return err
}

// GetDefaultPRBase retrieves the default base branch for the given repo
func GetDefaultPRBase(login *config.Login, owner, repo string) (string, error) {
	meta, _, err := login.Client().GetRepo(owner, repo)
	if err != nil {
		return "", fmt.Errorf("could not fetch repo meta: %s", err)
	}
	return meta.DefaultBranch, nil
}

// GetDefaultPRHead uses the currently checked out branch, tries to find a remote
// that has a branch with the same name, and extracts the owner from its URL.
// If no remote matches, owner is empty, meaning same as head repo owner.
func GetDefaultPRHead(localRepo *local_git.TeaRepo) (owner, branch string, err error) {
	var sha string
	if branch, sha, err = localRepo.TeaGetCurrentBranchNameAndSHA(); err != nil {
		return
	}

	remote, err := localRepo.TeaFindBranchRemote(branch, sha)
	if err != nil {
		err = fmt.Errorf("could not determine remote for current branch: %s", err)
		return
	}

	if remote == nil {
		// if no remote branch is found for the local branch,
		// we leave owner empty, meaning "use same repo as head" to gitea.
		return
	}

	url, err := local_git.ParseURL(remote.Config().URLs[0])
	if err != nil {
		return
	}
	owner, _ = utils.GetOwnerAndRepo(url.Path, "")
	return
}

// GetHeadSpec creates a head string as expected by gitea API
func GetHeadSpec(owner, branch, baseOwner string) string {
	if len(owner) != 0 && owner != baseOwner {
		return fmt.Sprintf("%s:%s", owner, branch)
	}
	return branch
}

// GetDefaultPRTitle transforms a string like a branchname to a readable text
func GetDefaultPRTitle(header string) string {
	// Extract the part after the last colon in the input string
	colonIndex := strings.LastIndex(header, ":")
	if colonIndex != -1 {
		header = header[colonIndex+1:]
	}

	title := noSpace.ReplaceAllString(header, "")
	title = spaceRegex.ReplaceAllString(title, " ")
	title = strings.TrimSpace(title)
	title = strings.Title(strings.ToLower(title))
	title = consecutive.ReplaceAllString(title, " ")

	return title
}
