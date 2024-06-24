package newrelic

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// The Newrelic API key is specific to the user.
// A user can have multiple account access.
// There is no concept of organization.

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "profile_id",
			Description: "Unique identifier of the current user.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCurrentUserId,
			Transform:   transform.FromField("ID"),
		},
	}, c...)
}

func getCurrentUserIdForConnection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	user, err := getCurrentUserIdMemoized(ctx,d, h)
	if err != nil {
		return nil, err
	}
	return user.(UserInfo).ID, nil
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getCurrentUserIdMemoized = plugin.HydrateFunc(getCurrentUserIdUncached).Memoize(memoize.WithCacheKeyFunction(getCurrentUserIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getCurrentUserId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getCurrentUserIdMemoized(ctx, d, h)
}

// Build a cache key for the call to getCurrentUserIdCacheKey.
func getCurrentUserIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getCurrentUserId"
	return key, nil
}

func getCurrentUserIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	apiKey := os.Getenv("NEW_RELIC_API_KEY")

	nrConfig := GetConfig(d.Connection)
	if nrConfig.APIKey != nil {
		apiKey = *nrConfig.APIKey
	}

	url := "https://api.newrelic.com/graphql" // GraphQL endpoint for Newrelic

	// Construct the GraphQL query
	var jsonData = []byte(`
	{
		"query": "{ actor { user { name email id } } }"
	}
	`)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Key", apiKey)

	// Initialize the HTTP client
	client := &http.Client{}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := ApiResponse{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res.Data.Actor.User, nil
}

// Top level of the JSON response
type ApiResponse struct {
	Data ActorData `json:"data"`
}

// Contains the actor information, nested under "data"
type ActorData struct {
	Actor UserActor `json:"actor"`
}

// Contains the user information, nested under "actor"
type UserActor struct {
	User UserInfo `json:"user"`
}

// Contains the actual user details
type UserInfo struct {
	Email string `json:"email"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
}
