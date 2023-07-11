# Table: newrelic_apm_application_metric_data

The `newrelic_apm_application_metric_data` table can be used to obtain actual APM metric data for a specific application.

**You must specify `app_id` in there where or join clause.**

> Note: Columns `From` and `To` must be double-quoted when used explicitly in the `select` or `where` clauses.

## Examples

### List metrics for an application after a certain date

```sql
select
  name,
  values
from
  newrelic_apm_application_metric_data
where
  app_id = 45
and
  "from" >= (now() - interval '1 YEAR')::date;
```
