import {
  DataFrame,
  DataQueryRequest,
  DataQueryResponse,
  DataSourceInstanceSettings,
  MetricFindValue,
  ScopedVars,
} from '@grafana/data';
import { DataSourceWithBackend, getTemplateSrv } from '@grafana/runtime';
import { CSVDataSourceOptions, CSVQuery } from './types';

export class DataSource extends DataSourceWithBackend<CSVQuery, CSVDataSourceOptions> {
  constructor(instanceSettings: DataSourceInstanceSettings<CSVDataSourceOptions>) {
    super(instanceSettings);
  }

  applyTemplateVariables(query: CSVQuery, scopedVars: ScopedVars): Record<string, any> {
    return {
      ...query,
      schema: query.schema?.map(({ name, type }) => ({
        name: getTemplateSrv().replace(name, scopedVars),
        type,
      })),
    };
  }

  async metricFindQuery?(query: CSVQuery, options: any): Promise<MetricFindValue[]> {
    const request = {
      targets: [
        {
          ...query,
          refId: 'metricFindQuery',
        },
      ],
      range: options.range,
      rangeRaw: options.rangeRaw,
    } as DataQueryRequest<CSVQuery>;

    let res: DataQueryResponse;

    try {
      res = await this.query(request).toPromise();
    } catch (err) {
      return Promise.reject(err);
    }

    if (!res || !res.data || res.data.length < 0) {
      return [];
    }

    return (res.data[0] as DataFrame).fields[0].values.toArray().map((_) => ({ text: _.toString() }));
  }
}
