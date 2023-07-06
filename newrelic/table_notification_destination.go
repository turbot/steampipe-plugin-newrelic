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

func tableNotificationDestination() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_notification_destination",
		Description: "Obtain information about notification destinations for a specific NewRelic account",
		List: &plugin.ListConfig{
			Hydrate: listNotificationDestinations,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "account_id",
					Require:   plugin.Required,
					Operators: []string{"="},
				},
			},
		},
		Columns: notificationDestinationColumns(),
	}
}

func listNotificationDestinations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	q := d.EqualsQuals
	accountId := int(q["account_id"].GetInt64Value())
	filters := ai.AiNotificationsDestinationFilter{}
	sorter := notifications.AiNotificationsDestinationSorter{
		Direction: notifications.AiNotificationsSortOrderTypes.ASC,
		Field:     notifications.AiNotificationsDestinationFieldsTypes.CREATED_AT,
	}
	cursor := ""

	for {
		destinations, err := client.Notifications.GetDestinationsWithContext(ctx, accountId, cursor, filters, sorter)
		if err != nil {
			return nil, fmt.Errorf("unable to obtain notification destinations for account %d: %v", accountId, err)
		}

		for _, destination := range destinations.Entities {
			d.StreamListItem(ctx, destination)
		}

		if destinations.NextCursor == "" {
			break
		}
		cursor = destinations.NextCursor
	}

	return nil, nil
}

func notificationDestinationColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "account_id", Description: "ID of the account the channel belongs to", Type: proto.ColumnType_INT},
		{Name: "active", Description: "Indicates if the destination is enabled", Type: proto.ColumnType_BOOL},
		{Name: "auth", Description: "Authentication type used for the destination", Type: proto.ColumnType_STRING},
		{Name: "is_user_authenticated", Description: "Indicates if current user is authenticated with the destination", Type: proto.ColumnType_BOOL},
		{Name: "created_at", Description: "Timestamp when the destination was created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(nrDateTransform)},
		{Name: "id", Description: "ID of the channel", Type: proto.ColumnType_STRING},
		{Name: "last_sent", Description: "Timestamp when last notification was sent to destination", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("LastSent").Transform(nrDateTransform)},
		{Name: "name", Description: "Name of the destination", Type: proto.ColumnType_STRING},
		{Name: "properties", Description: "Array of properties associated with the destination", Type: proto.ColumnType_JSON},
		{Name: "status", Description: "Status of the destination", Type: proto.ColumnType_STRING},
		{Name: "type", Description: "Type of the destination", Type: proto.ColumnType_STRING},
		{Name: "updated_at", Description: "Timestamp when destination was last updated", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("UpdatedAt").Transform(nrDateTransform)},
		{Name: "updated_by", Description: "The ID of the user which last updated the channel", Type: proto.ColumnType_INT},
	}
}
