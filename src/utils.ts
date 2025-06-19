import { DataSourceSettings } from '@grafana/data';
import { CSVDataSourceOptions, CSVQuery, defaultOptions, defaultQuery } from 'types';

export const getOptionsWithDefaults = (options: DataSourceSettings<CSVDataSourceOptions, {}>) => {
  if (options.jsonData.storage) {
    return options;
  }

  return { ...options, jsonData: { ...options.jsonData, ...defaultOptions } };
};

export const getQueryWithDefaults = (query: CSVQuery): CSVQuery => ({
  ...query,
  delimiter: query.delimiter ?? defaultQuery.delimiter,
  decimalSeparator: query.decimalSeparator ?? defaultQuery.decimalSeparator,
  header: query.header ?? defaultQuery.header,
  ignoreUnknown: query.ignoreUnknown ?? defaultQuery.ignoreUnknown,
  skipRows: query.skipRows ?? defaultQuery.skipRows,
  schema: query.schema ?? defaultQuery.schema,
});
