import { test, expect } from '@grafana/plugin-e2e';

test('Smoke test: plugin loads config page', async ({ createDataSourceConfigPage, page }) => {
  await createDataSourceConfigPage({ type: 'marcusolsson-csv-datasource' });

  await expect(await page.getByText('Type: CSV', { exact: true })).toBeVisible();
  await expect(await page.getByText('Storage Location', { exact: true })).toBeVisible();
});

test('Smoke test: plugin query editor works', async ({ createDataSource, page, panelEditPage }) => {
  const datasource = await createDataSource({ type: 'marcusolsson-csv-datasource' });

  await panelEditPage.datasource.set(datasource.name);

  await expect(await page.getByText('Fields')).toBeVisible();
});
