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
  const saveQuery = (newQuery: CSVQuery) => {
    if (newQuery) {
      onChange(newQuery, newQuery.schema[0].name);
    }
  };

  return (
    <QueryEditor
      onRunQuery={() => {}}
      onChange={saveQuery}
      query={{ ...query, ignoreUnknown: true }}
      datasource={datasource}
      limitFields={1}
      editorContext="variables"
    />
  );
};
