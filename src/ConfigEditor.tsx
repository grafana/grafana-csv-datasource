import { DataSourcePluginOptionsEditorProps, GrafanaTheme2 } from '@grafana/data';
import { Field, Input, RadioButtonGroup, useStyles2 } from '@grafana/ui';
import defaults from 'lodash/defaults';
import React, { ChangeEvent } from 'react';
import { CSVDataSourceOptions, defaultOptions } from './types';
import {
  AdvancedHttpSettings,
  Auth,
  ConfigSection,
  ConnectionSettings,
  DataSourceDescription,
  convertLegacyAuthProps,
} from '@grafana/experimental';
import { css } from '@emotion/css';
import { Divider } from 'components/Divider';

interface Props extends DataSourcePluginOptionsEditorProps<CSVDataSourceOptions> {}

/**
 * ConfigEditor lets the user configure connection details like the URL or
 * authentication.
 */
export const ConfigEditor: React.FC<Props> = ({ options, onOptionsChange }) => {
  const jsonData = defaults(options.jsonData, defaultOptions);
  const styles = useStyles2(getStyles);

  const onParamsChange = (e: ChangeEvent<HTMLInputElement>) => {
    onOptionsChange({
      ...options,
      jsonData: {
        ...jsonData,
        queryParams: e.currentTarget.value,
      },
    });
  };

  const onStorageChange = (value?: string) => {
    onOptionsChange({
      ...options,
      jsonData: {
        ...jsonData,
        storage: value!,
      },
    });
  };

  const onLocalPathChange = (e: ChangeEvent<HTMLInputElement>) => {
    onOptionsChange({
      ...options,
      url: e.currentTarget.value,
    });
  };

  return (
    <>
      <DataSourceDescription
        dataSourceName="CSV"
        docsLink="https://grafana.com/docs/plugins/marcusolsson-csv-datasource/latest/"
        hasRequiredFields={false}
      />

      <Divider />

      <Field label="Storage Location">
        <RadioButtonGroup
          options={[
            { label: 'HTTP', value: 'http' },
            { label: 'Local', value: 'local' },
          ]}
          value={jsonData.storage}
          onChange={onStorageChange}
        />
      </Field>

      <Divider />

      {jsonData.storage === 'http' ? (
        <>
          <ConnectionSettings config={options} onChange={onOptionsChange} urlPlaceholder="http://localhost:8080" />

          <Divider />

          <Auth
            {...convertLegacyAuthProps({
              config: options,
              onChange: onOptionsChange,
            })}
          />

          <Divider />

          <ConfigSection title="Additional settings" isCollapsible>
            <AdvancedHttpSettings config={options} onChange={onOptionsChange} />

            <div className={styles.space} />

            <Field label="Custom query parameters" description="Add custom parameters to your queries.">
              <Input
                width={40}
                value={jsonData.queryParams}
                onChange={onParamsChange}
                spellCheck={false}
                placeholder="limit=100"
              />
            </Field>
          </ConfigSection>
        </>
      ) : null}

      {jsonData.storage === 'local' ? (
        <Field label="Path">
          <Input
            value={options.url}
            onChange={onLocalPathChange}
            spellCheck={false}
            width={40}
            placeholder="Path to the CSV file"
          />
        </Field>
      ) : null}
    </>
  );
};

const getStyles = (theme: GrafanaTheme2) => {
  return {
    space: css({
      width: '100%',
      height: theme.spacing(2),
    }),
  };
};
