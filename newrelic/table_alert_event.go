package newrelic

import (
	"context"
	"fmt"
	"github.com/newrelic/newrelic-client-go/v2/pkg/alerts"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAlertEvent() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_alert_event",
		Description: "Obtain alert events from the given NewRelic account",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "product",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
				{
					Name:      "entity_type",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
				{
					Name:      "entity_group_id",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
				{
					Name:      "entity_id",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
				{
					Name:      "event_type",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
				{
					Name:      "incident_id",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
			},
			Hydrate: listAlertEvents,
		},
		Columns: alertEventColumns(),
	}
}

func listAlertEvents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_event.listAlertEvents", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	params := alerts.ListAlertEventsParams{}

	if d.EqualsQuals["product"] != nil {
		params.Product = d.EqualsQuals["product"].GetStringValue()
	}

	if d.EqualsQuals["entity_type"] != nil {
		params.EntityType = d.EqualsQuals["entity_type"].GetStringValue()
	}

	if d.EqualsQuals["entity_group_id"] != nil {
		params.EntityGroupID = int(d.EqualsQuals["entity_group_id"].GetInt64Value())
	}

	if d.EqualsQuals["entity_id"] != nil {
		params.EntityID = int(d.EqualsQuals["entity_id"].GetInt64Value())
	}

	if d.EqualsQuals["event_type"] != nil {
		params.EventType = d.EqualsQuals["event_type"].GetStringValue()
	}

	if d.EqualsQuals["incident_id"] != nil {
		params.IncidentID = int(d.EqualsQuals["incident_id"].GetInt64Value())
	}

	plugin.Logger(ctx).Debug("newrelic_alert_event.listAlertEvents", "params.Product", params.Product,
		"params.EntityType", params.EntityType, "params.EntityGroupID", params.EntityGroupID, "params.EntityID", params.EntityID,
		"params.EventType", params.EventType, "params.IncidentID", params.IncidentID)
	aes, err := client.Alerts.ListAlertEventsWithContext(ctx, &params)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_event.listAlertEvents", "query_error", err)
		return nil, fmt.Errorf("unable to obtain alert events: %v", err)
	}

	for _, ae := range aes {
		d.StreamListItem(ctx, ae)
	}

	return nil, nil
}

func alertEventColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier for the alert event.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "event_type",
			Description: "The type of the alert event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "product",
			Description: "The name of the product the alert event relates to.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "description",
			Description: "Description of the alert event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "entity_id",
			Description: "Identifier of the entity the alert event relates to.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "entity_type",
			Description: "The type of the entity the alert event relates to.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "entity_group_id",
			Description: "Identifier of the group the entity associated with this alert event.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "priority",
			Description: "Priority of the alert event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "incident_id",
			Description: "Identifier of the incident for which the alert event was raised.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "timestamp",
			Description: "Timestamp when alert event was raised.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Timestamp").Transform(epochTransform),
		},
	}
}
