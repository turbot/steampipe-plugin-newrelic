# Table: newrelic_alert_condition

The `newrelic_alert_condition` table can be used to obtain information on conditions why alerts were raised.

## Examples

### List all alert conditions

```sql
select
  id,
  name,
  type,
  enabled,
  entities,
  metric,
  runbook_url,
  terms,
  user_metric,
  user_value_function,
  scope,
  gc_metric,
  violation_close_timer,
  policy_id
from
  newrelic_alert_condition;
```

### List only enabled alert conditions

```sql
select
  name,
  type,
  metric,
  user_metric,
  user_value_function,
  gc_metric
from
  newrelic_alert_condition
where
  enabled = true;
```

### List only alert conditions for a specific metric

```sql
select
  name,
  type,
  metric,
  user_metric,
  user_value_function,
  gc_metric
from
  newrelic_alert_condition
where
  metric = 'cpu_percentage';
```
