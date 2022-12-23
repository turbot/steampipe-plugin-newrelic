package newrelic

import (
	"context"
	"fmt"
	"github.com/newrelic/newrelic-client-go/v2/pkg/apm"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableApmApplicationInstance() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_apm_application_instance",
		Description: "Obtain information about instances of applications for which APM are collected",
		List: &plugin.ListConfig{
			Hydrate: listApmApplicationInstances,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "app_id",
					Require:   plugin.Required,
					Operators: []string{"="},
				},
				{
					Name:    "id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: apmApplicationInstanceColumns(),
	}
}

func listApmApplicationInstances(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	q := d.KeyColumnQuals
	appId := int(q["app_id"].GetInt64Value())

	if q["id"] != nil {
		id := int(q["id"].GetInt64Value())
		i, err := client.APM.GetApplicationInstanceWithContext(ctx, appId, id)
		if err != nil {
			return nil, fmt.Errorf("unable to obtain APM application %d instance %d: %v", appId, id, err)
		}

		d.StreamListItem(ctx, i)
		return nil, nil
	}

	params := &apm.ListApplicationInstancesParams{}

	instances, err := client.APM.ListApplicationInstancesWithContext(ctx, appId, params)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain APM application %d instances: %v", appId, err)
	}

	for _, instance := range instances {
		d.StreamListItem(ctx, instance)
	}

	return nil, nil
}

func apmApplicationInstanceColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier of the application instance",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "application_name",
			Description: "The name of the application for which this is an instance of",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "host",
			Description: "Name/Identifier for the host on which the application instance is running",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "port",
			Description: "The port on which the application instance is configured",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "language",
			Description: "Language of the application instance",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "health_status",
			Description: "Current reported health status of the application instance",
			Type:        proto.ColumnType_STRING,
		},
		// Summary
		{
			Name:        "response_time",
			Description: "Current response time of the application instance (at last metric)",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Summary.ResponseTime"),
		},
		{
			Name:        "throughput",
			Description: "Current throughput of the application instance (at last metric)",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Summary.Throughput"),
		},
		{
			Name:        "error_rate",
			Description: "Current error rate of the application instance (at last metric)",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Summary.ErrorRate"),
		},
		{
			Name:        "apdex_target",
			Description: "The apdex target of the application instance",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Summary.ApdexTarget"),
		},
		{
			Name:        "apdex_score",
			Description: "Current apdex score of the application instance (at last metric)",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("Summary.ApdexScore"),
		},
		{
			Name:        "host_count",
			Description: "Count of hosts the application instance is installed on",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Summary.HostCount"),
		},
		{
			Name:        "instance_count",
			Description: "Count of instances of the application instance installed",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Summary.InstanceCount"),
		},
		{
			Name:        "concurrent_instance_count",
			Description: "Count of active concurrent instances (at last metric)",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Summary.ConcurrentInstanceCount"),
		},
		// End-User Summary
		{
			Name:        "end_user_response_time",
			Description: "Current response time from end-user perspective of the application instance (at last metric)",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("EndUserSummary.ResponseTime"),
		},
		{
			Name:        "end_user_throughput",
			Description: "Current throughput from end-user perspective of the application instance (at last metric)",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("EndUserSummary.Throughput"),
		},
		{
			Name:        "end_user_apdex_target",
			Description: "The apdex target from end-user perspective of the application instance (at last metric)",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("EndUserSummary.ApdexTarget"),
		},
		{
			Name:        "end_user_apdex_score",
			Description: "Current apdex score from end-user perspective of the application instance (at last metric)",
			Type:        proto.ColumnType_DOUBLE,
			Transform:   transform.FromField("EndUserSummary.ApdexScore"),
		},
		// Links
		{
			Name:        "app_id",
			Description: "Unique identifier for the application",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Links.Application"),
		},
		{
			Name:        "app_host",
			Description: "Unique identifier for the host on which this application instance is running",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Links.ApplicationHost"),
		},
	}
}
