import defaults from 'lodash/defaults';
import React, { FormEvent, useState } from 'react';
import { SelectableValue } from '@grafana/data';
import { InlineFieldRow, InlineField, Select, Switch, Input } from '@grafana/ui';
import { SchemaEditor } from './SchemaEditor';
import { CSVQuery, defaultQuery, FieldSchema } from '../types';

interface Props {
  query: CSVQuery;
  onChange: (query: CSVQuery) => void;
  onRunQuery: () => void;
  limit?: number;
  editorContext: string;
}

export const FieldEditor = ({ query, onChange, onRunQuery, limit, editorContext }: Props) => {
  const { header, skipRows, delimiter, decimalSeparator, ignoreUnknown, schema } = defaults(query, defaultQuery);

  const [numSkipRows, setNumSkipRows] = useState(skipRows?.toString());

  const delimOptions = [
    { label: 'Comma', value: ',' },
    { label: 'Semicolon', value: ';' },
    { label: 'Tab', value: '\t' },
  ];

  const onDelimiterChange = (value: SelectableValue<string>) => {
    onChange({ ...query, delimiter: value.value ?? ',' });
    onRunQuery();
  };

  const decimalSeparatorOptions = [
    { label: 'Point', value: '.' },
    { label: 'Comma', value: ',' },
  ];

  const onDecimalSeparatorChange = (value: SelectableValue<string>) => {
    onChange({ ...query, decimalSeparator: value.value ?? '.' });
    onRunQuery();
  };

  const onIgnoreUnknownChange = (e: FormEvent<HTMLInputElement>) => {
    onChange({ ...query, ignoreUnknown: e.currentTarget.checked });
    onRunQuery();
  };

  const onHeaderChange = (e: FormEvent<HTMLInputElement>) => {
    onChange({ ...query, header: e.currentTarget.checked });
    onRunQuery();
  };

  const onSkipRowsChange = (e: FormEvent<HTMLInputElement>) => {
    setNumSkipRows(e.currentTarget.value);
  };

  const onSchemaChange = (fields: FieldSchema[]) => {
    onChange({ ...query, schema: fields });
    onRunQuery();
  };

  return (
    <>
      <InlineFieldRow>
        <InlineField label="Delimiter" tooltip="Character used to separate columns">
          <Select
            width={13}
            value={delimOptions.find((_) => _.value === delimiter)}
            onChange={onDelimiterChange}
            options={delimOptions}
          />
        </InlineField>
        <InlineField
          label="Decimal separator"
          tooltip="Character used to separate the integral part from the fractional part of numbers."
        >
          <Select
            width={13}
            value={decimalSeparatorOptions.find((_) => _.value === decimalSeparator)}
            onChange={onDecimalSeparatorChange}
            options={decimalSeparatorOptions}
          />
        </InlineField>
        <InlineField label="Skip leading rows" tooltip="Number of rows to skip before looking for header">
          <Input
            width={5}
            value={numSkipRows}
            onChange={onSkipRowsChange}
            onBlur={() => {
              onChange({ ...query, skipRows: parseInt(numSkipRows, 10) });
              onRunQuery();
            }}
          />
        </InlineField>
        <InlineField label="Header" tooltip="Data contains a header row with field names">
          <Switch value={header} onChange={onHeaderChange} />
        </InlineField>
        <InlineField
          disabled={editorContext === 'variables'}
          label="Ignore unknown"
          tooltip="Ignore fields that aren't defined in the schema"
        >
          <Switch value={ignoreUnknown} onChange={onIgnoreUnknownChange} />
        </InlineField>
      </InlineFieldRow>
      <SchemaEditor value={schema} onChange={onSchemaChange} limit={limit} />
    </>
  );
};
