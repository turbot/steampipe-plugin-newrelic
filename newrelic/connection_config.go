package newrelic

import (
	"context"
	"errors"
	"github.com/newrelic/newrelic-client-go/v2/newrelic"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
	"os"
)

type PluginConfig struct {
	APIKey *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
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

	nrConfig := GetConfig(d.Connection)
	if nrConfig.APIKey != nil {
		apiKey = *nrConfig.APIKey
	}

	if apiKey == "" {
		return nil, errors.New("the 'api_key' must be set in the connection configuration file or 'NEW_RELIC_API_KEY' env var must be set. Please set and then restart Steampipe")
	}

	client, err := newrelic.New(newrelic.ConfigPersonalAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return client, nil
}
