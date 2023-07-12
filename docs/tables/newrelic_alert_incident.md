# Table: newrelic_alert_incident

The `newrelic_alert_incident` can be used to obtain information on incidents that an alert was raised for.

## Examples

### List all alert incidents

```sql
select 
  id,
  opened_at,
  closed_at,
  incident_preference,
  policy_id,
  violations
from
  newrelic_alert_incident;
```

### List alert incidents with policy names

```sql
select
  i.id,
  p.name as policy,
  i.opened_at,
  i.closed_at
from
  newrelic_alert_incident i,
  newrelic_alert_policy p
where
  i.policy_id = p.id;
```

### List open alert incidents

```sql
select 
  id,
  opened_at,
  closed_at,
  incident_preference,
  policy_id,
  violations
from
  newrelic_alert_incident
where
  closed_at is null;
```