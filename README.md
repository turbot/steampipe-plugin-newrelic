![image](https://hub.steampipe.io/images/plugins/turbot/newrelic-social-graphic.png)

# NewRelic plugin for Steampipe

Use SQL to instantly query NewRelic Alerts, Events, Dashboards and more. Open source CLI. No DB required.

- **[Get started ->](https://hub.steampipe.io/plugins/turbot/newrelic)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/newrelic/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-newrelic/issues)

## Quick start

### Install

Download and install the latest NewRelic plugin:

```shell
steampipe plugin install newrelic
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/newrelic#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/newrelic#configuration).

Configure your account details in `~/.steampipe/config/newrelic.spc`:

```hcl
connection "newrelic" {
  plugin = "newrelic"

  # NewRelic API Key. Required.
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

List APM Applications on your NewRelic account:

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

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-newrelic/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [NewRelic Plugin](https://github.com/turbot/steampipe-plugin-newrelic/labels/help%20wanted)
