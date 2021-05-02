import React, { useState, useEffect } from 'react';
import { InlineField, Select, QueryField } from '@grafana/ui';
import { FieldSchema } from 'types';
import { SelectableValue } from '@grafana/data';

interface Props {
  field: FieldSchema;
  onFieldChange: (field: FieldSchema) => void;
}

export const CSVQueryField = ({ field, onFieldChange }: Props) => {
  const [name, setName] = useState(field.name);

  useEffect(() => {
    setName(field.name);
  }, [field]);

  const onNameChange = (value: string) => setName(value);
  const onTypeChange = (selectableValue: SelectableValue<string>) => {
    onFieldChange({ ...field, type: selectableValue.value! });
  };

  return (
    <>
      <InlineField label="Field" tooltip={`Name of the CSV column to include.`} grow>
        <QueryField
          query={name}
          onRunQuery={() => onFieldChange({ ...field, name })}
          onChange={onNameChange}
          portalOrigin="csv"
        />
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
