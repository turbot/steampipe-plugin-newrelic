# Table: newrelic_plugin

The `newrelic_plugin` table can used to obtain information on plugins.

## Examples

### List all plugins

```sql
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

```sql
select
  name,
  description,
  short_name
from
  newrelic_plugin
where
  is_public = false;
```