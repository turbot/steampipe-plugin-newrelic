# Table: newrelic_apm_application_metric

The `newrelic_apm_application_metric` table can be used to obtain information on available metrics for a specific application.

**You must specify `app_id` in there where or join clause.**

## Examples

### List all metrics available for a specific application

```sql
select
  name,
  values
from
  newrelic_apm_application_metric
where
  app_id = 45;
```