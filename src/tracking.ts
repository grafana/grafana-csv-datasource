import { DataQueryRequest } from '@grafana/data';
import { reportInteraction } from '@grafana/runtime';
import { CSVQuery } from 'types';

export const trackRequest = (request: DataQueryRequest<CSVQuery>) => {
  request.targets.forEach((target) => {
    reportInteraction('grafana_csv_query_executed', {
      app: request.app,
      delimiter: target.delimiter,
      decimalSeparator: target.decimalSeparator,
      skipRows: target.skipRows,
      header: target.header,
      ignoreUnknown: target.ignoreUnknown,
      timezone: target.timezone ?? 'None',
      method: target.method ?? 'GET',
    });
  });
};
