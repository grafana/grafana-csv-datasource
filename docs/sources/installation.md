---
title: Installation
---

# Installation

You can install the plugin by using the [Grafana CLI](https://grafana.com/docs/grafana/latest/cli/) or downloading the plugin manually.
After installing the plugin, you can [configure it](configuration.md) and use it to [query data](query-data.md).

## Install using the Grafana CLI

To install the latest version of the plugin, run the following command on the Grafana server:

```
grafana-cli plugins install grafana-csv-datasource
```

## Install manually

1. Go to [Releases](https://github.com/grafana/grafana-csv-datasource/releases) on the GitHub project page.
1. Find the release you want to install.
1. Download the release by clicking the release asset called `marcusolsson-csv-datasource-<version>.zip`.
   You might need to uncollapse the **Assets** section to see it.
1. Install the plugin into the Grafana plugins directory.

   ```
   grafana-cli --pluginUrl ./grafana-csv-datasource-<version>.zip plugins install grafana-csv-datasource
   ```
1. Restart the Grafana server to load the plugin.
