# Table: newrelic_alert_event

The `newrelic_alert_event` table can be used to obtain information on alert events that have occurred.

## Examples

### List alert events for open incidents

```sql
select
  id,
  timestamp,
  event_type,
  product,
  description,
  entity_id,
  entity_type,
  entity_group_id,
  priority,
  incident_id
from
  newrelic_alert_event
where
  incident_id in (
    select
      id
    from
      newrelic_alert_incident
    where
      closed_at is null
  );
```

### List alert events for a specific product

```sql
select
  id,
  timestamp,
  event_type,
  product,
  description,
  entity_id,
  entity_type,
  entity_group_id,
  priority,
  incident_id
from
  newrelic_alert_event
where
  product = 'my-product';
```
