import { DataQuery, DataSourceJsonData } from '@grafana/data';

export type FieldSchema = {
  name: string;
  type: string;
};

export interface MyQuery extends DataQuery {
  delimiter: string;
  schema: FieldSchema[];
  header: boolean;
  ignoreUnknown: boolean;
  skipRows: number;
}

export const defaultQuery: Partial<MyQuery> = {
  delimiter: ',',
  header: true,
  ignoreUnknown: false,
  skipRows: 0,
  schema: [],
};

export interface MyDataSourceOptions extends DataSourceJsonData {
  queryParams?: string;
}
