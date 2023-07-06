# Table: newrelic_alert_channel

The `newrelic_alert_channel` table allows you to obtain information on Alert Notification Channels.

## Examples

### List all alert channels

```sql
select
  id,
  name,
  type,
  channel,
  recipients,
  teams,
  tags,
  url,
  key,
  route_key,
  base_url,
  payload_type,
  region,
  user_id,
  policies
from
  newrelic_alert_channel;
```

### List all slack alert channels

```sql
select
  name,
  channel,
  recipients,
  region,
  teams,
  tags
from
  newrelic_alert_channel
where
  type = 'slack';
```
