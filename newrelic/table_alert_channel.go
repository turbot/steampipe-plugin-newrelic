package newrelic

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAlertChannel() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_alert_channel",
		Description: "Obtain alert channels for a given NewRelic account",
		List: &plugin.ListConfig{
			Hydrate: listAlertChannels,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAlertChannel,
		},
		Columns: alertChannelColumns(),
	}
}

func getAlertChannel(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	channelId := int(d.EqualsQuals["id"].GetInt64Value())

	c, err := client.Alerts.GetChannelWithContext(ctx, channelId)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain alert channel %d: %v", channelId, err)
	}

	return c, nil
}

func listAlertChannels(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	acs, err := client.Alerts.ListChannelsWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain alert channels: %v", err)
	}

	for _, ac := range acs {
		d.StreamListItem(ctx, ac)
	}

	return nil, nil
}

func alertChannelColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier for the alert channel.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "name",
			Description: "Name of the alert channel.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "type",
			Description: "The type of alert channel (slack, email, webhook, etc).",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "recipients",
			Description: "The configured recipients of this alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.Recipients"),
		},
		// NOTE: Omitted AuthToken as may expose secret info
		// NOTE: Omitted APIKey as may expose secret info
		{
			Name:        "teams",
			Description: "The configured teams associated with this alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.Teams"),
		},
		{
			Name:        "tags",
			Description: "The tags associated with the alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.Tags"),
		},
		{
			Name:        "url",
			Description: "The URL associated with the alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.URL"),
		},
		{
			Name:        "channel",
			Description: "The channel associated with the alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.Channel"),
		},
		{
			Name:        "key",
			Description: "The key associated with the alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.Key"),
		},
		{
			Name:        "route_key",
			Description: "The route key associated with the alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.RouteKey"),
		},
		{
			Name:        "service_key",
			Description: "The service key associated with the alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.ServiceKey"),
		},
		{
			Name:        "base_url",
			Description: "The base URL associated with the alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.BaseURL"),
		},
		// Omitted AuthUsername as may expose secret info
		// Omitted AuthPassword as may expose secret info
		{
			Name:        "payload_type",
			Description: "The type of the payload sent to the alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.PayloadType"),
		},
		{
			Name:        "region",
			Description: "The region in which the alert channel is configured.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.Region"),
		},
		{
			Name:        "user_id",
			Description: "The identifier of the user whom created the alert channel.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Configuration.UserID"),
		},
		// Omitted Headers as may contain secret info
		// Omitted Payload as may contain secret info
		{
			Name:        "policies",
			Description: "An array of policy identifiers that link the alert channel to a policy.",
			Type:        proto.ColumnType_JSON,
			Transform:   transform.FromField("Links.PolicyIDs"),
		},
	}
}
