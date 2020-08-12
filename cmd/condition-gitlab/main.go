package main

import (
	gitlabCondition "github.com/go-semantic-release/condition-gitlab/pkg/condition"
	"github.com/go-semantic-release/semantic-release/v2/pkg/condition"
	"github.com/go-semantic-release/semantic-release/v2/pkg/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		CICondition: func() condition.CICondition {
			return &gitlabCondition.GitLab{}
		},
	})
}
