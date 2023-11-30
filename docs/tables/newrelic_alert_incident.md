---
title: "Steampipe Table: newrelic_alert_incident - Query New Relic Alert Incidents using SQL"
description: "Allows users to query New Relic Alert Incidents, providing comprehensive details about the incidents triggered in the New Relic account."
---

# Table: newrelic_alert_incident - Query New Relic Alert Incidents using SQL

New Relic Alerts is a flexible, centralized notification system that unlocks the operational potential of New Relic. Alerts provides a streamlined interface for managing alert policies and conditions, allowing users to evolve their alerting strategy to suit changing needs. It offers fine-grained control over alert conditions, ensuring that the right people are notified at the right time.

## Table Usage Guide

The `newrelic_alert_incident` table provides insights into alert incidents within New Relic. As a DevOps engineer, explore incident-specific details through this table, including the severity, duration, and associated policies. Utilize it to uncover information about incidents, such as those with high severity, long duration, and the policies that triggered them.

## Examples

### List all alert incidents
Explore all alert incidents to identify when they were opened and closed, their preferences, associated policies, and any violations. This is useful for gaining insights into potential issues and ensuring appropriate alert policies are in place.

```sql
select 
  id,
  opened_at,
  closed_at,
  incident_preference,
  policy_id,
  violations
from
  newrelic_alert_incident;
```

### List alert incidents with policy names
Determine the areas in which alert incidents are linked with specific policy names. This can be useful to understand the timeline of incidents and their corresponding policies, helping to manage and mitigate risks effectively.

```sql
select
  i.id,
  p.name as policy,
  i.opened_at,
  i.closed_at
from
  newrelic_alert_incident i,
  newrelic_alert_policy p
where
  i.policy_id = p.id;
```

### List open alert incidents
Explore which alert incidents are currently open in your NewRelic monitoring system. This can help you quickly assess active issues and prioritize your response efforts.

```sql
select 
  id,
  opened_at,
  closed_at,
  incident_preference,
  policy_id,
  violations
from
  newrelic_alert_incident
where
  closed_at is null;
```