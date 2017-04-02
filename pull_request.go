package main

import (
	"context"
	"log"

	"github.com/google/go-github/github"
)

// PullRequestGetter queries github for a specific pull request
type PullRequestGetter interface {
	Get(context.Context, string, string, int) (*github.PullRequest, *github.Response, error)
}

// PullRequestLister queries github for all pull requests
type PullRequestLister interface {
	List(context.Context, string, string, *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error)
}

// IssueGetter queries github for a specific issue
type IssueGetter interface {
	Get(context.Context, string, string, int) (*github.Issue, *github.Response, error)
}

type StatusGetter interface {
	GetCombinedStatus(context.Context, string, string, string, *github.ListOptions) (*github.CombinedStatus, *github.Response, error)
}

// processPullRequest filters out non-mergeable pull requests
func processPullRequest(issueClient IssueGetter, statusClient StatusGetter, mergeLabel string, input <-chan *github.PullRequest) <-chan *github.PullRequest {
	ret := make(chan *github.PullRequest)
	go func() {
		for pr := range input {
			if pr.GetState() != "open" {
				continue
			}

			issue, _, err := issueClient.Get(
				context.Background(),
				owner,
				repository,
				pr.GetNumber(),
			)
			if err != nil {
				continue
			}

			mergeable := false
			for _, label := range issue.Labels {
				mergeable = mergeable || *label.Name == mergeLabel
			}

			if !mergeable {
				continue
			}

			status, _, err := statusClient.GetCombinedStatus(
				context.Background(),
				owner,
				repository,
				*pr.Head.SHA,
				&github.ListOptions{},
			)
			if err != nil {
				continue
			}

			log.Printf("status for %q (%q): %q\n", pr.Head.GetRef(), pr.Head.GetSHA(), status.GetState())
			if status.GetState() != "success" {
				continue
			}

			if !*pr.Mergeable {
				continue
			}

			ret <- pr
		}
		close(ret)
	}()
	return ret
}
