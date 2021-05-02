import React from 'react';
import { CSVQuery } from './types';
import { TabbedQueryEditor } from 'components/TabbedQueryEditor';
import { FieldEditor } from 'components/FieldEditor';
import { DataSource } from 'datasource';

interface Props {
  query: CSVQuery;
  onChange: (query: CSVQuery) => void;
  onRunQuery: () => void;
  datasource: DataSource;
}

export const QueryEditor = (props: Props) => {
  const { query, onChange, onRunQuery } = props;

  return (
    <TabbedQueryEditor
      {...props}
      fieldsTab={<FieldEditor query={query} onChange={onChange} onRunQuery={onRunQuery} />}
      experimentalTab={<div></div>}
    />
  );
};
