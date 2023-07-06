# Table: newrelic_apm_label

The `newrelic_apm_label` table can be used to obtain information on labels available to tag onto applications or servers.

## Examples

### List all APM labels

```sql
select
  key,
  name,
  category,
  applications,
  servers
from
  newrelic_apm_label;
```