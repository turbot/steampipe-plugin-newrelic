package newrelic

import (
	"context"
	"fmt"
	"github.com/newrelic/newrelic-client-go/v2/pkg/apm"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"time"
)

func tableApmApplicationMetricData() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_apm_application_metric_data",
		Description: "Obtain metric data for application metrics",
		List: &plugin.ListConfig{
			Hydrate: listApmApplicationMetricData,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "app_id",
					Require:   plugin.Required,
					Operators: []string{"="},
				},
				{
					Name:    "from",
					Require: plugin.Optional,
				},
				{
					Name:    "to",
					Require: plugin.Optional,
				},
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: apmApplicationMetricDataColumns(),
	}
}

func listApmApplicationMetricData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_apm_application_metric_data.listApmApplicationMetricData", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	q := d.EqualsQuals
	appId := int(q["app_id"].GetInt64Value())

	params := apm.MetricDataParams{
		Names: []string{"*"},
	}
	if q["from"] != nil {
		f := q["from"].GetTimestampValue().AsTime()
		params.From = &f
	}
	if q["to"] != nil {
		t := q["to"].GetTimestampValue().AsTime()
		params.To = &t
	}
	if q["name"] != nil {
		params.Names = []string{q["name"].GetStringValue()}
	}

	plugin.Logger(ctx).Debug("newrelic_apm_application_metric_data.listApmApplicationMetricData", "app.Id", appId,
		"params.Names", params.Names, "params.From", params.From, "params.To", params.To)
	metrics, err := client.APM.GetMetricDataWithContext(ctx, appId, params)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_apm_application_metric_data.listApmApplicationMetricData", "query_error", err)
		return nil, fmt.Errorf("unable to obtain metric data for application %d: %v", appId, err)
	}

	for _, metric := range metrics {
		for _, ts := range metric.Timeslices {
			d.StreamListItem(ctx, MetricData{
				Name:   metric.Name,
				From:   ts.From,
				To:     ts.To,
				Values: ts.Values.Values,
			})
		}
	}

	return nil, nil
}

func apmApplicationMetricDataColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "app_id",
			Description: "Unique identifier of the application the metrics belong to.",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromQual("app_id"),
		},
		{
			Name:        "name",
			Description: "Name of the metric.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "from",
			Description: "Timestamp of beginning of the metrics.",
			Type:        proto.ColumnType_TIMESTAMP,
		},
		{
			Name:        "to",
			Description: "Timestamp of ending of the metrics.",
			Type:        proto.ColumnType_TIMESTAMP,
		},
		{
			Name:        "values",
			Description: "The actual metric data.",
			Type:        proto.ColumnType_JSON,
		},
	}
}

type MetricData struct {
	Name   string             `json:"name"`
	From   *time.Time         `json:"from"`
	To     *time.Time         `json:"to"`
	Values map[string]float64 `json:"values"`
}
