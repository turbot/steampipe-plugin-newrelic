package newrelic

import (
	"context"
	"errors"
	"os"

	"github.com/newrelic/newrelic-client-go/v2/newrelic"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type PluginConfig struct {
	APIKey *string `hcl:"api_key"`
	Region *string `hcl:"region"`
}

func ConfigInstance() interface{} {
	return &PluginConfig{}
}

func GetConfig(connection *plugin.Connection) PluginConfig {
	if connection == nil || connection.Config == nil {
		return PluginConfig{}
	}

	config, _ := connection.Config.(PluginConfig)
	return config
}

func connect(_ context.Context, d *plugin.QueryData) (*newrelic.NewRelic, error) {
	apiKey := os.Getenv("NEW_RELIC_API_KEY")
	region := os.Getenv("NEW_RELIC_REGION")

	nrConfig := GetConfig(d.Connection)
	if nrConfig.APIKey != nil {
		apiKey = *nrConfig.APIKey
	}

	if nrConfig.Region != nil {
		region = *nrConfig.Region
	}

	if apiKey == "" {
		return nil, errors.New("'api_key' must be set in the connection configuration file or 'NEW_RELIC_API_KEY' env var must be set. Please set and then restart Steampipe")
	}

	if region == "" {
		region = "us" // Default to US
	}

	client, err := newrelic.New(newrelic.ConfigPersonalAPIKey(apiKey), newrelic.ConfigRegion(region))
	if err != nil {
		return nil, err
	}

	return client, nil
}
