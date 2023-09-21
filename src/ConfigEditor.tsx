import { DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { DataSourceHttpSettings, Field, InlineField, LegacyForms, Legend, RadioButtonGroup } from '@grafana/ui';
import defaults from 'lodash/defaults';
import React, { ChangeEvent } from 'react';
import { CSVDataSourceOptions, defaultOptions } from './types';

const { Input, FormField } = LegacyForms;

interface Props extends DataSourcePluginOptionsEditorProps<CSVDataSourceOptions> {}

/**
 * ConfigEditor lets the user configure connection details like the URL or
 * authentication.
 */
export const ConfigEditor: React.FC<Props> = ({ options, onOptionsChange }) => {
  const jsonData = defaults(options.jsonData, defaultOptions);

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
      <Field>
        <RadioButtonGroup
          options={[
            { label: 'HTTP', value: 'http' },
            { label: 'Local', value: 'local' },
          ]}
          value={jsonData.storage}
          onChange={onStorageChange}
        />
      </Field>

      {jsonData.storage === 'http' ? (
        <>
          {/* DataSourceHttpSettings handles most the settings for connecting over
      HTTP. */}
          <DataSourceHttpSettings
            defaultUrl="http://localhost:8080"
            dataSourceConfig={options}
            onChange={onOptionsChange}
          />

          {/* The Grafana proxy strips query parameters from the URL set in
      DataSourceHttpSettings. To support custom query parameters, the user need
      to set them explicitly.  */}
          <h3 className="page-heading">Misc</h3>
          <div className="gf-form-group">
            <div className="gf-form-inline">
              <div className="gf-form max-width-30">
                <FormField
                  label="Custom query parameters"
                  labelWidth={14}
                  tooltip="Add custom parameters to your queries."
                  inputEl={
                    <Input
                      className="width-25"
                      value={jsonData.queryParams}
                      onChange={onParamsChange}
                      spellCheck={false}
                      placeholder="page=1&limit=100"
                    />
                  }
                />
              </div>
            </div>
          </div>
        </>
      ) : null}

      {jsonData.storage === 'local' ? (
        <>
          <Legend>Local</Legend>
          <InlineField label="Path" tooltip="Path to the CSV file.">
            <Input value={options.url} onChange={onLocalPathChange} spellCheck={false} />
          </InlineField>
        </>
      ) : null}
    </>
  );
};
