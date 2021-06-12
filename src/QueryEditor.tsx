import { QueryEditorProps } from '@grafana/data';
import { InfoBox, InlineField, InlineFieldRow, Switch } from '@grafana/ui';
import { FieldEditor } from 'components/FieldEditor';
import { TabbedQueryEditor } from 'components/TabbedQueryEditor';
import { DataSource } from 'datasource';
import React from 'react';
import { CSVDataSourceOptions, CSVQuery } from './types';

interface Props extends QueryEditorProps<DataSource, CSVQuery, CSVDataSourceOptions> {
  limitFields?: number;
  editorContext?: string;
}

export const QueryEditor = (props: Props) => {
  const { query, onChange, onRunQuery, limitFields, editorContext } = props;

  return (
    <TabbedQueryEditor
      {...props}
      fieldsTab={
        <FieldEditor
          query={query}
          onChange={onChange}
          onRunQuery={onRunQuery}
          limit={limitFields}
          editorContext={editorContext || 'default'}
        />
      }
      experimentalTab={
        <>
          <InfoBox severity="warning" style={{ maxWidth: '700px', whiteSpace: 'normal' }}>
            <p>
              {`The features listed here are experimental. They might change or be removed without notice. In the tooltip for
          each feature, there's a link to a pull request where you can submit feedback for that feature.`}
            </p>
          </InfoBox>
          <InlineFieldRow>
            <InlineField
              label="Enable regular expressions"
              tooltip={
                <>
                  <p>
                    {
                      'When enabled, field names become regular expressions and can be used to set a type for multiple fields at once.'
                    }
                  </p>
                  <a
                    href="https://github.com/marcusolsson/grafana-csv-datasource/issues/68"
                    target="_blank"
                    rel="noreferrer"
                  >
                    Share feedback
                  </a>
                </>
              }
            >
              <Switch
                value={!!query.experimental?.regex}
                onChange={(e) => {
                  onChange({
                    ...query,
                    experimental: {
                      regex: e.currentTarget.checked,
                    },
                  });
                  onRunQuery();
                }}
              />
            </InlineField>
          </InlineFieldRow>
        </>
      }
    />
  );
};
