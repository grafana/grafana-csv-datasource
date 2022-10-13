---
title: Configuration
---

# Configuration

For general information about configuring data sources, refer to [Data source management](https://grafana.com/docs/grafana/latest/administration/data-source-management/).

## Add a CSV data source

To add a CSV data source to Grafana:

1. In the side menu, select the **Configuration** tab (cog icon).
1. Select **Add data source** in the **Data Sources** tab.
1. Enter "CSV" in the search box to find the CSV data source.
1. Select the search result that says "CSV".
1. In the **URL** field, enter a URL that points to CSV content.

## Allow local mode

The plugin disables reading files from the local file system by default.
To allow local mode, add `allow_local_mode` to the plugin's configuration in your [Grafana configuration file](https://grafana.com/docs/grafana/latest/setup-grafana/configure-grafana/):

```ini
[plugin.grafana-csv-datasource]
allow_local_mode = true
```
