# Table: newrelic_account

The newrelic_account table can be used to obtain information about all accounts your credentials have access to view.

## Examples

### List all accounts

```sql
select
  id,
  name,
  reporting_event_types
from
  newrelic_account;
```