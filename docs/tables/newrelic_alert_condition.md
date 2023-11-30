---
title: "Steampipe Table: newrelic_alert_condition - Query New Relic Alert Conditions using SQL"
description: "Allows users to query New Relic Alert Conditions, specifically the condition's name, status, and severity, providing insights into application performance and potential anomalies."
---

# Table: newrelic_alert_condition - Query New Relic Alert Conditions using SQL

New Relic Alert Conditions is a feature within the New Relic platform that allows you to monitor and respond to issues across your applications and infrastructure. It provides a centralized way to set up and manage alerts based on various conditions, such as error rates, response times, and server load. New Relic Alert Conditions helps you stay informed about the health and performance of your applications and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `newrelic_alert_condition` table provides insights into alert conditions within the New Relic platform. As a DevOps engineer, explore condition-specific details through this table, including condition names, statuses, and severities. Utilize it to uncover information about conditions, such as those related to high error rates, slow response times, and overloaded servers, helping you to maintain optimal application performance.

## Examples

### List all alert conditions
Explore all the active alert conditions in your system to understand their configuration and status. This allows you to identify potential areas of concern and take proactive measures to prevent system failures.

```sql
select
  id,
  name,
  type,
  enabled,
  entities,
  metric,
  runbook_url,
  terms,
  user_metric,
  user_value_function,
  scope,
  gc_metric,
  violation_close_timer,
  policy_id
from
  newrelic_alert_condition;
```

### List only enabled alert conditions
Discover the segments that are active within the alert conditions. This can help you focus on the areas that are currently in use and require monitoring, ensuring efficient resource allocation.

```sql
select
  name,
  type,
  metric,
  user_metric,
  user_value_function,
  gc_metric
from
  newrelic_alert_condition
where
  enabled = true;
```

### List only alert conditions for a specific metric
Determine the areas in which specific alert conditions are linked to CPU percentage metrics. This allows for targeted monitoring and management of system performance.

```sql
select
  name,
  type,
  metric,
  user_metric,
  user_value_function,
  gc_metric
from
  newrelic_alert_condition
where
  metric = 'cpu_percentage';
```