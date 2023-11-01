package console

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (c *Console) getClusterParameters(ctx context.Context) (*ClusterParameters, error) {
	client := c.auth0.Oauth().Client(ctx, c.AccessToken)
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
	Name        string      `json:"name"`
	CreatedBy   string      `json:"createdBy"`
	Internal    bool        `json:"internal"`
	Description string      `json:"description"`
	Uuid        string      `json:"uuid"`
	Development bool        `json:"development"`
	Category    string      `json:"category"`
	Id          int         `json:"id"`
	Created     time.Time   `json:"created"`
	Deprecated  bool        `json:"deprecated"`
	Limits      interface{} `json:"limits"`
	ActivePlan  struct {
		Name      string `json:"name"`
		CreatedBy string `json:"createdBy"`
		Plan      struct {
			Zeebe struct {
				Broker struct {
					ClusterSize       int `json:"clusterSize"`
					ReplicationFactor int `json:"replicationFactor"`
					PartitionsCount   int `json:"partitionsCount"`
					Storage           struct {
						StorageClassName string `json:"storageClassName"`
						Resources        struct {
							Requests struct {
								Storage string `json:"storage"`
							} `json:"requests"`
							Limits struct {
								Storage string `json:"storage"`
							} `json:"limits"`
						} `json:"resources"`
						AutoResizing struct {
							Threshold string `json:"threshold"`
							Increase  string `json:"increase"`
						} `json:"autoResizing"`
					} `json:"storage"`
					Resources struct {
						Limits struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"limits"`
						Requests struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"requests"`
					} `json:"resources"`
					EnvVars string `json:"envVars"`
				} `json:"broker"`
				Gateway struct {
					Replicas   int  `json:"replicas"`
					Standalone bool `json:"standalone"`
					Backend    struct {
						Resources struct {
							Limits struct {
							} `json:"limits"`
							Requests struct {
							} `json:"requests"`
						} `json:"resources"`
					} `json:"backend"`
				} `json:"gateway"`
			} `json:"zeebe"`
			Operate struct {
				Backend struct {
					Resources struct {
						Limits struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"limits"`
						Requests struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"requests"`
					} `json:"resources"`
					EnvVars string `json:"envVars"`
				} `json:"backend"`
				Replicas      int `json:"replicas"`
				Elasticsearch struct {
					Resources struct {
						Limits struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"limits"`
						Requests struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"requests"`
					} `json:"resources"`
					Config struct {
						NodesCount int `json:"nodesCount"`
						Storage    struct {
							StorageClassName string `json:"storageClassName"`
							Resources        struct {
								Requests struct {
									Storage string `json:"storage"`
								} `json:"requests"`
								Limits struct {
									Storage string `json:"storage"`
								} `json:"limits"`
							} `json:"resources"`
							AutoResizing struct {
								Threshold string `json:"threshold"`
								Increase  string `json:"increase"`
							} `json:"autoResizing"`
						} `json:"storage"`
					} `json:"config"`
				} `json:"elasticsearch"`
			} `json:"operate"`
			Tasklist struct {
				Backend struct {
					Resources struct {
						Limits struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"limits"`
						Requests struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"requests"`
					} `json:"resources"`
					EnvVars string `json:"envVars"`
				} `json:"backend"`
				Replicas int `json:"replicas"`
			} `json:"tasklist"`
			Optimize struct {
				Backend struct {
					Resources struct {
						Limits struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"limits"`
						Requests struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"requests"`
					} `json:"resources"`
					EnvVars string `json:"envVars"`
				} `json:"backend"`
				Replicas int `json:"replicas"`
			} `json:"optimize"`
			ZeebeAnalytics struct {
				Backend struct {
					Resources struct {
						Limits struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"limits"`
						Requests struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"requests"`
					} `json:"resources"`
				} `json:"backend"`
				Replicas int `json:"replicas"`
			} `json:"zeebeAnalytics"`
			ConnectorBridge struct {
				Backend struct {
					Resources struct {
						Limits struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"limits"`
						Requests struct {
							Memory string `json:"memory"`
							Cpu    string `json:"cpu"`
						} `json:"requests"`
					} `json:"resources"`
				} `json:"backend"`
				Replicas int `json:"replicas"`
			} `json:"connectorBridge"`
		} `json:"plan"`
		Uuid    string    `json:"uuid"`
		Id      int       `json:"id"`
		Created time.Time `json:"created"`
	} `json:"activePlan"`
}

type Region struct {
	K8SContextUuid string `json:"k8sContextUuid"`
	Config         struct {
		AllowTrial         bool   `json:"allow_trial"`
		Hidden             bool   `json:"hidden"`
		KubeconfigPath     string `json:"kubeconfig_path"`
		KubernetesContext  string `json:"kubernetes_context"`
		Name               string `json:"name"`
		Region             string `json:"region"`
		ThanosUrl          string `json:"thanos_url"`
		Zone               string `json:"zone"`
		ClusterCount       int    `json:"clusterCount"`
		KubeConfigFilePath string `json:"kubeConfigFilePath"`
		K8SIdentifier      string `json:"k8sIdentifier"`
		ThanosEndpoint     string `json:"thanosEndpoint"`
	} `json:"config"`
}

type Regions []Region

func (r Regions) GetIDFromName(name string) (string, error) {
	for _, region := range r {
		if region.Config.Region == name {
			return region.K8SContextUuid, nil
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
		Uuid     string `json:"uuid"`
		Name     string `json:"name"`
		Versions struct {
			Zeebe                  string `json:"zeebe"`
			Operate                string `json:"operate"`
			Tasklist               string `json:"tasklist"`
			Optimize               string `json:"optimize"`
			ZeebeAnalytics         string `json:"zeebeAnalytics"`
			ConnectorBridge        string `json:"connectorBridge"`
			ElasticSearchOss       string `json:"elasticSearchOss"`
			ElasticSearchCurator   string `json:"elasticSearchCurator"`
			OperateEnvVars         string `json:"operateEnvVars"`
			ZeebeBrokerEnvVars     string `json:"zeebeBrokerEnvVars"`
			ZeebeGatewayEnvVars    string `json:"zeebeGatewayEnvVars"`
			TasklistEnvVars        string `json:"tasklistEnvVars"`
			OptimizeEnvVars        string `json:"optimizeEnvVars"`
			ZeebeAnalyticsEnvVars  string `json:"zeebeAnalyticsEnvVars"`
			ConnectorBridgeEnvVars string `json:"connectorBridgeEnvVars"`
		} `json:"versions"`
		CreatedBy string    `json:"createdBy"`
		Created   time.Time `json:"created"`
		Id        int       `json:"id"`
		UpdatedBy string    `json:"updatedBy"`
		Updated   time.Time `json:"updated"`
	} `json:"defaultGeneration"`
	AllowedGenerations []struct {
		Uuid     string `json:"uuid"`
		Name     string `json:"name"`
		Versions struct {
			Zeebe                  string `json:"zeebe"`
			Operate                string `json:"operate"`
			Tasklist               string `json:"tasklist"`
			Optimize               string `json:"optimize"`
			ZeebeAnalytics         string `json:"zeebeAnalytics"`
			ConnectorBridge        string `json:"connectorBridge"`
			ElasticSearchOss       string `json:"elasticSearchOss"`
			ElasticSearchCurator   string `json:"elasticSearchCurator"`
			OperateEnvVars         string `json:"operateEnvVars"`
			ZeebeBrokerEnvVars     string `json:"zeebeBrokerEnvVars"`
			ZeebeGatewayEnvVars    string `json:"zeebeGatewayEnvVars"`
			TasklistEnvVars        string `json:"tasklistEnvVars"`
			OptimizeEnvVars        string `json:"optimizeEnvVars"`
			ZeebeAnalyticsEnvVars  string `json:"zeebeAnalyticsEnvVars"`
			ConnectorBridgeEnvVars string `json:"connectorBridgeEnvVars"`
		} `json:"versions"`
		CreatedBy string    `json:"createdBy"`
		Created   time.Time `json:"created"`
		Id        int       `json:"id"`
		UpdatedBy string    `json:"updatedBy"`
		Updated   time.Time `json:"updated"`
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
