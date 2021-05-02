import React from 'react';

import { CSVQuery } from './types';
import { QueryEditor } from './QueryEditor';
import { DataSource } from 'datasource';

interface Props {
  query: CSVQuery;
  onChange: (query: CSVQuery, definition?: string) => void;
  datasource: DataSource;
}

export const VariableQueryEditor = ({ onChange, query, datasource }: Props) => {
  return (
    <QueryEditor onRunQuery={() => {}} onChange={(query) => onChange(query)} query={query} datasource={datasource} />
  );
};
