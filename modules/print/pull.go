// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package print

import (
	"fmt"
	"strings"

	"code.gitea.io/sdk/gitea"
)

var ciStatusSymbols = map[gitea.StatusState]string{
	gitea.StatusSuccess: "✓ ",
	gitea.StatusPending: "⭮ ",
	gitea.StatusWarning: "⚠ ",
	gitea.StatusError:   "✘ ",
	gitea.StatusFailure: "❌ ",
}

// PullDetails print an pull rendered to stdout
func PullDetails(pr *gitea.PullRequest, reviews []*gitea.PullReview, ciStatus *gitea.CombinedStatus) {
	base := pr.Base.Name
	head := formatPRHead(pr)
	state := formatPRState(pr)

	out := fmt.Sprintf(
		"# #%d %s (%s)\n@%s created %s\t**%s** <- **%s**\n\n%s\n\n",
		pr.Index,
		pr.Title,
		state,
		pr.Poster.UserName,
		FormatTime(*pr.Created, false),
		base,
		head,
		pr.Body,
	)

	if ciStatus != nil || len(reviews) != 0 || pr.State == gitea.StateOpen {
		out += "---\n"
	}

	out += formatReviews(pr, reviews)

	if ciStatus != nil {
		var summary, errors string
		for _, s := range ciStatus.Statuses {
			summary += ciStatusSymbols[s.State]
			if s.State != gitea.StatusSuccess {
				errors += fmt.Sprintf("  - [**%s**:\t%s](%s)\n", s.Context, s.Description, s.TargetURL)
			}
		}
		if len(ciStatus.Statuses) != 0 {
			out += fmt.Sprintf("- CI: %s\n%s", summary, errors)
		}
	}

	if pr.State == gitea.StateOpen {
		if pr.Mergeable {
			out += "- No Conflicts\n"
		} else {
			out += "- **Conflicting files**\n"
		}
	}

	if pr.AllowMaintainerEdit {
		out += "- Maintainers are allowed to edit\n"
	}

	outputMarkdown(out, getRepoURL(pr.HTMLURL))
}

func formatPRHead(pr *gitea.PullRequest) string {
	head := pr.Head.Name
	if pr.Head.RepoID != pr.Base.RepoID {
		if pr.Head.Repository != nil {
			head = pr.Head.Repository.Owner.UserName + ":" + head
		} else {
			head = "delete:" + head
		}
	}
	return head
}

func formatPRState(pr *gitea.PullRequest) string {
	if pr.Merged != nil {
		return "merged"
	}
	return string(pr.State)
}

func formatReviews(pr *gitea.PullRequest, reviews []*gitea.PullReview) string {
	result := ""
	if len(reviews) == 0 {
		return result
	}

	// deduplicate reviews by user (via review time & userID),
	reviewByUserOrTeam := make(map[string]*gitea.PullReview)
	for _, review := range reviews {
		switch review.State {
		case gitea.ReviewStateApproved,
			gitea.ReviewStateRequestChanges,
			gitea.ReviewStateRequestReview:
			if review.Reviewer != nil {
				if r, ok := reviewByUserOrTeam[fmt.Sprintf("user_%d", review.Reviewer.ID)]; !ok || review.Submitted.After(r.Submitted) {
					reviewByUserOrTeam[fmt.Sprintf("user_%d", review.Reviewer.ID)] = review
				}
			} else if review.ReviewerTeam != nil {
				if r, ok := reviewByUserOrTeam[fmt.Sprintf("team_%d", review.ReviewerTeam.ID)]; !ok || review.Submitted.After(r.Submitted) {
					reviewByUserOrTeam[fmt.Sprintf("team_%d", review.ReviewerTeam.ID)] = review
				}
			}

		}
	}

	// group reviews by type
	reviewByState := make(map[gitea.ReviewStateType][]string)
	for _, r := range reviewByUserOrTeam {
		if r.Reviewer != nil {
			reviewByState[r.State] = append(reviewByState[r.State],
				r.Reviewer.UserName,
			)
		} else if r.ReviewerTeam != nil {
			// only pulls to orgs can have team reviews
			org := pr.Base.Repository.Owner
			reviewByState[r.State] = append(reviewByState[r.State],
				fmt.Sprintf("%s/%s", org.UserName, r.ReviewerTeam.Name),
			)
		}
	}

	// stringify
	for state, user := range reviewByState {
		result += fmt.Sprintf("- %s by @%s\n", state, strings.Join(user, ", @"))
	}
	return result
}

// PullsList prints a listing of pulls
func PullsList(prs []*gitea.PullRequest, output string, fields []string) {
	printPulls(prs, output, fields)
}

// PullFields are all available fields to print with PullsList()
var PullFields = []string{
	"index",
	"state",
	"author",
	"author-id",
	"url",

	"title",
	"body",

	"mergeable",
	"base",
	"base-commit",
	"head",
	"diff",
	"patch",

	"created",
	"updated",
	"deadline",

	"assignees",
	"milestone",
	"labels",
	"comments",
}

func printPulls(pulls []*gitea.PullRequest, output string, fields []string) {
	labelMap := map[int64]string{}
	var printables = make([]printable, len(pulls))
	machineReadable := isMachineReadable(output)

	for i, x := range pulls {
		// pre-serialize labels for performance
		for _, label := range x.Labels {
			if _, ok := labelMap[label.ID]; !ok {
				labelMap[label.ID] = formatLabel(label, !machineReadable, "")
			}
		}
		// store items with printable interface
		printables[i] = &printablePull{x, &labelMap}
	}

	t := tableFromItems(fields, printables, machineReadable)
	t.print(output)
}

type printablePull struct {
	*gitea.PullRequest
	formattedLabels *map[int64]string
}

func (x printablePull) FormatField(field string, machineReadable bool) string {
	switch field {
	case "index":
		return fmt.Sprintf("%d", x.Index)
	case "state":
		return formatPRState(x.PullRequest)
	case "author":
		return formatUserName(x.Poster)
	case "author-id":
		return x.Poster.UserName
	case "url":
		return x.HTMLURL
	case "title":
		return x.Title
	case "body":
		return x.Body
	case "created":
		return FormatTime(*x.Created, machineReadable)
	case "updated":
		return FormatTime(*x.Updated, machineReadable)
	case "deadline":
		if x.Deadline == nil {
			return ""
		}
		return FormatTime(*x.Deadline, machineReadable)
	case "milestone":
		if x.Milestone != nil {
			return x.Milestone.Title
		}
		return ""
	case "labels":
		var labels = make([]string, len(x.Labels))
		for i, l := range x.Labels {
			labels[i] = (*x.formattedLabels)[l.ID]
		}
		return strings.Join(labels, " ")
	case "assignees":
		var assignees = make([]string, len(x.Assignees))
		for i, a := range x.Assignees {
			assignees[i] = formatUserName(a)
		}
		return strings.Join(assignees, " ")
	case "comments":
		return fmt.Sprintf("%d", x.Comments)
	case "mergeable":
		isMergeable := x.Mergeable && x.State == gitea.StateOpen
		return formatBoolean(isMergeable, !machineReadable)
	case "base":
		return x.Base.Ref
	case "base-commit":
		return x.MergeBase
	case "head":
		return formatPRHead(x.PullRequest)
	case "diff":
		return x.DiffURL
	case "patch":
		return x.PatchURL
	}
	return ""
}
