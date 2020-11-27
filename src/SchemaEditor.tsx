import React, { FormEvent, useState, useEffect } from 'react';
import { InlineFieldRow, InlineField, Icon, Input, Select, Button } from '@grafana/ui';

import { FieldSchema } from './types';
import {} from '@emotion/core';
import { SelectableValue } from '@grafana/data';

interface Props {
  value: FieldSchema[];
  onChange: (value: FieldSchema[]) => void;
}

export const SchemaEditor = ({ value, onChange }: Props) => {
  const [internalValue, setInternalValue] = useState(value);

  useEffect(() => {
    setInternalValue(value);
  }, [value]);

  const onNameChange = (idx: number) => (e: FormEvent<HTMLInputElement>) => {
    setInternalValue(internalValue.map((field, i) => (i === idx ? { ...field, name: e.currentTarget.value } : field)));
  };
  const onTypeChange = (idx: number) => (selectableValue: SelectableValue<string>) => {
    const res = internalValue.map((field, i) => (i === idx ? { ...field, type: selectableValue.value! } : field));
    setInternalValue(res);
    onChange(res);
  };
  const onAppendField = () => {
    const res = [...internalValue, { name: '', type: 'string' }];
    setInternalValue(res);
    onChange(res);
  };
  const onAddField = (idx: number) => {
    const res = [...internalValue.slice(0, idx + 1), { name: '', type: 'string' }, ...internalValue.slice(idx + 1)];
    setInternalValue(res);
    onChange(res);
  };
  const onRemoveField = (idx: number) => {
    const res = internalValue.filter((_, i) => i !== idx);
    setInternalValue(res);
    onChange(res);
  };

  return (
    <>
      {internalValue.map((_, i) => (
        <InlineFieldRow>
          <InlineField label="Name">
            <Input width={25} value={_.name} onChange={onNameChange(i)} onBlur={() => onChange(internalValue)} />
          </InlineField>
          <InlineField label="Type">
            <Select
              width={12}
              value={_.type}
              onChange={onTypeChange(i)}
              options={[
                { label: 'String', value: 'string' },
                { label: 'Number', value: 'number' },
                { label: 'Boolean', value: 'boolean' },
              ]}
            />
          </InlineField>
          <a className="gf-form-label">
            <Icon name="plus" size="lg" onClick={() => onAddField(i)} />
          </a>
          <a className="gf-form-label">
            <Icon name="minus" size="lg" onClick={() => onRemoveField(i)} />
          </a>
        </InlineFieldRow>
      ))}
      {internalValue.length === 0 ? (
        <InlineFieldRow>
          <Button variant="secondary" icon="plus" onClick={onAppendField}>
            Add field
          </Button>
        </InlineFieldRow>
      ) : null}
    </>
  );
};
