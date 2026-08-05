import { test, expect } from '@grafana/plugin-e2e';

// Disable the new dashboard layouts so plugin-e2e's addPanel() helper uses the
// stable "Add panel" flow instead of the sidebar/edit-pane path.
//
// Grafana 13.2.0 (the nightly build in CI's e2e matrix) renamed the sidebar
// test ids, e.g. "edit pane configure panel button" -> "sidebar configure panel
// button". @grafana/plugin-e2e@3.10.0 pins @grafana/e2e-selectors@13.1.0, which
// only knows the old name, so addPanel() waits for a test id that no longer
// exists and panelEditPage setup times out. No stable e2e-selectors release
// carries the new name yet (only Grafana's nightly prerelease does).
//
// Remove this override once we upgrade to a @grafana/plugin-e2e release whose
// bundled @grafana/e2e-selectors includes the Grafana 13.2.0 sidebar selectors.
test.use({ featureToggles: { dashboardNewLayouts: false } });

test('Smoke test: plugin loads config page', async ({ createDataSourceConfigPage, page }) => {
  await createDataSourceConfigPage({ type: 'marcusolsson-csv-datasource' });

  await expect(await page.getByText('Storage Location', { exact: true })).toBeVisible();
});

test('Smoke test: plugin query editor works', async ({ createDataSource, page, panelEditPage }) => {
  const datasource = await createDataSource({ type: 'marcusolsson-csv-datasource' });

  await panelEditPage.datasource.set(datasource.name);

  await expect(await page.getByText('Fields')).toBeVisible();
});
