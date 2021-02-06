import { DataSourcePlugin } from '@grafana/data';
import { DataSource } from './datasource';
import { ConfigEditor } from './ConfigEditor';
import { DashboardQueryEditor } from './DashboardQueryEditor';
import { VariableQueryEditor } from './VariableQueryEditor';
import { CSVQuery, CSVDataSourceOptions } from './types';

export const plugin = new DataSourcePlugin<DataSource, CSVQuery, CSVDataSourceOptions>(DataSource)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(DashboardQueryEditor)
  .setVariableQueryEditor(VariableQueryEditor);
