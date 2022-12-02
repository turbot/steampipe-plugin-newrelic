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
			"newrelic_account":         tableAccount(),
			"newrelic_alert_channel":   tableAlertChannel(),
			"newrelic_alert_condition": tableAlertCondition(),
			"newrelic_alert_event":     tableAlertEvent(),
			"newrelic_alert_incident":  tableAlertIncident(),
			"newrelic_alert_policy":    tableAlertPolicy(),
		},
	}

	return p
}
