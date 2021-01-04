# CSV data source for Grafana

[![Build](https://github.com/marcusolsson/grafana-csv-datasource/workflows/CI/badge.svg)](https://github.com/marcusolsson/grafana-csv-datasource/actions?query=workflow%3A%22CI%22)
[![Release](https://github.com/marcusolsson/grafana-csv-datasource/workflows/Release/badge.svg)](https://github.com/marcusolsson/grafana-csv-datasource/actions?query=workflow%3ARelease)
[![Marketplace](https://img.shields.io/badge/dynamic/json?color=orange&label=marketplace&prefix=v&query=%24.items%5B%3F%28%40.slug%20%3D%3D%20%22marcusolsson-csv-datasource%22%29%5D.version&url=https%3A%2F%2Fgrafana.com%2Fapi%2Fplugins)](https://grafana.com/grafana/plugins/marcusolsson-csv-datasource)
[![Downloads](https://img.shields.io/badge/dynamic/json?color=orange&label=downloads&query=%24.items%5B%3F%28%40.slug%20%3D%3D%20%22marcusolsson-csv-datasource%22%29%5D.downloads&url=https%3A%2F%2Fgrafana.com%2Fapi%2Fplugins)](https://grafana.com/grafana/plugins/marcusolsson-csv-datasource)
[![License](https://img.shields.io/github/license/marcusolsson/grafana-csv-datasource)](LICENSE)

A data source for loading CSV data into [Grafana](https://grafana.com).

> **Important:** This plugin is still in early development. Only use it if you intend to contribute to its development, by providing feedback.

![Screenshot](https://github.com/marcusolsson/grafana-csv-datasource/raw/master/src/img/screenshot.png)

## Public data sets

Here are a few publicly available CSV data sets that you can try out.

### Confirmed cases of COVID-19

**Source:** [CSSE COVID-19 Data Repository](https://github.com/CSSEGISandData/COVID-19)

#### Configuration

- **URL:** https://github.com/CSSEGISandData/COVID-19/blob/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv

### Natural Gas Prices

**Source:** [DataHub](https://datahub.io/core/natural-gas)

#### Configuration

- **URL:** https://datahub.io/core/natural-gas/r/daily.csv
