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
import { Observable } from 'rxjs';
import { trackRequest } from 'tracking';

export class DataSource extends DataSourceWithBackend<CSVQuery, CSVDataSourceOptions> {
  jsonData: CSVDataSourceOptions;

  constructor(instanceSettings: DataSourceInstanceSettings<CSVDataSourceOptions>) {
    super(instanceSettings);

    this.jsonData = instanceSettings.jsonData;
  }

  query(request: DataQueryRequest<CSVQuery>): Observable<DataQueryResponse> {
    trackRequest(request);
    return super.query(request);
  }

  applyTemplateVariables(query: CSVQuery, scopedVars: ScopedVars): Record<string, any> {
    const apply = (text: string) => getTemplateSrv().replace(text, scopedVars);

    return {
      ...query,
      schema: query.schema?.map(({ name, type }) => ({
        name: apply(name),
        type,
      })),

      // HTTP settings
      path: apply(query.path),
      queryParams: apply(query.queryParams),
      params: query.params?.map((param) => param.map(apply)),
      headers: query.headers?.map((header) => header.map(apply)),
      body: apply(query.body),
    };
  }

  /**
   * This line adds support for annotation queries in >=7.2.
   */
  annotations = {};

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

    let res: DataQueryResponse | undefined;

    try {
      res = await this.query(request).toPromise();
    } catch (err) {
      return Promise.reject(err);
    }

    if (res && (!res.data.length || !res.data[0].fields.length)) {
      return [];
    }

    return res ? (res.data[0] as DataFrame).fields[0].values.toArray().map((_) => ({ text: _.toString() })) : [];
  }
}
