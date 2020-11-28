import defaults from 'lodash/defaults';
import React, { FormEvent, useState } from 'react';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { SchemaEditor } from './SchemaEditor';
import { DataSource } from './DataSource';
import { MyDataSourceOptions, MyQuery } from './types';
import { defaultQuery, FieldSchema } from './types';
import { InlineFieldRow, InlineField, Select, Switch, Input } from '@grafana/ui';

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>;

export const QueryEditor = ({ onRunQuery, onChange, query }: Props) => {
  const { header, skipRows, delimiter, ignoreUnknown, schema } = defaults(query, defaultQuery);

  const [numSkipRows, setNumSkipRows] = useState(skipRows?.toString());

  const delimOptions = [
    { label: 'Comma', value: ',' },
    { label: 'Semicolon', value: ';' },
    { label: 'Tab', value: '\t' },
  ];

  const onDelimiterChange = (value: SelectableValue<string>) => {
    onChange({ ...query, delimiter: value.value! });
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
            width={15}
            value={delimOptions.find(_ => _.value === delimiter)}
            onChange={onDelimiterChange}
            options={delimOptions}
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
        <InlineField label="Ignore unknown" tooltip="Ignore fields that aren't defined in the schema">
          <Switch value={ignoreUnknown} onChange={onIgnoreUnknownChange} />
        </InlineField>
      </InlineFieldRow>
      <SchemaEditor value={schema} onChange={onSchemaChange} />
    </>
  );
};
