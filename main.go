package main

import (
	"github.com/turbot/steampipe-plugin-newrelic/newrelic"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: newrelic.Plugin})
}
