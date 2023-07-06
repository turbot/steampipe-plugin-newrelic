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
