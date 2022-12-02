package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"steampipe-plugin-newrelic/newrelic"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: newrelic.Plugin})
}
