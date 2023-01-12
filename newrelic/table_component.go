package newrelic

import (
	"context"
	"fmt"
	"github.com/newrelic/newrelic-client-go/v2/pkg/plugins"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableComponent() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_component",
		Description: "Obtain information about components in your NewRelic account",
		List: &plugin.ListConfig{
			Hydrate: listComponents,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "plugin_id",
					Require: plugin.Optional,
				},
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getComponent,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: componentColumns(),
	}
}

func getComponent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	componentId := int(d.EqualsQuals["id"].GetInt64Value())

	c, err := client.Plugins.GetComponent(componentId)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain component %d: %v", componentId, err)
	}

	return c, nil
}

func listComponents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	params := plugins.ListComponentsParams{HealthStatus: true}
	q := d.EqualsQuals

	if q["plugin_id"] != nil {
		params.PluginID = int(q["plugin_id"].GetInt64Value())
	}
	if q["name"] != nil {
		params.Name = q["name"].GetStringValue()
	}

	cs, err := client.Plugins.ListComponents(&params)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain components: %v", err)
	}

	for _, c := range cs {
		d.StreamListItem(ctx, c)
	}

	return nil, nil
}

func componentColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier for the component",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "name",
			Description: "The name of the component",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "health_status",
			Description: "The health status of the component",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "summary_metrics",
			Description: "An array of metric summaries associated with the component",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "plugin_id",
			Description: "Only populated if passed in as a query parameter",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromQual("plugin_id"),
		},
	}
}
