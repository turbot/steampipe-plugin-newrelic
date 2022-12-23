package newrelic

import (
	"context"
	"fmt"
	"github.com/newrelic/newrelic-client-go/v2/pkg/apm"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableApmApplicationMetric() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_apm_application_metric",
		Description: "Obtain a list of available metrics for a specific application",
		List: &plugin.ListConfig{
			Hydrate: listApmApplicationMetrics,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "app_id",
					Require:   plugin.Required,
					Operators: []string{"="},
				},
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: apmApplicationMetricColumns(),
	}
}

func listApmApplicationMetrics(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	q := d.KeyColumnQuals
	appId := int(q["app_id"].GetInt64Value())

	params := apm.MetricNamesParams{}
	if q["name"] != nil {
		params.Name = q["name"].GetStringValue()
	}

	metrics, err := client.APM.GetMetricNamesWithContext(ctx, appId, params)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain APM application %d metrics: %v", appId, err)
	}

	for _, metric := range metrics {
		d.StreamListItem(ctx, metric)
	}

	return nil, nil
}

func apmApplicationMetricColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "name",
			Description: "Name of the application metric",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "values",
			Description: "An array of the values you can obtain for this metric",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "app_id",
			Description: "Identifier for the application",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromQual("app_id"),
		},
	}
}
