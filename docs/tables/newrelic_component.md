# Table: newrelic_component

The `newrelic_component` table can be used to obtain information on components monitored in New Relic.

## Examples

### List all components

```sql
select
  id,
  name
  health_status,
  summary_metrics
from
  newrelic_component;
```

### List components for a specific plugin

```sql
select
  id,
  name
  health_status,
  summary_metrics
from
  newrelic_component
where
  plugin_id = 634;
```