# Table: newrelic_alert_policy

The `newrelic_alert_policy` can be used to obtain information about policies for alerts.

## Examples

### List all alert policies

```sql
select
  id,
  name,
  incident_preference,
  created_at,
  updated_at
from
  newrelic_alert_policy;
```

### Obtain a single alert policy by id

```sql
select
  id,
  name,
  incident_preference,
  created_at,
  updated_at
from
  newrelic_alert_policy
where
  id = 142354;
```