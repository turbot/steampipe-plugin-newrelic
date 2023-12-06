---
title: "Steampipe Table: newrelic_apm_application_instance - Query New Relic APM Application Instances using SQL"
description: "Allows users to query New Relic APM Application Instances, specifically providing metrics about individual instances of applications monitored by New Relic APM."
---

# Table: newrelic_apm_application_instance - Query New Relic APM Application Instances using SQL

New Relic's Application Performance Monitoring (APM) is a software that provides real-time monitoring and detailed performance analytics of software applications. It helps in identifying performance issues, understanding dependencies, and improving software performance. An Application Instance in New Relic APM represents a single running instance of your application, providing in-depth performance metrics for that instance.

## Table Usage Guide

The `newrelic_apm_application_instance` table provides insights into individual instances of applications monitored by New Relic APM. As a DevOps engineer or application developer, you can use this table to explore detailed metrics about each application instance, including throughput, response time, and error rate. Utilize it to identify performance bottlenecks, understand dependencies, and improve the overall performance of your applications.

**Important Notes**
- You must specify the `app_id` in the `where` clause to query this table.

## Examples

### List all application instances monitored by apm for a specific application
Determine the performance and health status of a specific application by analyzing its response time, error rate, and other key metrics. This can aid in identifying any potential issues or bottlenecks that may be affecting the application's overall performance.

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

### Obtain information on instances of all applications
Explore which applications have the most instances to understand resource allocation and usage patterns. This can help in optimizing resource distribution and identifying potential bottlenecks.

```sql
select
  i.application_name,
  sum(i.instance_count) as instances
from
  newrelic_apm_application_instance i,
  newrelic_apm_application a
where
  i.app_id = a.id
group by
  i.application_name;
```