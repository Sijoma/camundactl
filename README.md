# camundactl ⌨

## Config file

The default location of the config file is `/Users/<your_name>/.camundactl.yaml`

You can copy the example from this repository and update your access token. 
The Oauth2 flow only works on the development & integration stage at this moment. 

## Usage

camunda ctl allows to provison Camunda SaaS resources

Usage:
    camundactl [command]

Available Commands:
```text
    cluster     CRUD cluster commands
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    login       Authenticate to Camunda Console and store the accessToken in the configuration file.
    org         Set your current Camunda org
    version     Prints version info
```

Flags:
```text 
    --accessToken string     console access token
    --client_id string       the id of the client
    --client_secret string   the secret of the client
    --config string          config file (default is $HOME/.camundactl.yaml)
    -h, --help                   help for camundactl
    --stage string           the console stage to be used, either 'dev'. 'int' or 'prod' (default "prod")
    -t, --toggle                 Help message for toggle
```

Use "camundactl [command] --help" for more information about a command.


## Cluster create

Example: camundactl cluster create my-fancy-cluster

Usage:
  camundactl cluster create [cluster name] [flags]

Flags:
```text 
      --auto true              whether auto updates are active, defaults to true (default true)
      --channel Alpha          the channel of the cluster for example Alpha (default "Alpha")
      --gen Trial Cluster      the cluster type for example Trial Cluster (default "Camunda 8.3.1")
  -h, --help                   help for create
      --region Trial Cluster   the cluster type for example Trial Cluster (default "europe-west1")
      --stage dev              the stage label, one of dev, `test`, `stage`, `prod` (default "dev")
      --type Trial Cluster     the cluster type for example Trial Cluster (default "Trial Cluster")
```
