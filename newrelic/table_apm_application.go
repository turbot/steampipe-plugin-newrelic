package newrelic

import (
	"context"
	"fmt"
	"github.com/newrelic/newrelic-client-go/v2/pkg/apm"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableApmApplication() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_apm_application",
		Description: "Obtain applications for which APM are collected",
		List: &plugin.ListConfig{
			Hydrate: listApmApplications,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "language",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getApmApplication,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: apmApplicationColumns(),
	}
}

func getApmApplication(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_apm_application.getApmApplication", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	appId := int(d.EqualsQuals["id"].GetInt64Value())

	plugin.Logger(ctx).Debug("newrelic_apm_application.getApmApplication", "app.Id", appId)
	app, err := client.APM.GetApplicationWithContext(ctx, appId)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_apm_application.getApmApplication", "query_error", err)
		return nil, fmt.Errorf("unable to obtain APM application with id %d: %v", appId, err)
	}

	return app, nil
}

func listApmApplications(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_apm_application.listApmApplications", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	params := &apm.ListApplicationsParams{}
	q := d.EqualsQuals
	if q["name"] != nil {
		params.Name = q["name"].GetStringValue()
	}
	if q["language"] != nil {
		params.Language = q["language"].GetStringValue()
	}

	plugin.Logger(ctx).Debug("newrelic_apm_application.listApmApplications", "params.Name", params.Name, "params.Language", params.Language)
	apps, err := client.APM.ListApplicationsWithContext(ctx, params)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_apm_application.listApmApplications", "query_error", err)
		return nil, fmt.Errorf("unable to obtain APM applications: %v", err)
	}

	for _, app := range apps {
		d.StreamListItem(ctx, app)
	}

	return nil, nil
}

func apmApplicationColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier of the application.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "name",
			Description: "Friendly/display name of the application.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "language",
			Description: "Language of the application.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "health_status",
			Description: "Current reported health status of the application.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "reporting",
			Description: "Indicates if reporting is enabled for the application.",
			Type:        proto.ColumnType_BOOL,
		},
		{
			Name:        "last_reported_at",
			Description: "Last report received.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("LastReportedAt").NullIfZero(),
		},
		// Summary
		{
			Name:        "response_time",
			Description: "Current response time of the application (at last metric).",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Summary.ResponseTime"),
		},
		{
			Name:        "throughput",
			Description: "Current throughput of the application (at last metric).",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Summary.Throughput"),
		},
		{
			Name:        "error_rate",
			Description: "Current error rate of the application (at last metric).",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Summary.ErrorRate"),
		},
		{
			Name:        "apdex_target",
			Description: "The apdex target of the application.",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Summary.ApdexTarget"),
		},
		{
			Name:        "apdex_score",
			Description: "Current apdex score of the application (at last metric).",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Summary.ApdexScore"),
		},
		{
			Name:        "host_count",
			Description: "Count of hosts the application is installed on.",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Summary.HostCount"),
		},
		{
			Name:        "instance_count",
			Description: "Count of instances of the application installed.",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Summary.InstanceCount"),
		},
		{
			Name:        "concurrent_instance_count",
			Description: "Count of active concurrent instances (at last metric).",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Summary.ConcurrentInstanceCount"),
		},

		// End-User Summary
		{
			Name:        "end_user_response_time",
			Description: "Current response time from end-user perspective of the application (at last metric).",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("EndUserSummary.ResponseTime"),
		},
		{
			Name:        "end_user_throughput",
			Description: "Current throughput from end-user perspective of the application (at last metric).",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("EndUserSummary.Throughput"),
		},
		{
			Name:        "end_user_apdex_target",
			Description: "The apdex target from end-user perspective of the application (at last metric).",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("EndUserSummary.ApdexTarget"),
		},
		{
			Name:        "end_user_apdex_score",
			Description: "Current apdex score from end-user perspective of the application (at last metric).",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("EndUserSummary.ApdexScore"),
		},
		// Settings
		{
			Name:        "apdex_threshold",
			Description: "The setting for apdex threshold against the application.",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Settings.AppApdexThreshold"),
		},
		{
			Name:        "end_user_apdex_threshold",
			Description: "The setting for apdex threshold from the end user perspective.",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Settings.EndUserApdexThreshold"),
		},
		{
			Name:        "real_user_monitoring",
			Description: "Indicates if real user monitoring is enabled for the application.",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Settings.EnableRealUserMonitoring"),
		},
		{
			Name:        "server_side_config",
			Description: "Indicates if server side config is used for the application.",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Settings.UseServerSideConfig"),
		},
		// Links
		{
			Name:        "servers",
			Description: "An array of identifiers for the servers associated with the application.",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Links.ServerIDs"),
		},
		{
			Name:        "hosts",
			Description: "An array of identifiers for the hosts associated with the applications.",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Links.HostIDs"),
		},
		{
			Name:        "instances",
			Description: "An array of identifiers for the instances of the application.",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Links.InstanceIDs"),
		},
		{
			Name:        "alert_policy_id",
			Description: "The identifier for the alert policy associated with the application.",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Links.AlertPolicyID"),
		},
	}
}
