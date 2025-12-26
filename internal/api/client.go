package api

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func Client(PAT string) *githubv4.Client {
	token := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: PAT})
	httpClient := oauth2.NewClient(context.Background(), token)
	return githubv4.NewClient(httpClient)
}
