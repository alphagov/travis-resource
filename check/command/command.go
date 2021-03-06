package command

import (
	"context"
	"fmt"

	"github.com/alphagov/travis-resource/messager"
	"github.com/alphagov/travis-resource/model"
	"github.com/shuheiktgw/go-travis"
)

type CheckCommand struct {
	TravisClient *travis.Client
	Request      model.CheckRequest
	Messager     *messager.ResourceMessager
}

func (c *CheckCommand) SendResponse(buildId uint) {
	var response model.CheckResponse
	if buildId != 0 {
		response = model.CheckResponse{model.Version{fmt.Sprint(buildId)}}
	} else {
		response = model.CheckResponse{}
	}

	c.Messager.SendJsonResponse(response)
}
func (c *CheckCommand) GetBuildId() (uint, error) {
	state := travis.BuildStatePassed
	if c.Request.Source.CheckOnState != "" {
		state = c.Request.Source.CheckOnState
	}
	if c.Request.Source.CheckAllBuilds {
		state = ""
	}

	options := travis.BuildsByRepoOption{
		BranchName: []string{c.Request.Source.Branch},
		State:      []string{state},
	}

	builds, _, err := c.TravisClient.Builds.ListByRepoSlug(
		context.Background(),
		c.Request.Source.Repository,
		&options,
	)

	if err != nil {
		return 0, err
	}
	if len(builds) == 0 {
		return 0, nil
	}

	return *builds[0].Id, nil
}
