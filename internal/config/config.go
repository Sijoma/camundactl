package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

func CreateConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			log.Println(err)
		}

		// Search config in home directory with name ".camundactl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".camundactl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		// Create config file if it does not exist
		home, err := os.UserHomeDir()
		if err != nil {
			log.Println(err)
		}

		configName := "/.camundactl.yaml"
		fmt.Println("config ", home, configName)
		_, err = os.Create(home + configName)
		if err != nil {
			log.Println(err)
		}
	}

}

const accessTokenKey = ".auth"

func StoreAccessToken(stage string, token *oauth2.Token) {
	viper.Set(stage+accessTokenKey, token)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("could not write config file " + err.Error())
		return
	}
}

func GetAccessToken(stage string) (*oauth2.Token, error) {
	configToken := viper.Get(stage + accessTokenKey)
	tok, _ := yaml.Marshal(configToken)
	accessToken := &oauth2.Token{}
	err := yaml.Unmarshal(tok, accessToken)
	if err != nil {
		fmt.Println("error unmarshalling token")
		return nil, err
	}

	return accessToken, nil
}

const activeOrgIDKey = ".org"

func StoreOrgID(stage string, orgID string) {
	viper.Set(stage+activeOrgIDKey, orgID)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("could not write config file ", err)
		return
	}
}

func GetActiveOrgID(stage string) string {
	orgID := viper.GetString(stage + activeOrgIDKey)
	return orgID
}
