package console

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const clusterPlanRoute = "/api/plans"

type clusterCreateResponse struct {
	ClusterId string `json:"clusterId"`
}

type ClusterCreateRequest struct {
	Name         string `json:"name"`
	PlanTypeId   string `json:"planTypeId"`
	ChannelId    string `json:"channelId"`
	GenerationId string `json:"generationId"`
	RegionId     string `json:"k8sContextId"`
	AutoUpdate   bool   `json:"autoUpdate"`
	StageLabel   string `json:"stageLabel"`
}

type NamedClusterCreateRequest struct {
	Name       string
	PlanType   string
	Channel    string
	Generation string
	Region     string
	AutoUpdate bool
	StageLabel string
	//Trial      Cluster
}

const clusterRoute = "/clusters"

func (c *Console) CreateCluster(ctx context.Context, cluster NamedClusterCreateRequest) (string, error) {
	parameters, err := c.getClusterParameters(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to get parameters for cluster creation %w", err)
	}

	regionID, err := parameters.Regions.GetIDFromName(cluster.Region)
	if err != nil {
		return "", err
	}

	planTypeID, err := parameters.ClusterPlanTypes.GetIDFromName(cluster.PlanType)
	if err != nil {
		return "", err
	}

	channelID, generationID, err := parameters.Channels.GetIDsFromName(cluster.Channel, cluster.Generation)
	if err != nil {
		return "", err
	}

	requestedCluster := ClusterCreateRequest{
		Name:         cluster.Name,
		PlanTypeId:   planTypeID,
		ChannelId:    channelID,
		GenerationId: generationID,
		RegionId:     regionID,
		AutoUpdate:   false,
		StageLabel:   "dev",
	}
	client := c.auth0.Oauth().Client(ctx, c.accessToken)
	endpoint := "https://" + c.auth0.EndpointURL() + "/api/orgs/" + c.ActiveOrg.Uuid + clusterRoute

	payload, err := json.Marshal(requestedCluster)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusOK|http.StatusAccepted|http.StatusCreated {
		err = fmt.Errorf("error %w", err)
		return "", err
	}

	var clusterCreated clusterCreateResponse
	err = json.NewDecoder(resp.Body).Decode(&clusterCreated)
	if err != nil {
		return "", err
	}
	return clusterCreated.ClusterId, err
}

func (c *Console) DeleteCluster(ctx context.Context, id string) error {
	client := c.auth0.Oauth().Client(ctx, c.accessToken)
	endpoint := "https://" + c.auth0.EndpointURL() + "/api/orgs/" + c.ActiveOrg.Uuid + clusterRoute + "/" + id

	request, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusOK|http.StatusAccepted|http.StatusCreated {
		err = fmt.Errorf("error %s, %w", string(all), err)
		return err
	}

	return nil
}
