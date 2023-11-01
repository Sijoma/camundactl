package console

import "time"

type Organization struct {
	Name        string    `json:"name"`
	Internal    bool      `json:"internal"`
	Uuid        string    `json:"uuid"`
	Created     time.Time `json:"created"`
	Permissions struct {
		Org struct {
			Clusters struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"clusters"`
			Clients struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"clients"`
			Settings struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"settings"`
			Usage struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"usage"`
			Billing struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"billing"`
			Users struct {
				General struct {
					Create bool `json:"create"`
					Read   bool `json:"read"`
					Update bool `json:"update"`
					Delete bool `json:"delete"`
				} `json:"general"`
				Owner struct {
					Create bool `json:"create"`
					Read   bool `json:"read"`
					Update bool `json:"update"`
					Delete bool `json:"delete"`
				} `json:"owner"`
				Admin struct {
					Create bool `json:"create"`
					Read   bool `json:"read"`
					Update bool `json:"update"`
					Delete bool `json:"delete"`
				} `json:"admin"`
				Member struct {
					Create bool `json:"create"`
					Read   bool `json:"read"`
					Update bool `json:"update"`
					Delete bool `json:"delete"`
				} `json:"member"`
			} `json:"users"`
			Activity struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"activity"`
			Diagrams struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"diagrams"`
			Forms struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"forms"`
			Instances struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"instances"`
			Workflows struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"workflows"`
			Webide struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"webide"`
			EarlyAccessFeatures struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"earlyAccessFeatures"`
		} `json:"org"`
		Cluster struct {
			Clients struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"clients"`
			Alerts struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"alerts"`
			Operate struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"operate"`
			Tasklist struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"tasklist"`
			Optimize struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"optimize"`
			IpWhitelist struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"ipWhitelist"`
			ConnectorSecrets struct {
				Create bool `json:"create"`
				Read   bool `json:"read"`
				Update bool `json:"update"`
				Delete bool `json:"delete"`
			} `json:"connectorSecrets"`
			Settings struct {
				ResourceBasedAuthorizations struct {
					Activate struct {
						Create bool `json:"create"`
						Read   bool `json:"read"`
						Update bool `json:"update"`
						Delete bool `json:"delete"`
					} `json:"activate"`
					Configure struct {
						Create bool `json:"create"`
						Read   bool `json:"read"`
						Update bool `json:"update"`
						Delete bool `json:"delete"`
					} `json:"configure"`
				} `json:"resourceBasedAuthorizations"`
			} `json:"settings"`
		} `json:"cluster"`
	} `json:"permissions"`
}
