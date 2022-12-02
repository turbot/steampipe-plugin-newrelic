package newrelic

import (
    "context"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
    p := &plugin.Plugin{
        Name: "steampipe-plugin-newrelic",
        ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
            NewInstance: ConfigInstance,
            Schema:      ConfigSchema,
        },
        DefaultTransform: transform.FromGo(),
        TableMap: map[string]*plugin.Table{
            // TODO: Map Tables
        },
    }

    return p
}
