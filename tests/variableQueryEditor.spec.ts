import { expect, test } from '@grafana/plugin-e2e';

// skipped: variableEditPage.datasource.set() can't find the datasource picker
// in Grafana 12.2.0 (selector mismatch in @grafana/plugin-e2e). Un-skip when
// plugin-e2e adds support for the new variable edit page.
test.skip('should successfully create a variable', async ({ variableEditPage, readProvisionedDataSource }) => {
  const ds = await readProvisionedDataSource({ fileName: 'datasource.yaml', name: 'Employees' });
  await variableEditPage.setVariableType('Query');
  await variableEditPage.datasource.set(ds.name);
  const queryDataRequest = variableEditPage.waitForQueryDataRequest();
  await variableEditPage.runQuery();
  await queryDataRequest;
  await expect(variableEditPage).toDisplayPreviews(
    ['Almaz Russom', 'Brenda Tilman', 'Mada Rawdha Tahan', 'Yuan Yang'],
    { timeout: 15_000 }
  );
});
