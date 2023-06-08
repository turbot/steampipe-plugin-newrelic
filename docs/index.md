---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/newrelic.svg"
brand_color: "#1CE783"
display_name: "NewRelic"
short_name: "newrelic"
description: "Steampipe plugin for querying NewRelic Alerts, Events and other resources."
og_description: Query NewRelic with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/newrelic-social-graphic.png"
---

# NewRelic + Steampipe

[NewRelic](https://newrelic.com/) is a SaaS providing Monitoring, Alerting, Dashboards for applications, infrastructure, etc.

[Steampipe](https://steampipe.io/) is an open source CLI for querying cloud APIs using SQL from [Turbot](https://turbot.com/)

## Documentation

- [Table definitions / examples](https://hub.steampipe.io/plugins/turbot/newrelic/tables)

## Getting Started

### Installation

```shell
steampipe plugin install newrelic
```

### Credentials

| Item | Description  |
| ---- |--------------|
| Credentials | You will require a [NewRelic API Token](https://docs.newrelic.com/docs/apis/intro-apis/new-relic-api-keys) |
| Resolution | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/newrelic.spc`).<br />2. Credentials specified in environment variables, e.g., `NEW_RELIC_API_KEY`. |

### Configuration

Configuration can take place in the config file (which takes precedence) `~/.steampipe/config/newrelic.spc` or in Environment Variables.

Environment Variables:
- `NEW_RELIC_API_KEY` for the API key (ex: `854335b43rc4t32rt3c347238v5`)

Configuration File:

```hcl
connection "newrelic" {
  plugin  = "newrelic"
  token   = "854335b43rc4t32rt3c347238v5"
}
```
