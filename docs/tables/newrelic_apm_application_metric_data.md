# Table: newrelic_apm_application_metric_data

The `newrelic_apm_application_metric_data` table can be used to obtain actual APM metric data for a specific application.

**You must specify `app_id` in there where or join clause.**

## Examples

### List metrics for an application between two time periods

```sql
select
from
  newrelic_apm_application_metric_data
where
  app_id = 45
and
  from = (now() - interval '1 DAY')::date
and
 to = now()::date;
```