---
title: "Steampipe Table: newrelic_apm_application_deployment - Query New Relic APM Application Deployments using SQL"
description: "Allows users to query New Relic APM Application Deployments, providing insights into deployment details such as revision, timestamp, user, and change log."
---

# Table: newrelic_apm_application_deployment - Query New Relic APM Application Deployments using SQL

New Relic Application Performance Monitoring (APM) is a tool that provides real-time monitoring and alerting of application performance and health. It offers end-to-end transaction tracing, allowing you to isolate and solve application issues before they affect customers. With New Relic APM, you can track deployments, identify errors, and drill down into detailed performance data.

## Table Usage Guide

The `newrelic_apm_application_deployment` table provides insights into application deployments within New Relic APM. As a DevOps engineer, explore deployment-specific details through this table, including revision, timestamp, user, and change log. Utilize it to track the history of application deployments, identify changes, and understand the impact of each deployment on application performance.

**Important Notes**
- You must specify the `app_id` in the `where` clause to query this table.

## Examples

### List all deployments
Explore all deployments related to a specific application, including their individual details and timestamps. This aids in tracking the history of deployments and identifying any changes made by users over time.

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
This example allows you to identify all deployments associated with a particular application, helping you track changes and updates made over time. This can be useful in managing application versions and understanding the history of application modifications.

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