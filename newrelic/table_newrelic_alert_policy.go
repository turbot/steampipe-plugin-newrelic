package newrelic

import (
	"context"
	"fmt"

	"github.com/newrelic/newrelic-client-go/v2/pkg/alerts"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAlertPolicy() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_alert_policy",
		Description: "Obtain alert policies from the given New Relic account",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "name",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
			},
			Hydrate: listAlertPolicies,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAlertPolicy,
		},
		Columns: commonColumns(alertPolicyColumns()),
	}
}

func getAlertPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_policy.getAlertPolicy", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	policyId := int(d.EqualsQuals["id"].GetInt64Value())

	plugin.Logger(ctx).Debug("newrelic_alert_policy.getAlertPolicy", "policy.Id", policyId)
	p, err := client.Alerts.GetPolicy(policyId)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_policy.getAlertPolicy", "query_error", err)
		return nil, fmt.Errorf("unable to obtain policy %d: %v", policyId, err)
	}

	return p, nil
}

func listAlertPolicies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_policy.listAlertPolicies", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	params := alerts.ListPoliciesParams{}

	if d.EqualsQuals["name"] != nil {
		params.Name = d.EqualsQuals["name"].GetStringValue()
	}

	plugin.Logger(ctx).Debug("newrelic_alert_policy.listAlertPolicies", "params.Name", params.Name)
	ps, err := client.Alerts.ListPoliciesWithContext(ctx, &params)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_policy.listAlertPolicies", "query_error", err)
		return nil, fmt.Errorf("unable to obtain alert policies: %v", err)
	}

	for _, p := range ps {
		d.StreamListItem(ctx, p)
	}

	return nil, nil
}

func alertPolicyColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier for the alert policy.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "name",
			Description: "Name of the policy.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "incident_preference",
			Description: "The preference type of the incident (PER_POLICY, PER_CONDITION, PER_CONDITION_AND_TARGET).",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "created_at",
			Description: "Timestamp at which the policy was created.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("CreatedAt").Transform(epochTransform),
		},
		{
			Name:        "updated_at",
			Description: "Timestamp at which the policy was updated.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("UpdatedAt").Transform(epochTransform),
		},
	}
}
