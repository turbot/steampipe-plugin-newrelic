package newrelic

import (
	"context"
	"fmt"
	"github.com/newrelic/newrelic-client-go/v2/pkg/nrtime"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"time"
)

func nrDateTransform(_ context.Context, input *transform.TransformData) (interface{}, error) {
	if input.Value == nil {
		return nil, nil
	}

	x := string(input.Value.(nrtime.DateTime))
	if x == "" {
		return nil, nil
	}

	return time.Parse(time.RFC3339, x)
}

func epochTransform(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	if input.Value == nil {
		return nil, nil
	}

	x := fmt.Sprintf("%v", input.Value)
	plugin.Logger(ctx).Warn(x)
	return time.Parse("2006-01-02 15:04:05.999 +0000 UTC", x)
}
