import { DataSourcePluginOptionsEditorProps, GrafanaTheme2 } from '@grafana/data';
import { Field, Input, RadioButtonGroup, useStyles2 } from '@grafana/ui';
import React, { ChangeEvent } from 'react';
import { CSVDataSourceOptions } from './types';
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
import { getOptionsWithDefaults } from 'utils';

interface Props extends DataSourcePluginOptionsEditorProps<CSVDataSourceOptions> {}

/**
 * ConfigEditor lets the user configure connection details like the URL or
 * authentication.
 */
export const ConfigEditor: React.FC<Props> = (props) => {
  const { onOptionsChange } = props;
  const optionsWithDefaults = getOptionsWithDefaults(props.options);
  const styles = useStyles2(getStyles);

  const onParamsChange = (e: ChangeEvent<HTMLInputElement>) => {
    onOptionsChange({
      ...optionsWithDefaults,
      jsonData: {
        ...optionsWithDefaults.jsonData,
        queryParams: e.currentTarget.value,
      },
    });
  };

  const onStorageChange = (value?: string) => {
    onOptionsChange({
      ...optionsWithDefaults,
      jsonData: {
        ...optionsWithDefaults.jsonData,
        storage: value!,
      },
    });
  };

  const onLocalPathChange = (e: ChangeEvent<HTMLInputElement>) => {
    onOptionsChange({
      ...optionsWithDefaults,
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
          value={optionsWithDefaults.jsonData.storage}
          onChange={onStorageChange}
        />
      </Field>

      <Divider />

      {optionsWithDefaults.jsonData.storage === 'http' ? (
        <>
          <ConnectionSettings
            config={optionsWithDefaults}
            onChange={onOptionsChange}
            urlPlaceholder="http://localhost:8080"
          />

          <Divider />

          <Auth
            {...convertLegacyAuthProps({
              config: optionsWithDefaults,
              onChange: onOptionsChange,
            })}
          />

          <Divider />

          <ConfigSection title="Additional settings" isCollapsible>
            <AdvancedHttpSettings config={optionsWithDefaults} onChange={onOptionsChange} />

            <div className={styles.space} />

            <Field label="Custom query parameters" description="Add custom parameters to your queries.">
              <Input
                width={40}
                value={optionsWithDefaults.jsonData.queryParams}
                onChange={onParamsChange}
                spellCheck={false}
                placeholder="limit=100"
              />
            </Field>
          </ConfigSection>
        </>
      ) : null}

      {optionsWithDefaults.jsonData.storage === 'local' ? (
        <Field label="Path">
          <Input
            value={optionsWithDefaults.url}
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
