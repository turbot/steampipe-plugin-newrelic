---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/newrelic.svg"
brand_color: "#1CE783"
display_name: "New Relic"
short_name: "newrelic"
description: "Steampipe plugin for querying New Relic Alerts, Events and other resources."
og_description: Query New Relic with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/newrelic-social-graphic.png"
---

# New Relic + Steampipe

[New Relic](https://newrelic.com/) is a SaaS providing Monitoring, Alerting, Dashboards for applications, infrastructure, etc.

[Steampipe](https://steampipe.io/) is an open source CLI for querying cloud APIs using SQL from [Turbot](https://turbot.com/)

List APM Applications on your New Relic account:

```sql
select
  id,
  name,
  error_rate,
  health_status,
  response_time
from
  newrelic_apm_application;
```

```
+-----------+------+------------+---------------+---------------+
| id        | name | error_rate | health_status | response_time |
+-----------+------+------------+---------------+---------------+
| 511153982 | test | 0          | gray          | 0             |
+-----------+------+------------+---------------+---------------+
```

## Documentation

- [Table definitions / examples →](https://hub.steampipe.io/plugins/turbot/newrelic/tables)

## Quick start

### Install

Download and install the latest New Relic plugin:

```shell
steampipe plugin install newrelic
```

### Credentials

| Item | Description                                                                                                                                                                                              |
| ---- |----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | You will require a [New Relic API Token](https://docs.newrelic.com/docs/apis/intro-apis/new-relic-api-keys)                                                                                               |
| Permissions | User API Keys are associated with a user account, they have the same permissions as the user which may mean they can access multiple accounts.                                                              |
| Radius | Each connection represents one New Relic user, this can be across multiple accounts if the user has permissions on multiple accounts. |                                                                    |
| Resolution | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/newrelic.spc`).<br />2. Credentials specified in environment variables, e.g., `NEW_RELIC_API_KEY` and `NEW_RELIC_REGION`. |

### Configuration

Installing the latest New Relic plugin will create a config file (`~/.steampipe/config/newrelic.spc`) with a single connection named `newrelic`:

Configure your account details in `~/.steampipe/config/newrelic.spc`:

```hcl
connection "newrelic" {
plugin = "newrelic"

    # New Relic API Key. Required.
    # This can also be set via the 'NEW_RELIC_API_KEY' environment variable.
    # api_key = "NRAK-XX0X0XX00XXXX0000XXXXXXXXX0X"

    # New Relic Region - valid values are 'us' or 'eu' (default, if not chosen, is 'us'). Optional.
    # This can also be set via the 'NEW_RELIC_REGION' environment variable.
    # region = "us"
}
```

Alternatively, you can also use the standard New Relic environment variables to configure your credentials **only if other arguments (`api_key`, `region`) are not specified** in the connection:

```shell
export NEW_RELIC_API_KEY=NRAK-XX0X0XX00XXXX0000XXXXXXXXX0X
export NEW_RELIC_REGION=us
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-newrelic
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
