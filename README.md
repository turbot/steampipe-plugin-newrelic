![image](https://hub.steampipe.io/images/plugins/turbot/newrelic-social-graphic.png)

# New Relic plugin for Steampipe

Use SQL to instantly query New Relic Alerts, Events, Dashboards and more. Open source CLI. No DB required.

- **[Get started ->](https://hub.steampipe.io/plugins/turbot/newrelic)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/newrelic/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-newrelic/issues)

## Quick start

### Install

Download and install the latest New Relic plugin:

```shell
steampipe plugin install newrelic
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/newrelic#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/newrelic#configuration).

Configure your account details in `~/.steampipe/config/newrelic.spc`:

```hcl
connection "newrelic" {
  plugin = "newrelic"

  # Authentication information
  api_key = "NRAK-XX0X0XX00XXXX0000XXXXXXXXX0X"
  region = "us"
}
```

or through environment variables:

```shell
export NEW_RELIC_API_KEY=NRAK-XX0X0XX00XXXX0000XXXXXXXXX0X
export NEW_RELIC_REGION=us
```

Run steampipe:

```shell
steampipe query
```

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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/overview) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/overview) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```shell
git clone https://github.com/turbot/steampipe-plugin-newrelic.git
cd steampipe-plugin-newrelic
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```shell
make
```

Configure the plugin:

```shell
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/newrelic.spc
```

Try it!

```shell
steampipe query
> .inspect newrelic
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [New Relic Plugin](https://github.com/turbot/steampipe-plugin-newrelic/labels/help%20wanted)
