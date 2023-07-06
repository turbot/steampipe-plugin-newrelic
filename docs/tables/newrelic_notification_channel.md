# Table: newrelic_notification_channel

The `newrelic_notification_channel` can be used to obtain information about notification channels associated with a specific account.

**You must specify `account_id` in the where or join clause.**

## Examples

### List all notification channels for a specific account

```sql
select
  id,
  name,
  account_id,
  active,
  created_at,
  updated_at,
  updated_by,
  destination_id,
  product,
  properties,
  status,
  type
from
  newrelic_notification_channel
where
  account_id = 21355;
```