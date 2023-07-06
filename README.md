![image](https://hub.steampipe.io/images/plugins/turbot/newrelic-social-graphic.png)

# NewRelic plugin for Steampipe

Use SQL to instantly query NewRelic Alerts, Events, Dashboards and more. Open source CLI. No DB required.

- **[Get started ->](https://hub.steampipe.io/plugins/turbot/newrelic)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/newrelic/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-newrelic/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install newrelic
```

Setup the configuration:

```shell
vi ~/.steampipe/config/newrelic.spc
```

or set the following Environment Variables

- `NEW_RELIC_API_KEY` : The API Key to use for the New Relic API.
- `NEW_RELIC_REGION` : The region to use `us` or `eu`.

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

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-pagerduty/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [NewRelic Plugin](https://github.com/turbot/steampipe-plugin-newrelic/labels/help%20wanted)
