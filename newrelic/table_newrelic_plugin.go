package newrelic

import (
	"context"
	"fmt"

	"github.com/newrelic/newrelic-client-go/v2/pkg/plugins"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePlugin() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_plugin",
		Description: "Obtain information about plugins installed in your New Relic account",
		List: &plugin.ListConfig{
			Hydrate: listPlugins,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getPlugins,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: pluginColumns(),
	}
}

func getPlugins(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_plugin.getPlugins", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	pluginId := int(d.EqualsQuals["id"].GetInt64Value())
	params := plugins.GetPluginParams{Detailed: true}

	plugin.Logger(ctx).Debug("newrelic_plugin.getPlugins", "plugin.Id", pluginId, "params.Detailed", params.Detailed)
	p, err := client.Plugins.GetPlugin(pluginId, &params)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_plugin.getPlugins", "query_error", err)
		return nil, fmt.Errorf("unable to obtain plugin %d: %v", pluginId, err)
	}

	return p, nil
}

func listPlugins(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_plugin.listPlugins", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	params := plugins.ListPluginsParams{
		Detailed: true,
	}

	plugin.Logger(ctx).Debug("newrelic_plugin.listPlugins", "params.Detailed", params.Detailed)
	ps, err := client.Plugins.ListPlugins(&params)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_plugin.listPlugins", "query_error", err)
		return nil, fmt.Errorf("unable to obtain plugins: %v", err)
	}

	for _, p := range ps {
		d.StreamListItem(ctx, p)
	}

	return nil, nil
}

func pluginColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier for the plugin.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "name",
			Description: "Name of the plugin.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "guid",
			Description: "The GUID associated with the plugin.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("GUID"),
		},
		{
			Name:        "description",
			Description: "The description of the plugin.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Details.Description"),
		},
		{
			Name:        "publisher",
			Description: "The publisher of the plugin.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "component_agent_count",
			Description: "Count of component agents associated with the plugin.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "created_at",
			Description: "Timestamp at which the plugin record was created.",
			Type:        proto.ColumnType_STRING, // TODO: See if can be converted to timestamp from actual value (SDK is string)
			Transform:   transform.FromField("Details.CreatedAt"),
		},
		{
			Name:        "updated_at",
			Description: "Timestamp at which the plugin was last updated.",
			Type:        proto.ColumnType_STRING, // TODO: See if can be converted to timestamp from actual value (SDK is string)
			Transform:   transform.FromField("Details.UpdatedAt"),
		},
		{
			Name:        "short_name",
			Description: "The short version of the plugin name.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Details.ShortName"),
		},
		{
			Name:        "publisher_support_url",
			Description: "The support URL from the plugin's publisher.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Details.PublisherSupportURL"),
		},
		{
			Name:        "publisher_about_url",
			Description: "The about URL from the plugin's publisher.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Details.PublisherAboutURL"),
		},
		{
			Name:        "download_url",
			Description: "The download URL for the plugin.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Details.DownloadURL"),
		},
		{
			Name:        "published_version",
			Description: "The published version of the plugin.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Details.PublishedVersion"),
		},
		{
			Name:        "has_unpublished_changes",
			Description: "Indicates if the plugin has unpublished changes.",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Details.HasUnpublishedChanges"),
		},
		{
			Name:        "is_public",
			Description: "Indicates if the plugin is publicly available.",
			Type:        proto.ColumnType_BOOL,
			Transform:   transform.FromField("Details.IsPublic"),
		},
		{
			Name:        "summary_metrics",
			Description: "An array of metric summaries associated with the plugin.",
			Type:        proto.ColumnType_JSON,
		},
	}
}
