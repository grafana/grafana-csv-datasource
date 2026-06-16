import { DataSourcePlugin } from '@grafana/data';
import { ConfigEditor } from './ConfigEditor';
import { DataSource } from './datasource';
import { QueryEditor } from './QueryEditor';
import { CSVDataSourceOptions, CSVQuery } from './types';

export const plugin = new DataSourcePlugin<DataSource, CSVQuery, CSVDataSourceOptions>(DataSource)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(QueryEditor);
