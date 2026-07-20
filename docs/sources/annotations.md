---
title: Annotations
menuTitle: Annotations
description: This document explains how to setup annotations with CSV datasource
aliases:
keywords:
  - data source
  - csv
labels:
  products:
    - oss
    - grafana cloud
weight: 600
---

{{< admonition type="warning" >}}
This plugin is in maintenance mode. It won't receive new features, and bug fixes aren't guaranteed. We recommend using the [Infinity data source](https://grafana.com/grafana/plugins/yesoreyeram-infinity-datasource/) instead. To get started quickly with CSV data in Grafana, refer to [How to visualize CSV data with Grafana](https://grafana.com/blog/2025/02/05/how-to-visualize-csv-data-with-grafana/).
{{< /admonition >}}

[Annotations](https://grafana.com/docs/grafana/latest/dashboards/annotations) let you extract data from a data source and use it to annotate a dashboard.

To use the CSV data source for annotations, follow the instructions on [Querying other data sources](https://grafana.com/docs/grafana/latest/dashboards/annotations/#querying-other-data-sources). Make sure to select the CSV from the list of data sources.

Configure a query with _at least_ two fields:

- A **String** field for the annotation text
- A **Time** field for the annotation time

If you want to add titles or tags to the annotations, you can add additional **Fields** with the appropriate types.

For more information on how to configure a query, refer to [Query editor](query-editor.md).
