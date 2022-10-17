module.exports = {
  title: 'CSV for Grafana',
  url: 'https://grafana.github.io',
  baseUrl: '/grafana-csv-datasource/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.svg',
  organizationName: 'grafana', // Usually your GitHub org/user name.
  projectName: 'grafana-csv-datasource', // Usually your repo name.
  scripts: [],
  themeConfig: {
    navbar: {
      title: 'CSV Data Source for Grafana',
      logo: {
        alt: 'Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          href: 'https://github.com/grafana/grafana-csv-datasource',
          label: 'GitHub',
          position: 'right',
        },
        {
          href: 'https://grafana.com/plugins/marcusolsson-csv-datasource',
          label: 'Marketplace',
          position: 'right',
        },
      ],
    },
    footer: {
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Installation',
              to: '/',
            },
            {
              label: 'Configuration',
              to: 'configuration/',
            },
            {
              label: 'Query editor',
              to: 'query-editor/',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Discussions',
              href: 'https://github.com/grafana/grafana-csv-datasource/discussions',
            },
            {
              label: 'Support',
              href: 'https://github.com/grafana/grafana-csv-datasource/discussions/categories/q-a',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Grafana Labs`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl: 'https://github.com/grafana/grafana-csv-datasource/edit/main/website/',
          routeBasePath: '/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
