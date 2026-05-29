import { CustomVariableSupport, type DataQueryRequest, type DataQueryResponse } from '@grafana/data';
import { uniqueId } from 'lodash';
import type { Observable } from 'rxjs';
import type { CSVQuery } from 'types';
import type { DataSource } from './datasource';
import { VariableQueryEditor } from './VariableQueryEditor';

export class VariableSupport extends CustomVariableSupport<DataSource> {
  constructor() {
    super();
  }
  editor = VariableQueryEditor;
  query(request: DataQueryRequest<CSVQuery>): Observable<DataQueryResponse> {
    // Make sure that every query has a refId
    const queries = request.targets.map((query) => {
      return { ...query, refId: query.refId || uniqueId('tempVar') };
    });
    return this.query({ ...request, targets: queries });
  }
}
