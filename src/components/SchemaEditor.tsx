import { Button, InlineFieldRow } from '@grafana/ui';
import React from 'react';
import { FieldSchema } from '../types';
import { CSVQueryField } from './CSVQueryField';

interface Props {
  value: FieldSchema[];
  onChange: (value: FieldSchema[]) => void;
  limit?: number;
}

export const SchemaEditor = ({ value, onChange, limit }: Props) => {
  const fields = value.length > 0 ? value : [{ name: '', type: 'string' }];
  const canAddMore = !limit || fields.length < limit;

  const onFieldChange = (idx: number) => (fieldSchema: FieldSchema) => {
    const res = fields.map((field, i) => (i === idx ? fieldSchema : field));
    onChange(res);
  };

  const onAddField = (idx: number) => {
    if (canAddMore) {
      const res = [...fields.slice(0, idx + 1), { name: '', type: 'string' }, ...fields.slice(idx + 1)];
      onChange(res);
    }
  };

  const onRemoveField = (idx: number) => {
    const res = fields.filter((_, i) => i !== idx);
    onChange(res);
  };

  return (
    <>
      {fields.map((_, i) => (
        <InlineFieldRow key={i}>
          <CSVQueryField field={_} onFieldChange={onFieldChange(i)} />
          {canAddMore && (
            <Button variant="secondary" title="plus" onClick={() => onAddField(i)} icon="plus" aria-label="Add Field" />
          )}
          {fields.length > 1 && (
            <Button
              variant="secondary"
              title="minus"
              onClick={() => onRemoveField(i)}
              icon="minus"
              aria-label="Remove Field"
            />
          )}
        </InlineFieldRow>
      ))}
    </>
  );
};
