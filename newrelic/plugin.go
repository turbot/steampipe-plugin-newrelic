package newrelic

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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
			"newrelic_account":                     tableAccount(),
			"newrelic_alert_channel":               tableAlertChannel(),
			"newrelic_alert_condition":             tableAlertCondition(),
			"newrelic_alert_event":                 tableAlertEvent(),
			"newrelic_alert_incident":              tableAlertIncident(),
			"newrelic_alert_policy":                tableAlertPolicy(),
			"newrelic_apm_application":             tableApmApplication(),
			"newrelic_apm_application_deployment":  tableApmApplicationDeployment(),
			"newrelic_apm_application_instance":    tableApmApplicationInstance(),
			"newrelic_apm_application_metric":      tableApmApplicationMetric(),
			"newrelic_apm_application_metric_data": tableApmApplicationMetricData(),
			"newrelic_plugin":                      tablePlugin(),
			"newrelic_component":                   tableComponent(),
			"newrelic_notification_channel":        tableNotificationChannel(),
			"newrelic_notification_destination":    tableNotificationDestination(),
		},
	}

	return p
}
