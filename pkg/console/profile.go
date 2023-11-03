package console

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Console) fetchOrgs(ctx context.Context) ([]Organization, error) {
	client := c.auth0.Oauth().Client(ctx, c.accessToken)
	endpoint := c.auth0.AccountsURL() + "/api/organizations/my"
	//fmt.Println("POST endpoint: " + endpoint)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusOK|http.StatusAccepted|http.StatusCreated {
		err = fmt.Errorf("request for organizations failed")
		return nil, err
	}

	var orgs []Organization
	err = json.NewDecoder(resp.Body).Decode(&orgs)
	if err != nil {
		return nil, fmt.Errorf("failed to decode organizations: %w", err)
	}

	if len(orgs) < 1 {
		return nil, fmt.Errorf("no organizations found")
	}

	return orgs, err
}
