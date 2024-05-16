package newrelic

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableApmApplicationDeployment() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_apm_application_deployment",
		Description: "Obtain information about application deployments",
		List: &plugin.ListConfig{
			Hydrate: listApmApplicationDeployments,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "app_id",
					Require:   plugin.Required,
					Operators: []string{"="},
				},
			},
		},
		Columns: commonColumns(apmApplicationDeploymentColumns()),
	}
}

func listApmApplicationDeployments(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_apm_application_deployment.listApmApplicationDeployments", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	q := d.EqualsQuals
	appId := int(q["app_id"].GetInt64Value())

	plugin.Logger(ctx).Debug("newrelic_apm_application_deployment.listApmApplicationDeployments", "app.Id", appId)
	deployments, err := client.APM.ListDeploymentsWithContext(ctx, appId)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_apm_application_deployment.listApmApplicationDeployments", "query_error", err)
		return nil, fmt.Errorf("unable to obtain APM application %d deployments: %v", appId, err)
	}

	for _, deployment := range deployments {
		d.StreamListItem(ctx, deployment)
	}

	return nil, nil
}

func apmApplicationDeploymentColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier of the deployment.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "description",
			Description: "The description of the deployment.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "revision",
			Description: "The revision of the deployment.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "changelog",
			Description: "The changelog entry associated with the deployment.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "user",
			Description: "The user associated with the deployment.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "timestamp",
			Description: "The timestamp of when the deployment was deployed.",
			Type:        proto.ColumnType_TIMESTAMP,
		},
		// Links
		{
			Name:        "app_id",
			Description: "The identifier of the application the deployment is linked to.",
			Type:        proto.ColumnType_INT,
			Transform:   transform.FromField("Links.ApplicationID"),
		},
	}
}
