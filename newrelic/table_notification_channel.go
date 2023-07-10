package newrelic

import (
	"context"
	"fmt"
	"github.com/newrelic/newrelic-client-go/v2/pkg/ai"
	"github.com/newrelic/newrelic-client-go/v2/pkg/notifications"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNotificationChannel() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_notification_channel",
		Description: "Obtain information about notification channels for a specific NewRelic account",
		List: &plugin.ListConfig{
			Hydrate: listNotificationChannels,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "account_id",
					Require:   plugin.Required,
					Operators: []string{"="},
				},
			},
		},
		Columns: notificationChannelColumns(),
	}
}

func listNotificationChannels(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	q := d.EqualsQuals
	accountId := int(q["account_id"].GetInt64Value())
	filters := ai.AiNotificationsChannelFilter{}
	sorter := notifications.AiNotificationsChannelSorter{
		Direction: notifications.AiNotificationsSortOrderTypes.ASC,
		Field:     notifications.AiNotificationsChannelFieldsTypes.CREATED_AT,
	}
	cursor := ""

	for {
		channels, err := client.Notifications.GetChannelsWithContext(ctx, accountId, cursor, filters, sorter)
		if err != nil {
			return nil, fmt.Errorf("unable to obtain notification channels for aaccount %d: %v", accountId, err)
		}

		for _, channel := range channels.Entities {
			d.StreamListItem(ctx, channel)
		}

		if channels.NextCursor == "" {
			break
		}
		cursor = channels.NextCursor
	}

	return nil, nil
}

func notificationChannelColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "account_id", Description: "ID of the account the channel belongs to.", Type: proto.ColumnType_INT},
		{Name: "active", Description: "Indicates if the channel is enabled.", Type: proto.ColumnType_BOOL},
		{Name: "created_at", Description: "Timestamp when the channel was created.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(nrDateTransform)},
		{Name: "destination_id", Description: "The ID of the notification destination.", Type: proto.ColumnType_STRING},
		{Name: "id", Description: "ID of the channel.", Type: proto.ColumnType_STRING},
		{Name: "name", Description: "Name of the channel.", Type: proto.ColumnType_STRING},
		{Name: "product", Description: "Name of the product the channel is created for.", Type: proto.ColumnType_STRING},
		{Name: "properties", Description: "Array of properties associated with the channel.", Type: proto.ColumnType_JSON},
		{Name: "status", Description: "Status of the channel.", Type: proto.ColumnType_STRING},
		{Name: "type", Description: "Type of the channel", Type: proto.ColumnType_STRING},
		{Name: "updated_at", Description: "Timestamp when  channel was last updated.", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("UpdatedAt").Transform(nrDateTransform)},
		{Name: "updated_by", Description: "The ID of the user which last updated the channel.", Type: proto.ColumnType_INT},
	}
}
