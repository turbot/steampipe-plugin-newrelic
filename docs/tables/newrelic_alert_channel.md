---
title: "Steampipe Table: newrelic_alert_channel - Query New Relic Alert Channels using SQL"
description: "Allows users to query New Relic Alert Channels, specifically the configuration and status of each alert channel."
---

# Table: newrelic_alert_channel - Query New Relic Alert Channels using SQL

New Relic Alert Channels are part of the New Relic Alerts service that helps you monitor and respond to changes in your system. It provides a centralized way to manage alert notifications for various resources, enabling rapid response to system issues. Alert Channels allow you to define where and how you want to receive alert notifications.

## Table Usage Guide

The `newrelic_alert_channel` table provides insights into alert channels within New Relic Alerts service. As a DevOps engineer, explore channel-specific details through this table, including configuration, status, and associated metadata. Utilize it to manage and monitor your alert channels, ensuring timely and appropriate notifications for system issues.

## Examples

### List all alert channels
Explore various alert channels to understand their types, recipients, associated teams, and other related details. This is useful in gaining insights into how alerts are managed and disseminated within your organization.

```sql+postgres
select
  id,
  name,
  type,
  channel,
  recipients,
  teams,
  tags,
  url,
  key,
  route_key,
  base_url,
  payload_type,
  region,
  user_id,
  policies
from
  newrelic_alert_channel;
```

```sql+sqlite
select
  id,
  name,
  type,
  channel,
  recipients,
  teams,
  tags,
  url,
  key,
  route_key,
  base_url,
  payload_type,
  region,
  user_id,
  policies
from
  newrelic_alert_channel;
```

### List all slack alert channels
Explore all the alert channels set up on Slack to gain insights into the distribution of alerts across different teams and regions. This can help in managing the alert system more efficiently and ensuring all relevant parties are notified.

```sql+postgres
select
  name,
  channel,
  recipients,
  region,
  teams,
  tags
from
  newrelic_alert_channel
where
  type = 'slack';
```

```sql+sqlite
select
  name,
  channel,
  recipients,
  region,
  teams,
  tags
from
  newrelic_alert_channel
where
  type = 'slack';
```