---
id: configuration
title: Configuration
---

## Add a CSV data source

1. In the side menu, click the **Configuration** tab (cog icon)
1. Click **Add data source** in the top-right corner of the **Data Sources** tab
1. Enter "CSV" in the search box to find the CSV data source
1. Click the search result that says "CSV"
1. In **URL**, enter a URL that points to CSV content

## Allow local mode

This allows you to read files from local computer.

Reading files from the local file system is disabled by default. To allow local mode, add the following to your Grafana config file:

```ini
[plugin.marcusolsson-csv-datasource]
allow_local_mode = true
```

> Note: Local mode option is not available in Grafana Cloud and other hosted grafana environments. In such cases, use a web server such as nginx to serve the CSV files over http and then use http url such as `http://localhost/my-csv-app/my-csv-file.csv`.
