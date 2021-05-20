import React from 'react';
import { CSVQuery } from './types';
import { TabbedQueryEditor } from 'components/TabbedQueryEditor';
import { FieldEditor } from 'components/FieldEditor';
import { DataSource } from 'datasource';
import { InfoBox, InlineField, InlineFieldRow, Switch } from '@grafana/ui';

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
