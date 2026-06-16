import { DataSourceSettings } from '@grafana/data';

import { getOptionsWithDefaults, getQueryWithDefaults } from './utils';
import { CSVDataSourceOptions, CSVQuery } from './types';

describe('utils', () => {
  describe('getOptionsWithDefaults', () => {
    const BLANK_OPTIONS: DataSourceSettings<CSVDataSourceOptions> = {
      access: '',
      basicAuth: false,
      basicAuthUser: '',
      id: -1,
      isDefault: false,
      jsonData: {},
      name: '',
      orgId: -1,
      readOnly: false,
      secureJsonFields: {},
      type: '',
      typeLogoUrl: '',
      typeName: '',
      uid: '',
      url: '',
      user: '',
      database: '',
      withCredentials: false,
    };

    it('should return options unchanged when storage is already set', () => {
      const options = { ...BLANK_OPTIONS, jsonData: { storage: 'local' } };

      const result = getOptionsWithDefaults(options);

      expect(result).toBe(options);
      expect(result.jsonData.storage).toBe('local');
    });

    it('should add default options when storage is not set', () => {
      const options = { ...BLANK_OPTIONS, jsonData: {} };

      const result = getOptionsWithDefaults(options);

      expect(result.jsonData.storage).toBe('http');
    });

    it('should preserve existing jsonData properties when storage is already set', () => {
      const options = { ...BLANK_OPTIONS, jsonData: { storage: 'local', queryParams: 'test=value' } };

      const result = getOptionsWithDefaults(options);

      expect(result.jsonData.storage).toBe('local');
      expect(result.jsonData.queryParams).toBe('test=value');
    });

    it('should preserve existing jsonData when adding defaults', () => {
      const options = { ...BLANK_OPTIONS, jsonData: { queryParams: 'test=value' } };

      const result = getOptionsWithDefaults(options);

      expect(result.jsonData.storage).toBe('http');
      expect(result.jsonData.queryParams).toBe('test=value');
    });
  });

  describe('getQueryWithDefaults', () => {
    const createBaseQuery = (): CSVQuery => ({
      refId: 'A',
      delimiter: ',',
      schema: [],
      header: true,
      ignoreUnknown: false,
      skipRows: 0,
      timezone: 'UTC',
      decimalSeparator: '.',
      method: 'GET',
      path: '',
      queryParams: '',
      params: [],
      headers: [],
      body: '',
      experimental: {
        regex: false,
      },
    });

    it('should return query with all defaults when no properties are set', () => {
      const query: any = createBaseQuery();
      delete query.delimiter;
      delete query.decimalSeparator;
      delete query.header;
      delete query.ignoreUnknown;
      delete query.skipRows;
      delete query.schema;

      const result = getQueryWithDefaults(query);

      expect(result.delimiter).toBe(',');
      expect(result.decimalSeparator).toBe('.');
      expect(result.header).toBe(true);
      expect(result.ignoreUnknown).toBe(false);
      expect(result.skipRows).toBe(0);
      expect(result.schema).toEqual([]);
    });

    it('should preserve set values and only use defaults for undefined properties', () => {
      const query = createBaseQuery();
      query.delimiter = ';';
      query.header = false;
      query.skipRows = 5;
      query.schema = [{ name: 'test', type: 'string' }];

      const result = getQueryWithDefaults(query);

      expect(result.delimiter).toBe(';');
      expect(result.decimalSeparator).toBe('.');
      expect(result.header).toBe(false);
      expect(result.ignoreUnknown).toBe(false);
      expect(result.skipRows).toBe(5);
      expect(result.schema).toEqual([{ name: 'test', type: 'string' }]);
    });

    it('should preserve falsy values (empty string, 0, false)', () => {
      const query = createBaseQuery();
      query.delimiter = '';
      query.decimalSeparator = '';
      query.header = false;
      query.ignoreUnknown = false;
      query.skipRows = 0;
      query.schema = [];

      const result = getQueryWithDefaults(query);

      expect(result.delimiter).toBe('');
      expect(result.decimalSeparator).toBe('');
      expect(result.header).toBe(false);
      expect(result.ignoreUnknown).toBe(false);
      expect(result.skipRows).toBe(0);
      expect(result.schema).toEqual([]);
    });

    it('should preserve other query properties', () => {
      const query = createBaseQuery();
      query.refId = 'B';
      query.method = 'POST';
      query.path = '/test.csv';
      query.timezone = 'America/New_York';
      query.experimental.regex = true;

      const result = getQueryWithDefaults(query);

      expect(result.refId).toBe('B');
      expect(result.method).toBe('POST');
      expect(result.path).toBe('/test.csv');
      expect(result.timezone).toBe('America/New_York');
      expect(result.experimental.regex).toBe(true);
    });
  });
});
