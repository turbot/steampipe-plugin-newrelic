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
