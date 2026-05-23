import { test, expect } from '@grafana/plugin-e2e';

test('smoke: should render query editor', async ({ panelEditPage, readProvisionedDataSource }) => {
  const ds = await readProvisionedDataSource({ fileName: 'datasource.yaml' });
  await panelEditPage.datasource.set(ds.name);
  // verify the query editor row loaded with the CSV datasource
  const queryRow = panelEditPage.getQueryEditorRow('A');
  await expect(queryRow.getByText(ds.name)).toBeVisible({ timeout: 15_000 });
});

test('data query should return employee data from local CSV', async ({
  gotoPanelEditPage,
  readProvisionedDashboard,
}) => {
  const dashboard = await readProvisionedDashboard({ fileName: 'csv/demo.json' });
  const panelEditPage = await gotoPanelEditPage({ dashboard, id: '2' });
  // the provisioned panel already has a table visualization with data
  await expect(panelEditPage.panel.locator).toContainText('Brenda Tilman', { timeout: 15_000 });
  await expect(panelEditPage.panel.locator).toContainText('Marketing');
});

test('data query should return natural gas price data from HTTP CSV', async ({
  gotoPanelEditPage,
  readProvisionedDashboard,
}) => {
  test.setTimeout(60_000);
  const dashboard = await readProvisionedDashboard({ fileName: 'csv/demo.json' });
  const panelEditPage = await gotoPanelEditPage({ dashboard, id: '1' });
  // the provisioned panel fetches CSV from GitHub raw - verify it loads data
  await expect(panelEditPage.panel.locator).not.toContainText('No data', { timeout: 30_000 });
});
