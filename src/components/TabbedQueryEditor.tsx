import { css } from '@emotion/css';
import { TimeRange } from '@grafana/data';
import { CodeEditor, InfoBox, InlineField, InlineFieldRow, Input, RadioButtonGroup, useTheme } from '@grafana/ui';
import { DataSource } from 'datasource';
import React, { useState } from 'react';
import AutoSizer from 'react-virtualized-auto-sizer';
import { CSVQuery, Pair } from '../types';
import { KeyValueEditor } from './KeyValueEditor';
import { PathEditor } from './PathEditor';
import { getQueryWithDefaults } from 'utils';

// Display a warning message when user adds any of the following headers.
const sensitiveHeaders = ['authorization', 'proxy-authorization', 'x-api-key'];

interface Props {
  onChange: (query: CSVQuery) => void;
  onRunQuery: () => void;
  query: CSVQuery;
  limitFields?: number;
  datasource: DataSource;
  range?: TimeRange;

  fieldsTab: React.ReactNode;
  experimentalTab: React.ReactNode;
}

export const TabbedQueryEditor = ({ query, onChange, onRunQuery, fieldsTab, experimentalTab, datasource }: Props) => {
  const [bodyType, setBodyType] = useState('plaintext');
  const [tabIndex, setTabIndex] = useState(0);
  const theme = useTheme();

  const q = getQueryWithDefaults(query);

  const onBodyChange = (body: string) => {
    onChange({ ...q, body });
    onRunQuery();
  };

  const onParamsChange = (params: Array<Pair<string, string>>) => {
    onChange({ ...q, params });
    onRunQuery();
  };

  const onHeadersChange = (headers: Array<Pair<string, string>>) => {
    onChange({ ...q, headers });
    onRunQuery();
  };

  const tabs = [
    {
      title: 'Fields',
      disabled: false,
      content: fieldsTab,
    },
    {
      title: 'Path',
      disabled: false,
      content:
        datasource.jsonData.storage === 'http' ? (
          <PathEditor
            method={q.method ?? 'GET'}
            onMethodChange={(method) => {
              onChange({ ...q, method });
              onRunQuery();
            }}
            path={q.path ?? ''}
            onPathChange={(path) => {
              onChange({ ...q, path: path });
              onRunQuery();
            }}
          />
        ) : (
          <InlineField
            label="Relative path"
            tooltip={'The path here is relative to the URL defined in the data source configuration.'}
          >
            <Input
              placeholder={'file.csv'}
              value={q.path ?? ''}
              onChange={(e) => {
                onChange({ ...q, path: e.currentTarget.value });
                onRunQuery();
              }}
            />
          </InlineField>
        ),
    },
    {
      title: 'Params',
      disabled: datasource.jsonData.storage !== 'http',
      content: (
        <KeyValueEditor
          addRowLabel={'Add param'}
          columns={['Key', 'Value']}
          values={q.params ?? []}
          onChange={onParamsChange}
          onBlur={() => onRunQuery()}
        />
      ),
    },
    {
      title: 'Headers',
      disabled: datasource.jsonData.storage !== 'http',
      content: (
        <KeyValueEditor
          addRowLabel={'Add header'}
          columns={['Key', 'Value']}
          values={q.headers ?? []}
          onChange={onHeadersChange}
          onBlur={() => onRunQuery()}
        />
      ),
    },
    {
      title: 'Body',
      disabled: datasource.jsonData.storage !== 'http',
      content: (
        <>
          <InlineFieldRow>
            <InlineField label="Syntax highlighting">
              <RadioButtonGroup
                value={bodyType}
                onChange={(v) => setBodyType(v ?? 'plaintext')}
                options={[
                  { label: 'Text', value: 'plaintext' },
                  { label: 'JSON', value: 'json' },
                  { label: 'XML', value: 'xml' },
                ]}
              />
            </InlineField>
          </InlineFieldRow>
          <InlineFieldRow>
            <AutoSizer
              disableHeight
              className={css`
                margin-bottom: ${theme.spacing.sm};
              `}
            >
              {(size) => (
                <CodeEditor
                  value={q.body || ''}
                  language={bodyType}
                  width={size.width}
                  height="200px"
                  showMiniMap={false}
                  showLineNumbers={true}
                  onBlur={onBodyChange}
                />
              )}
            </AutoSizer>
          </InlineFieldRow>
        </>
      ),
    },
    {
      title: 'Experimental',
      disabled: false,
      content: experimentalTab,
    },
  ].filter((tab) => !tab.disabled);

  return (
    <>
      {tabs.length > 1 && (
        <InlineFieldRow>
          <InlineField>
            <RadioButtonGroup
              onChange={(e) => setTabIndex(e ?? 0)}
              value={tabIndex}
              options={tabs.map((tab, idx) => ({ label: tab.title, value: idx }))}
            />
          </InlineField>
        </InlineFieldRow>
      )}

      {q.method === 'GET' && q.body && (
        <InfoBox severity="warning" style={{ maxWidth: '700px', whiteSpace: 'normal' }}>
          {"GET requests can't have a body. The body you've defined will be ignored."}
        </InfoBox>
      )}

      {(q.headers ?? []).map(([key, _]) => key.toLowerCase()).find((_) => sensitiveHeaders.includes(_)) && (
        <InfoBox severity="warning" style={{ maxWidth: '700px', whiteSpace: 'normal' }}>
          {
            "It looks like you're adding credentials in the header. Since queries are stored unencrypted, it's strongly recommended that you add any secrets to the data source config instead."
          }
        </InfoBox>
      )}

      {tabs[tabIndex].content}
    </>
  );
};

export const formatCacheTimeLabel = (s: number) => {
  if (s < 60) {
    return s + 's';
  } else if (s < 3600) {
    return s / 60 + 'm';
  }

  return s / 3600 + 'h';
};
