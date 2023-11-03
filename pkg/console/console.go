package console

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/sijoma/camundactl/internal/auth0"
	"github.com/sijoma/camundactl/internal/config"
)

type Console struct {
	auth0         auth0.Auth0App
	accessToken   *oauth2.Token
	ActiveOrg     Organization
	Organizations []Organization
	stage         string
}

func NewConsole(ctx context.Context, stage string) *Console {
	configToken, err := config.GetAccessToken(stage)
	if err != nil {
		fmt.Println("There was an error getting the access token", configToken)
	}

	console := Console{
		auth0: auth0.NewAuth0App(stage),
		stage: stage,
	}

	if configToken.AccessToken != "" {
		console.accessToken = configToken
		err = console.UpdateProfile(ctx)
		if err != nil {
			fmt.Println("unable to update profile")
			return nil
		}
	}

	return &console
}

func NewMachineConsole(stage, clientID string) *Console {
	instance := Console{
		auth0: auth0.NewAuth0m2m(stage, clientID),
		stage: stage,
	}
	return &instance
}

func (c *Console) UpdateProfile(ctx context.Context) error {
	orgs, err := c.fetchOrgs(ctx)
	if err != nil {
		fmt.Println("unable to get organizations")
		return err
	}

	c.Organizations = orgs

	activeOrgID := config.GetActiveOrgID(c.stage)
	if activeOrgID == "" {
		// Default to first org
		fmt.Println("defaulting to first organization")
		activeOrgID = orgs[0].Uuid
	}
	err = c.SetOrg(activeOrgID)
	if err != nil {
		return fmt.Errorf("unable to set active org %w", err)
	}
	return nil
}

func (c *Console) PrintOrgs() {
	for _, org := range c.Organizations {
		fmt.Printf("Name: %s | ID: %s\n", org.Name, org.Uuid)
	}
}

func (c *Console) SetOrg(orgID string) error {
	for _, org := range c.Organizations {
		if org.Uuid == orgID {
			c.ActiveOrg = org
			config.StoreOrgID(c.stage, orgID)
			return nil
		}
	}
	return fmt.Errorf("organization with ID %s not found", orgID)
}
