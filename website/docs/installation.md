---
id: installation
title: Installation
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

You can install the plugin using [grafana-cli](https://grafana.com/docs/grafana/latest/administration/cli/), or by downloading the plugin manually.

## Install using grafana-cli

To install the latest version of the plugin, run the following command on the Grafana server:

<Tabs
  groupId="operating-systems"
  defaultValue="linux"
  values={[
    {label: 'Linux', value: 'linux'},
    {label: 'macOS', value: 'macos'},
    {label: 'Windows', value: 'windows'},
  ]}>
  <TabItem value="linux">

```bash
grafana-cli plugins install marcusolsson-csv-datasource
```

  </TabItem>
  <TabItem value="macos">

```bash
grafana-cli plugins install marcusolsson-csv-datasource
```

  </TabItem>
  <TabItem value="windows">

```bash
grafana-cli.exe plugins install marcusolsson-csv-datasource
```

  </TabItem>
</Tabs>

## Install manually

1. Go to [Releases](https://github.com/marcusolsson/grafana-csv-datasource/releases) on the GitHub project page
1. Find the release you want to install
1. Download the release by clicking the release asset called `marcusolsson-csv-datasource-<version>.zip`. You may need to uncollapse the **Assets** section to see it.
1. Install the plugin into the Grafana plugins directory

   <Tabs
     groupId="operating-systems"
     defaultValue="linux"
     values={[
       {label: 'Linux', value: 'linux'},
       {label: 'macOS', value: 'macos'},
       {label: 'Windows', value: 'windows'},
     ]}>
     <TabItem value="linux">

     ```bash
     grafana-cli --pluginUrl ./marcusolsson-csv-datasource-<version>.zip plugins install marcusolsson-csv-datasource
     ```

     </TabItem>
     <TabItem value="macos">

     ```bash
     grafana-cli --pluginUrl ./marcusolsson-csv-datasource-<version>.zip plugins install marcusolsson-csv-datasource
     ```

     </TabItem>
     <TabItem value="windows">

     ```bash
     grafana-cli.exe --pluginUrl marcusolsson-csv-datasource-<version>.zip plugins install marcusolsson-csv-datasource
     ```

     </TabItem>
   </Tabs>

1. Restart the Grafana server to load the plugin
