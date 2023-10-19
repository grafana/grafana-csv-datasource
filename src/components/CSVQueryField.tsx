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
        <QueryField
          query={field.name}
          onChange={onNameChange}
          portalOrigin="csv"
          // https://github.com/grafana/grafana/commit/f6d3a5cc9411acf20c9a8a497c993667aa825062
          // auto-adds the `onblur: empty-function` prop,
          // but for older grafana versions we need to add this.
          // we can remove this when we stop supporting grafana versions before that commit.
          onBlur={() => {}}
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
