package newrelic

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAlertIncident() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_alert_incident",
		Description: "Obtain alert incidents from the given New Relic account",
		List: &plugin.ListConfig{
			Hydrate: listAlertIncidents,
		},
		Columns: commonColumns(alertIncidentColumns()),
	}
}

func listAlertIncidents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_incident.listAlertIncidents", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	plugin.Logger(ctx).Debug("newrelic_alert_incident.listAlertIncidents", "onlyOpen", false, "excludeViolations", false)
	ais, err := client.Alerts.ListIncidentsWithContext(ctx, false, false)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_incident.listAlertIncidents", "query_error", err)
		return nil, fmt.Errorf("unable to obtain alert incidents: %v", err)
	}

	for _, ai := range ais {
		d.StreamListItem(ctx, ai)
	}

	return nil, nil
}

func alertIncidentColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier for the alert incident.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "opened_at",
			Description: "Timestamp of when the incident was created.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("OpenedAt").Transform(epochTransform),
		},
		{
			Name:        "closed_at",
			Description: "Timestamp of when the incident was closed.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("CreatedAt").Transform(epochTransform),
		},
		{
			Name:        "incident_preference",
			Description: "The preference of the incident.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "policy_id",
			Description: "Identifier of the policy the incident is associated with.",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Links.PolicyID"),
		},
		{
			Name:        "violations",
			Description: "An array of violation identifiers associated with the incident.",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Links.Violations"),
		},
	}
}
