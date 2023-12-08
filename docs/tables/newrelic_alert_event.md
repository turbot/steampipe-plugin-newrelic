---
title: "Steampipe Table: newrelic_alert_event - Query New Relic Alert Events using SQL"
description: "Allows users to query New Relic Alert Events, providing insights into the alerts triggered in the New Relic system."
---

# Table: newrelic_alert_event - Query New Relic Alert Events using SQL

New Relic Alert Events is a service within New Relic that enables monitoring and responding to issues across your applications and infrastructure. It offers a centralized way to set up and manage alerts for various resources, helping you stay informed about the health and performance of your resources and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `newrelic_alert_event` table provides insights into alert events within New Relic. As a DevOps engineer or system administrator, you can explore event-specific details through this table, including alert policies, conditions, and associated metadata. Utilize it to uncover information about alert events, such as those triggered by specific conditions, the associated policies, and the overall status of your system's health.

## Examples

### List alert events for open incidents
Determine the areas in which alert events are linked to currently open incidents. This is useful for identifying ongoing issues that require immediate attention or further investigation.

```sql+postgres
select
  id,
  timestamp,
  event_type,
  product,
  description,
  entity_id,
  entity_type,
  entity_group_id,
  priority,
  incident_id
from
  newrelic_alert_event
where
  incident_id in (
    select
      id
    from
      newrelic_alert_incident
    where
      closed_at is null
  );
```

```sql+sqlite
select
  id,
  timestamp,
  event_type,
  product,
  description,
  entity_id,
  entity_type,
  entity_group_id,
  priority,
  incident_id
from
  newrelic_alert_event
where
  incident_id in (
    select
      id
    from
      newrelic_alert_incident
    where
      closed_at is null
  );
```

### List alert events for a specific product
Explore alert events associated with a specific product to gain insights into incidents, their priority, and related entities. This can be particularly useful for troubleshooting product-related issues and understanding their impact.

```sql+postgres
select
  id,
  timestamp,
  event_type,
  product,
  description,
  entity_id,
  entity_type,
  entity_group_id,
  priority,
  incident_id
from
  newrelic_alert_event
where
  product = 'my-product';
```

```sql+sqlite
select
  id,
  timestamp,
  event_type,
  product,
  description,
  entity_id,
  entity_type,
  entity_group_id,
  priority,
  incident_id
from
  newrelic_alert_event
where
  product = 'my-product';
```