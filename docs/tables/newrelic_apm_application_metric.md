---
title: "Steampipe Table: newrelic_apm_application_metric - Query New Relic APM Application Metrics using SQL"
description: "Allows users to query APM Application Metrics in New Relic, providing insights into application performance and potential anomalies."
---

# Table: newrelic_apm_application_metric - Query New Relic APM Application Metrics using SQL

New Relic's Application Performance Monitoring (APM) is a cloud-based service that provides insights into your software applications' performance and usage. It helps you monitor, troubleshoot, and optimize your entire software stack. APM Application Metrics provide detailed performance data of your applications, which is crucial for maintaining optimal application performance and enhancing user experience.

## Table Usage Guide

The `newrelic_apm_application_metric` table provides insights into the performance metrics of applications monitored by New Relic's APM service. As a DevOps engineer or application developer, you can explore detailed metrics through this table, including response times, throughput, and error rates. Utilize it to monitor application performance, identify potential issues, and optimize your applications for better user experience.

**Important Notes**
- You must specify the `app_id` in the `where` clause to query this table.

## Examples

### List all metrics available for a specific application
Explore which metrics are available for a specific application to understand its performance and usage. This can help in identifying areas that may require attention or optimization.

```sql+postgres
select
  name,
  values
from
  newrelic_apm_application_metric
where
  app_id = 45;
```

```sql+sqlite
select
  name,
  values
from
  newrelic_apm_application_metric
where
  app_id = 45;
```

### List all available cpu based metrics for all applications
Discover the segments that provide a comprehensive view of CPU-related metrics across all applications. This is beneficial for performance monitoring and optimization, helping to identify areas that may require attention or improvement.

```sql+postgres
select
  a.name as app,
  m.name as metric,
  m.values as values
from
  newrelic_apm_application_metric m,
  newrelic_apm_application a
where
  m.app_id = a.id
and m.name ilike '%cpu%';
```

```sql+sqlite
select
  a.name as app,
  m.name as metric,
  m.values as values
from
  newrelic_apm_application_metric m,
  newrelic_apm_application a 
where
  m.app_id = a.id
and m.name like '%cpu%';
```