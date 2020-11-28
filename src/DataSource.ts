import { DataSourceInstanceSettings, ScopedVars } from '@grafana/data';
import { DataSourceWithBackend, getTemplateSrv } from '@grafana/runtime';
import { CSVDataSourceOptions, CSVQuery } from './types';

export class DataSource extends DataSourceWithBackend<CSVQuery, CSVDataSourceOptions> {
  constructor(instanceSettings: DataSourceInstanceSettings<CSVDataSourceOptions>) {
    super(instanceSettings);
  }

  applyTemplateVariables(query: CSVQuery, scopedVars: ScopedVars): Record<string, any> {
    return {
      ...query,
      schema: query.schema.map(({ name, type }) => ({
        name: getTemplateSrv().replace(name, scopedVars),
        type,
      })),
    };
  }
}
