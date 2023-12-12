---
title: "Steampipe Table: newrelic_plugin - Query New Relic Plugins using SQL"
description: "Allows users to query New Relic Plugins, providing insights into plugin details such as the plugin's name, publisher, and associated metadata."
---

# Table: newrelic_plugin - Query New Relic Plugins using SQL

New Relic Plugins are tools that allow you to extend New Relic's functionality and monitor the performance of specific software services that are not covered by New Relic's standard offerings. These plugins are typically developed by third-party software vendors or by New Relic's own development team. They provide additional monitoring capabilities for a wide range of applications, databases, infrastructure components and other software services.

## Table Usage Guide

The `newrelic_plugin` table provides insights into the plugins used within New Relic. As a DevOps engineer, you can explore plugin-specific details through this table, including the plugin's name, publisher, and associated metadata. Use it to manage and monitor the performance of specific software services that are not covered by New Relic's standard offerings.

## Examples

### List all plugins
Explore the variety of plugins available, including their publishers and status, to better understand the range of tools and functionalities at your disposal. This can help in decision-making for plugin selection and usage, keeping you updated on their latest versions and changes.

```sql+postgres
select
  id,
  name,
  guid,
  description,
  publisher,
  component_agent_count,
  created_at,
  updated_at,
  short_name,
  publisher_support_url,
  publisher_about_url,
  download_url,
  published_version,
  has_unpublished_changes,
  is_public,
  summary_metrics
from
  newrelic_plugin;
```

```sql+sqlite
select
  id,
  name,
  guid,
  description,
  publisher,
  component_agent_count,
  created_at,
  updated_at,
  short_name,
  publisher_support_url,
  publisher_about_url,
  download_url,
  published_version,
  has_unpublished_changes,
  is_public,
  summary_metrics
from
  newrelic_plugin;
```

### List private plugins
Explore which NewRelic plugins are set to private, allowing you to manage and control access to your specific plugins. This can be particularly useful in maintaining security and restricting unwanted access.

```sql+postgres
select
  name,
  description,
  short_name
from
  newrelic_plugin
where
  is_public = false;
```

```sql+sqlite
select
  name,
  description,
  short_name
from
  newrelic_plugin
where
  is_public = 0;
```