---
title: Installation
menuTitle: Installation
description: This document explains how to install CSV datasource
aliases:
keywords:
  - data source
  - csv
labels:
  products:
    - oss
    - grafana cloud
weight: 200
---

You can install the CSV plugin using [grafana-cli](https://grafana.com/docs/grafana/latest/administration/cli/), or by downloading the plugin manually.

## Install using grafana-cli

To install the latest version of the plugin, run the following command on the Grafana server:

In linux/macos, you will be installing using the following command

```bash
grafana-cli plugins install marcusolsson-csv-datasource
```

whereas in windows machine, use the following comand

```bash
grafana-cli.exe plugins install marcusolsson-csv-datasource
```

## Install manually

1. Go to [Releases](https://github.com/grafana/grafana-csv-datasource/releases) on the GitHub project page
2. Find the release you want to install
3. Download the release by clicking the release asset called `marcusolsson-csv-datasource-<version>.zip`. You may need to uncollapse the **Assets** section to see it.
4. Install the plugin into the Grafana plugins directory. In the linux/macos, use the following command

   ```bash
   grafana-cli --pluginUrl ./marcusolsson-csv-datasource-<version>.zip plugins install marcusolsson-csv-datasource
   ```

   whereas in windows, use the following command

   ```bash
   grafana-cli.exe --pluginUrl marcusolsson-csv-datasource-<version>.zip plugins install marcusolsson-csv-datasource
   ```

5. Restart the Grafana server to load the plugin
