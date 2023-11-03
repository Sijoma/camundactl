package console

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (c *Console) getClusterParameters(ctx context.Context) (*ClusterParameters, error) {
	client := c.auth0.Oauth().Client(ctx, c.accessToken)
	endpoint := "https://" + c.auth0.EndpointURL() + "/api/orgs/" + c.ActiveOrg.Uuid + "/clusters/parameters"

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
		err = fmt.Errorf("request for parameters failed")
		return nil, err
	}

	var parameters ClusterParameters
	err = json.NewDecoder(resp.Body).Decode(&parameters)
	if err != nil {
		return nil, fmt.Errorf("failed to decode parameters: %w", err)
	}

	return &parameters, err
}

type ClusterPlanType struct {
	Name        string `json:"name"`
	CreatedBy   string `json:"createdBy"`
	Internal    bool   `json:"internal"`
	Description string `json:"description"`
	Uuid        string `json:"uuid"`
	Id          int    `json:"id"`
}

type Region struct {
	Uuid   string `json:"k8sContextUuid"`
	Config struct {
		Name   string `json:"name"`
		Region string `json:"region"`
	} `json:"config"`
}

type Regions []Region

func (r Regions) GetIDFromName(name string) (string, error) {
	for _, region := range r {
		if region.Config.Region == name {
			return region.Uuid, nil
		}
	}
	return "", fmt.Errorf("region with name %s not found", name)
}

type ClusterPlanTypes []ClusterPlanType

func (c ClusterPlanTypes) GetIDFromName(name string) (string, error) {
	for _, clusterPlanType := range c {
		if clusterPlanType.Name == name {
			return clusterPlanType.Uuid, nil
		}
	}
	return "", fmt.Errorf("cluster plan type with name %s not found", name)
}

type Channel struct {
	Uuid              string    `json:"uuid"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	CreatedBy         string    `json:"createdBy"`
	Created           time.Time `json:"created"`
	IsDefault         bool      `json:"isDefault"`
	Id                int       `json:"id"`
	UpdatedBy         string    `json:"updatedBy"`
	Updated           time.Time `json:"updated"`
	DefaultGeneration struct {
		Uuid string `json:"uuid"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"defaultGeneration"`
	AllowedGenerations []struct {
		Uuid string `json:"uuid"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"allowedGenerations"`
}

type Channels []Channel

func (c Channels) GetIDsFromName(channelName, generationName string) (channelID string, generationID string, err error) {
	for _, channel := range c {
		if channel.Name == channelName {
			for _, generation := range channel.AllowedGenerations {
				if generation.Name == generationName {
					return channel.Uuid, generation.Uuid, nil
				}
			}
		}
	}
	return "", "", fmt.Errorf("combination of channel %s and generation  %s not found", channelName, generationName)
}

type ClusterParameters struct {
	Channels                      Channels         `json:"channels"`
	ClusterPlanTypes              ClusterPlanTypes `json:"clusterPlanTypes"`
	ClusterPlanTypesFromSalesPlan ClusterPlanTypes `json:"clusterPlanTypesFromSalesPlan"`
	Regions                       Regions          `json:"regions"`
}
