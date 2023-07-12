package newrelic

import (
	"context"
	"fmt"

	"github.com/newrelic/newrelic-client-go/v2/pkg/accounts"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAccount() *plugin.Table {
	return &plugin.Table{
		Name:        "newrelic_account",
		Description: "Obtain accounts visible to your user from New Relic.",
		List: &plugin.ListConfig{
			Hydrate: listAccounts,
		},
		Columns: accountColumns(),
	}
}

func listAccounts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_account.listAccounts", "connection_error", err)
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	params := accounts.ListAccountsParams{
		Scope: &accounts.RegionScopeTypes.GLOBAL,
	}

	plugin.Logger(ctx).Debug("newrelic_account.listAccounts", "params.Scope", *params.Scope)
	as, err := client.Accounts.ListAccountsWithContext(ctx, params)
	if err != nil {
		plugin.Logger(ctx).Error("newrelic_account.listAccounts", "query_error", err)
		return nil, fmt.Errorf("unable to obtain accounts: %v", err)
	}

	for _, a := range as {
		d.StreamListItem(ctx, a)
	}

	return nil, nil
}

func accountColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Unique identifier for the account.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "name",
			Description: "Name of the account.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "reporting_event_types",
			Description: "An array of event types that are currently reporting in the account.",
			Type:        proto.ColumnType_JSON,
		},
	}
}
