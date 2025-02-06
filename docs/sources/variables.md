---
title: Variables
menuTitle: Variables
description: This document explains how to setup variables with CSV datasource
aliases:
keywords:
  - data source
  - csv
labels:
  products:
    - oss
    - grafana cloud
weight: 500
---

{{< admonition type="warning" >}}
This plugin is now in maintenance mode, no new features will be added. We recommend using the [Infinity data source plugin](https://grafana.com/grafana/plugins/yesoreyeram-infinity-datasource/) instead.  If you want to get started
quickly with CSV and Grafana, please read [How to Visualize CSV Data with Grafana](https://grafana.com/blog/2025/02/05/how-to-visualize-csv-data-with-grafana/), which uses the recommended approach.
{{< /admonition >}}

[Query variables](https://grafana.com/docs/grafana/latest/variables/variable-types/add-query-variable) let you extract data from a data source and use it to populate a dashboard variable.

To query the CSV data source for variables, follow the instructions on how to [Add a query variable](https://grafana.com/docs/grafana/latest/variables/variable-types/add-query-variable). Make sure to select the CSV from the list of data sources.

For more information on how to configure a query, refer to [Query editor](query-editor.md).
