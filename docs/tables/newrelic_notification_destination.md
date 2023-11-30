---
title: "Steampipe Table: newrelic_notification_destination - Query New Relic Notification Destinations using SQL"
description: "Allows users to query New Relic Notification Destinations, providing insights into the notification channels configured in New Relic."
---

# Table: newrelic_notification_destination - Query New Relic Notification Destinations using SQL

New Relic Notification Destinations are channels configured in New Relic to receive notifications. These channels can be of different types like email, Slack, PagerDuty, etc. They are a crucial part of New Relic's Alert system, tying together the alert policies and the users who need to be notified.

## Table Usage Guide

The `newrelic_notification_destination` table provides insights into the notification channels configured in New Relic. As a DevOps engineer, you can explore details about these channels, including their types and configurations, through this table. Utilize it to manage and optimize your alert system by understanding the distribution of notifications across different channels.

## Examples

### List all notification destinations for a specific account
Determine the areas in which notifications are being sent for a specific user account. This can be useful to audit and manage where notifications are directed, ensuring the right individuals or teams are alerted.

```sql
select
  id,
  name,
  account_id,
  active,
  created_at,
  updated_at,
  updated_by,
  auth,
  is_user_authenticated,
  properties,
  status,
  type
from
  newrelic_notification_destination
where
  account_id = 21355;
```