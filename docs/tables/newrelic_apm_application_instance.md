# Table: newrelic_apm_application_instance

The `newrelic_apm_application_instance` table can use used to obtain information about instances of the applications covered by APM.

**You must specify `app_id` in there where or join clause.**

## Examples

### List all application instances monitored by apm

```sql
select
  id,
  application_name,
  host,
  port,
  language,
  health_status,
  response_time,
  throughput,
  error_rate,
  apdex_target,
  apdex_score,
  host_count,
  instance_count,
  concurrent_instance_count,
  end_user_response_time,
  end_user_throughput,
  end_user_apdex_target,
  end_user_apdex_score,
  end_user_apdex_threshold,
  app_id,
  app_host
from
  newrelic_apm_application_instance
where
  app_id = 45;
```