package console

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type generationPayload struct {
	GenerationUUID      string `json:"generationUuid,omitempty"`
	ChannelUUID         string `json:"channelUuid,omitempty"`
	ClusterPlanUUID     string `json:"clusterPlanUuid,omitempty"`
	ClusterPlanTypeUUID string `json:"clusterPlanTypeUuid,omitempty"`
}

const externalOrgRoute = "/external/organizations/"

func (c *Console) PatchCluster(orgId, clusterId, generationUUID string) error {
	ctx := context.Background()
	client := c.auth0.Oauth().Client(ctx, c.AccessToken)
	endpoint := "https://" + c.auth0.EndpointURL() + externalOrgRoute + orgId + "/clusters/" + clusterId
	fmt.Println("PATCHING endpoint: " + endpoint)

	payload, _ := json.Marshal(generationPayload{GenerationUUID: generationUUID})

	request, err := http.NewRequest(http.MethodPatch, endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusOK|http.StatusAccepted|http.StatusCreated {
		err = errors.New(string(b))
	}
	return err
}
