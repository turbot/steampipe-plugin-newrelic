# Table: newrelic_apm_application_deployment

The `newrelic_apm_application_deployment` table can be used to obtain information about deployments of applications covered by APM.

**You must specify `app_id` in there where or join clause.**

## Examples

### List all deployments

```sql
select
  id,
  description,
  revision,
  changelog,
  user,
  timestamp,
  app_id
from
  newrelic_apm_application_deployment
where
  app_id = 45;
```

### List all deployments for a specific application by name

```sql
select
  id,
  description,
  revision,
  changelog,
  user,
  timestamp,
  app_id
from
  newrelic_apm_application_deployment
where
  app_id in (
    select id 
    from newrelic_apm_application 
    where name = 'my-app'
);
```