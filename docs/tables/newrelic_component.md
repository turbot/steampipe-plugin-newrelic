---
title: "Steampipe Table: newrelic_component - Query New Relic Components using SQL"
description: "Allows users to query New Relic Components, specifically providing insights into the performance and availability of monitored applications."
---

# Table: newrelic_component - Query New Relic Components using SQL

New Relic is a software analytics company that offers cloud-based software to help website and application owners track the performances of their services. It provides SaaS Application Performance Management and Digital Intelligence products. New Relic's Components are part of its APM product, which monitor and analyze the performance of software applications.

## Table Usage Guide

The `newrelic_component` table provides insights into the components of applications monitored in New Relic. As a DevOps engineer, explore component-specific details through this table, including metrics, attributes, and associated metadata. Utilize it to uncover information about components, such as their performance, usage, and the overall health of the applications they belong to.

## Examples

### List all components
Explore the health status and metrics summary of all components in your New Relic account to gain insights into their performance and condition. This could be useful in identifying components that require attention or optimization.

```sql+postgres
select
  id,
  name
  health_status,
  summary_metrics
from
  newrelic_component;
```

```sql+sqlite
select
  id,
  name,
  health_status,
  summary_metrics
from
  newrelic_component;
```

### List components for a specific plugin
Explore the components associated with a specific plugin to gain insights into their health status and key metrics. This can be useful in assessing the performance and reliability of the plugin in question.

```sql+postgres
select
  id,
  name
  health_status,
  summary_metrics
from
  newrelic_component
where
  plugin_id = 634;
```

```sql+sqlite
select
  id,
  name,
  health_status,
  summary_metrics
from
  newrelic_component
where
  plugin_id = 634;
```