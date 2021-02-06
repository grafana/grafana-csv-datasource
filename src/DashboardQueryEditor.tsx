import React from 'react';

import { QueryEditorProps } from '@grafana/data';
import { DataSource } from './datasource';
import { CSVDataSourceOptions, CSVQuery } from './types';
import { QueryEditor } from './QueryEditor';

type Props = QueryEditorProps<DataSource, CSVQuery, CSVDataSourceOptions>;

export const DashboardQueryEditor = ({ onRunQuery, onChange, query }: Props) => {
  return <QueryEditor onRunQuery={onRunQuery} onChange={onChange} query={query} />;
};
