# Table: newrelic_apm_application

The `newrelic_apm_application` table can be used to obtain information about applications being monitored by APM.

## Examples

### List all applications monitored by apm

```sql
select
  id,
  name,
  language,
  health_status,
  reporting,
  last_reported_at,
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
  real_user_monitoring,
  server_side_config,
  servers,
  hosts,
  instances,
  alert_policy_id
from
  newrelic_apm_application;
```

### List basic information for all applications written in a specific programming language

```sql
select
  id,
  name,
  language,
  health_status,
  host_count,
  instance_count,
from
  newrelic_apm_application
where
  language = 'java';
```

### List applications using a specific named policy

```sql
select
  id,
  name,
  language,
  health_status,
  host_count,
  instance_count,
from
  newrelic_apm_application
where
  alert_policy_id = (
    select 
      id 
    from 
      newrelic_alert_policy 
    where 
      name = 'test'
  );
```