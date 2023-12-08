---
title: "Steampipe Table: newrelic_alert_policy - Query New Relic Alert Policies using SQL"
description: "Allows users to query New Relic Alert Policies, specifically providing insights into the conditions and thresholds that trigger alerts."
---

# Table: newrelic_alert_policy - Query New Relic Alert Policies using SQL

New Relic Alert Policies is a feature within New Relic's observability platform that allows you to set up and manage alert conditions for your applications and infrastructure. It provides a centralized way to monitor and respond to issues, enabling you to stay informed about the health and performance of your resources and take appropriate actions when predefined thresholds are met. New Relic Alert Policies helps in maintaining the stability of your systems by providing real-time alerting capabilities.

## Table Usage Guide

The `newrelic_alert_policy` table provides insights into Alert Policies within New Relic's observability platform. As a DevOps engineer or system administrator, you can explore policy-specific details through this table, including the conditions that trigger alerts and the thresholds that have been set. Utilize it to monitor the health and performance of your applications and infrastructure, and to respond promptly to any issues that may arise.

## Examples

### List all alert policies
Explore the various alert policies, including when they were created and last updated, to better manage and respond to incidents in your New Relic environment. This can help improve your incident response strategy by keeping you informed about all the existing alert policies.

```sql+postgres
select
  id,
  name,
  incident_preference,
  created_at,
  updated_at
from
  newrelic_alert_policy;
```

```sql+sqlite
select
  id,
  name,
  incident_preference,
  created_at,
  updated_at
from
  newrelic_alert_policy;
```

### Obtain a single alert policy by id
Gain insights into a specific alert policy by using its unique identifier. This is useful in understanding the alert's preferences and tracking when it was created or last updated.

```sql+postgres
select
  id,
  name,
  incident_preference,
  created_at,
  updated_at
from
  newrelic_alert_policy
where
  id = 142354;
```

```sql+sqlite
select
  id,
  name,
  incident_preference,
  created_at,
  updated_at
from
  newrelic_alert_policy
where
  id = 142354;
```