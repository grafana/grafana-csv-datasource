import { SelectableValue } from '@grafana/data';
import { InlineField, QueryField, Select } from '@grafana/ui';
import React from 'react';
import { FieldSchema } from 'types';

interface Props {
  field: FieldSchema;
  onFieldChange: (field: FieldSchema) => void;
}

export const CSVQueryField = ({ field, onFieldChange }: Props) => {
  const onNameChange = (value: string) => {
    onFieldChange({ ...field, name: value });
  };
  const onTypeChange = (selectableValue: SelectableValue<string>) => {
    onFieldChange({ ...field, type: selectableValue.value! });
  };

  return (
    <>
      <InlineField label="Field" tooltip={`Name of the CSV column to include.`} grow>
        <QueryField query={field.name} onChange={onNameChange} portalOrigin="csv" />
      </InlineField>
      <InlineField label="Type" tooltip="Set the type of a field. By default, all fields have type String.">
        <Select
          width={12}
          value={field.type}
          onChange={onTypeChange}
          options={[
            { label: 'String', value: 'string' },
            { label: 'Number', value: 'number' },
            { label: 'Time', value: 'time' },
            { label: 'Boolean', value: 'boolean' },
          ]}
        />
      </InlineField>
    </>
  );
};
