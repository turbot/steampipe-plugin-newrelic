---
title: "Steampipe Table: newrelic_account - Query New Relic Accounts using SQL"
description: "Allows users to query New Relic Accounts, specifically providing information about the account's ID, name, product, and other related details."
---

# Table: newrelic_account - Query New Relic Accounts using SQL

New Relic is a performance management solution that enables developers to diagnose and fix application performance problems in real time. It provides visibility into the performance and usage of your software applications. New Relic helps you to understand how your applications are performing and where bottlenecks may be occurring.

## Table Usage Guide

The `newrelic_account` table provides insights into accounts within New Relic. As a developer or DevOps engineer, explore account-specific details through this table, including account ID, name, product, and other related details. Utilize it to uncover information about accounts, such as the products associated with each account, the status of those products, and the overall details of each account.

## Examples

### List all accounts
Explore which New Relic accounts are available and determine the types of reporting events associated with each. This can help with monitoring and understanding the types of alerts and events that are being generated.

```sql+postgres
select
  id,
  name,
  reporting_event_types
from
  newrelic_account;
```

```sql+sqlite
select
  id,
  name,
  reporting_event_types
from
  newrelic_account;
```