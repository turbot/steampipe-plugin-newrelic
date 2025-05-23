## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#55](https://github.com/turbot/steampipe-plugin-newrelic/pull/55))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#55](https://github.com/turbot/steampipe-plugin-newrelic/pull/55))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#53](https://github.com/turbot/steampipe-plugin-newrelic/pull/53))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#53](https://github.com/turbot/steampipe-plugin-newrelic/pull/53))

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#47](https://github.com/turbot/steampipe-plugin-newrelic/pull/47))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#47](https://github.com/turbot/steampipe-plugin-newrelic/pull/47))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-newrelic/blob/main/docs/LICENSE). ([#47](https://github.com/turbot/steampipe-plugin-newrelic/pull/47))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#46](https://github.com/turbot/steampipe-plugin-newrelic/pull/46))

## v0.1.2 [2023-12-06]

_Bug fixes_

- Fixed the invalid Go module path of the plugin. ([#43](https://github.com/turbot/steampipe-plugin-newrelic/pull/43))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#32](https://github.com/turbot/steampipe-plugin-newrelic/pull/32))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#30](https://github.com/turbot/steampipe-plugin-newrelic/pull/30))
- Recompiled plugin with Go version `1.21`. ([#30](https://github.com/turbot/steampipe-plugin-newrelic/pull/30))

## v0.0.1 [2023-07-13]

_What's new?_

- New tables added
  - [newrelic_account](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_account)
  - [newrelic_alert_channel](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_alert_channel)
  - [newrelic_alert_condition](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_alert_condition)
  - [newrelic_alert_event](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_alert_event)
  - [newrelic_alert_incident](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_alert_incident)
  - [newrelic_alert_policy](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_alert_policy)
  - [newrelic_apm_application](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_apm_application)
  - [newrelic_apm_application_deployment](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_apm_application_deployment)
  - [newrelic_apm_application_instance](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_apm_application_instance)
  - [newrelic_apm_application_metric](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_apm_application_metric)
  - [newrelic_apm_application_metric_data](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_apm_application_metric_data)
  - [newrelic_component](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_component)
  - [newrelic_notification_channel](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_notification_channel)
  - [newrelic_notification_destination](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_notification_destination)
  - [newrelic_plugin](https://hub.steampipe.io/plugins/turbot/newrelic/tables/newrelic_plugin)
