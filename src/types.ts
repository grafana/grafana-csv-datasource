import { DataQuery, DataSourceJsonData } from '@grafana/data';

export type FieldSchema = {
  name: string;
  type: string;
};

export interface CSVQuery extends DataQuery {
  delimiter: string;
  schema: FieldSchema[];
  header: boolean;
  ignoreUnknown: boolean;
  skipRows: number;
}

export const defaultQuery: Partial<CSVQuery> = {
  delimiter: ',',
  header: true,
  ignoreUnknown: false,
  skipRows: 0,
  schema: [],
};

export interface CSVDataSourceOptions extends DataSourceJsonData {
  queryParams?: string;
}
