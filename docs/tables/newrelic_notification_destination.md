# Table: newrelic_notification_destination

The `newrelic_notification_destination` table can be used to obtain information about destinations for notifications for a specific New Relic account.

**You must specify `account_id` in the where or join clause.**

## Examples

### List all notification destinations for a specific account

```sql
select
  id,
  name,
  account_id,
  active,
  created_at,
  updated_at,
  updated_by,
  auth,
  is_user_authenticated,
  properties,
  status,
  type
from
  newrelic_notification_destination
where
  account_id = 21355;
```