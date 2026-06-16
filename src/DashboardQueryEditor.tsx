import { QueryEditorProps } from '@grafana/data';
import React from 'react';
import { DataSource } from './datasource';
import { QueryEditor } from './QueryEditor';
import { CSVDataSourceOptions, CSVQuery } from './types';

type Props = QueryEditorProps<DataSource, CSVQuery, CSVDataSourceOptions>;

export const DashboardQueryEditor = ({ onRunQuery, onChange, query, datasource }: Props) => {
  return <QueryEditor onRunQuery={onRunQuery} onChange={onChange} query={query} datasource={datasource} />;
};
