---
title: "Steampipe Table: newrelic_apm_application_metric_data - Query New Relic APM Application Metric Data using SQL"
description: "Allows users to query New Relic APM Application Metric Data, specifically providing insights into application performance and potential areas of improvement."
---

# Table: newrelic_apm_application_metric_data - Query New Relic APM Application Metric Data using SQL

New Relic APM (Application Performance Monitoring) is a service that provides real-time monitoring and detailed performance analytics for your applications. It allows you to track transactions, errors, and other key metrics, enabling you to identify and diagnose performance issues quickly. With New Relic APM, you can gain a deep understanding of how your applications are performing in production, and where they can be optimized.

## Table Usage Guide

The `newrelic_apm_application_metric_data` table provides insights into application performance metrics within New Relic's Application Performance Monitoring (APM). As a DevOps engineer or application developer, explore application-specific details through this table, including transaction times, error rates, and other key performance indicators. Utilize it to uncover information about application performance, such as slow transactions, high error rates, and other potential areas of improvement.

## Examples

### List metrics for an application after a certain date
Explore the performance metrics of a specific application for the past year. This is useful in identifying trends or issues that have arisen over time, helping to inform future development and troubleshooting efforts.

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