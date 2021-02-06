import React from 'react';

import { CSVQuery } from './types';
import { QueryEditor } from './QueryEditor';

interface Props {
  query: CSVQuery;
  onChange: (query: CSVQuery, definition?: string) => void;
}

export const VariableQueryEditor = ({ onChange, query }: Props) => {
  return <QueryEditor onRunQuery={() => {}} onChange={query => onChange(query)} query={query} />;
};
