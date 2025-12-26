package api

import (
	"context"
	"fmt"

	"github.com/gleich/lumber/v2"
	"github.com/shurcooL/githubv4"
)

type Repo struct {
	URL           string
	Name          string
	NameWithOwner string
	IsMirror      bool
	IsDisabled    bool
	IsArchived    bool
	IsEmpty       bool
	IsFork        bool
}

func Repos(client *githubv4.Client) ([]Repo, error) {
	fmt.Println()
	lumber.Info("Fetching repos")

	repos := []Repo{}
	vars := map[string]interface{}{
		"cursor": (*githubv4.String)(nil),
	}
	for {
		query := struct {
			Viewer struct {
				Repositories struct {
					Nodes    []Repo
					PageInfo struct {
						EndCursor   string
						HasNextPage bool
					}
				} `graphql:"repositories(affiliations: OWNER, first: 100, after: $cursor)"`
			}
		}{}

		err := client.Query(context.Background(), &query, vars)
		if err != nil {
			return []Repo{}, err
		}

		repos = append(repos, query.Viewer.Repositories.Nodes...)
		if !query.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
		vars["cursor"] = githubv4.String(query.Viewer.Repositories.PageInfo.EndCursor)
	}
	lumber.Success("Loaded", len(repos), "repos")
	return repos, nil
}
