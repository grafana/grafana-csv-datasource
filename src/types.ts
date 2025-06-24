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
  timezone: string;
  decimalSeparator: string;

  method: string;
  path: string;
  queryParams: string;
  params: Array<Pair<string, string>>;
  headers: Array<Pair<string, string>>;
  body: string;

  experimental: {
    regex: boolean;
  };
}

export const defaultQuery: Pick<
  CSVQuery,
  'delimiter' | 'decimalSeparator' | 'header' | 'ignoreUnknown' | 'skipRows' | 'schema'
> = {
  delimiter: ',',
  decimalSeparator: '.',
  header: true,
  ignoreUnknown: false,
  skipRows: 0,
  schema: [],
};

export interface CSVDataSourceOptions extends DataSourceJsonData {
  storage?: string;
  queryParams?: string;
}

export const defaultOptions: Partial<CSVDataSourceOptions> = {
  storage: 'http',
};

export type Pair<T, K> = [T, K];
