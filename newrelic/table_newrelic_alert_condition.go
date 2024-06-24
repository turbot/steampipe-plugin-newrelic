package newrelic

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAlertCondition() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_alert_condition",
		Description: "Obtain alert conditions for a given New Relic account",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("policy_id"),
			Hydrate:    listAlertConditions,
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "id",
					Require:   plugin.Required,
					Operators: []string{"="},
				},
				{
					Name:      "policy_id",
					Require:   plugin.Required,
					Operators: []string{"="},
				},
			},
			Hydrate: getAlertCondition,
		},
		Columns: commonColumns(alertConditionColumns()),
	}
}

func getAlertCondition(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_condition.getAlertCondition", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	conditionId := int(d.EqualsQuals["id"].GetInt64Value())
	policyId := int(d.EqualsQuals["policy_id"].GetInt64Value())

	plugin.Logger(ctx).Debug("newrelic_alert_condition.getAlertCondition", "condition.Id", conditionId, "policy.Id", policyId)
	c, err := client.Alerts.GetConditionWithContext(ctx, policyId, conditionId)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_condition.getAlertCondition", "query_error", err)
		return nil, fmt.Errorf("unable to obtain alert condition %d for policy %d: %v", conditionId, policyId, err)
	}

	return c, nil
}

func listAlertConditions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_condition.listAlertConditions", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	policyId := int(d.EqualsQuals["policy_id"].GetInt64Value())

	plugin.Logger(ctx).Debug("newrelic_alert_condition.listAlertConditions", "policy.Id", policyId)
	acs, err := client.Alerts.ListConditionsWithContext(ctx, policyId)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_alert_condition.listAlertConditions", "query_error", err)
		return nil, fmt.Errorf("unable to obtain alert conditions for policy %d: %v", policyId, err)
	}

	for _, ac := range acs {
		d.StreamListItem(ctx, ac)
	}

	return nil, nil
}

func alertConditionColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier for the alert condition.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "name",
			Description: "The name of the alert condition.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "type",
			Description: "The type of the alert condition.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "enabled",
			Description: "Indicates if the alert condition is enabled.",
			Type:        proto.ColumnType_BOOL,
		},
		{
			Name:        "entities",
			Description: "An array of entities associated with the alert condition.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "metric",
			Description: "The metric type associated with the alert condition.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "runbook_url",
			Description: "The url of the runbook associated with the alert condition.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("RunbookURL"),
		},
		{
			Name:        "terms",
			Description: "An array of term objects associated with the alert condition.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "user_metric",
			Description: "User defined metric associated with the alert condition.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("UserDefined.Metric"),
		},
		{
			Name:        "user_value_function",
			Description: "User defined value function associated with the alert condition.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("UserDefined.ValueFunction"),
		},
		{
			Name:        "scope",
			Description: "The scope of the alert condition.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Scope"),
		},
		{
			Name:        "gc_metric",
			Description: "The GC metric associated with the alert condition.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("GCMetric"),
		},
		{
			Name:        "violation_close_timer",
			Description: "",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "policy_id",
			Description: "Identifier for the policy alert condition belongs to.",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromQual("policy_id"),
		},
	}
}
