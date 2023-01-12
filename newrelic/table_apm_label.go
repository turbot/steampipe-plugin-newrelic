package newrelic

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableApmLabel() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_apm_label",
		Description: "Obtain all labels in the account",
		List: &plugin.ListConfig{
			Hydrate: listApmLabels,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getApmLabel,
			KeyColumns: plugin.SingleColumn("key"),
		},
		Columns: apmLabelColumns(),
	}
}

func getApmLabel(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	key := d.EqualsQuals["key"].GetStringValue()

	label, err := client.APM.GetLabelWithContext(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain label with key '%s': %v", key, err)
	}

	return label, nil
}

func listApmLabels(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	labels, err := client.APM.ListLabelsWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain labels: %v", err)
	}

	for _, label := range labels {
		d.StreamListItem(ctx, label)
	}

	return nil, nil
}

func apmLabelColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "key",
			Description: "The identifier for the label",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "name",
			Description: "The name of the label",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "category",
			Description: "The category the label belongs to",
			Type:        proto.ColumnType_STRING,
		},
		// Links
		{
			Name:        "applications",
			Description: "An array of application identifiers to which the label is applied",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Links.Applications"),
		},
		{
			Name:        "servers",
			Description: "An array of server identifiers to which the label is applied",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Links.Servers"),
		},
	}
}
