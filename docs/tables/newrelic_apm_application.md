---
title: "Steampipe Table: newrelic_apm_application - Query New Relic APM Applications using SQL"
description: "Allows users to query New Relic APM Applications, specifically the details related to application performance, providing insights into application health and potential performance bottlenecks."
---

# Table: newrelic_apm_application - Query New Relic APM Applications using SQL

New Relic APM is a software application that helps monitor and manage the performance of your applications. It provides detailed performance metrics for every aspect of your environment. With New Relic APM, you can understand how your applications are performing in real-time, troubleshoot issues, and improve application performance over time.

## Table Usage Guide

The `newrelic_apm_application` table provides insights into application performance within New Relic APM. As a DevOps engineer, explore application-specific details through this table, including response times, throughput, error rates, and Apdex score. Utilize it to uncover information about applications, such as those with high error rates, poor response times, and to verify Apdex scores.

## Examples

### List all applications monitored by apm
Explore the performance and health status of all applications under monitoring. This query helps in gaining insights into application response times, error rates, and apdex scores, allowing you to assess their overall performance and troubleshoot any potential issues.

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
Explore applications written in a specific programming language to gain insights into their health status and the count of hosts and instances. This helps in assessing the performance and health of applications and can guide in resource allocation.

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
Explore which applications are utilizing a specific policy to gain insights into the policy's impact and effectiveness. This query is particularly useful for managing and optimizing policy usage across different applications.

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