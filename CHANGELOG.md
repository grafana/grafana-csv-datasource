# Changelog

## 0.6.5 (2023-05-03)

- **Chore** backend binaries are now compiled with go 1.20.4

## 0.6.4 (2023-04-19)

- **Chore** backend binaries are now compiled with go 1.20.3

## 0.6.3 (2021-12-03)

- **Chore** backend binaries are now compiled with go 1.19.3
- **Chore** frontend npm dependencies updated
- **Chore** added spellcheck

## 0.6.2 (2021-10-14)

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.6.1...v0.6.2)

- Fixed the broken docs and links

## 0.6.1 (2021-06-22)

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.6.0...v0.6.1)

### Bug fixes

- allow_local_mode accepts any value

## 0.6.0 (2021-06-21)

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.5.0...v0.6.0)

### Enhancements

- Disable local mode by default. To use local mode, allow it in your grafana.ini:

  ```ini
  [plugin.marcusolsson-csv-datasource]
  allow_local_mode = true
  ```

## 0.5.0 (2021-03-21)

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.4.1...v0.5.0)

### Enhancements

- Improved query editor with support for HTTP params, headers, and body.
- Add support for relative paths ([#69](https://github.com/grafana/grafana-csv-datasource/issues/69))
- Add support for decimal separators ([#43](https://github.com/grafana/grafana-csv-datasource/issues/43))
- **EXPERIMENTAL:** Add support for regular expressions in field names ([#68](https://github.com/grafana/grafana-csv-datasource/issues/68)). Must be enabled in the Experimental tab in the query editor.

### Bug fixes

- **BREAKING:** Remove default Accept header ([#56](https://github.com/grafana/grafana-csv-datasource/issues/56)). If your data source expects `Accept: text/csv` on the request, you now need to add it yourself in the Params tab.

## 0.4.1 (2021-03-21)

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.4.0...v0.4.1)

### Bug fixes

- Wrong data format is detected

## 0.4.0 (2021-03-21)

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.3.3...v0.4.0)

### Enhancements

- Add support for annotation queries
- Add support for variables queries ([#30](https://github.com/grafana/grafana-csv-datasource/issues/30))
- Upgrade @grafana/\* packages
- Upgrade Grafana Go SDK

## 0.3.3 (2021-02-05)

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.3.2...v0.3.3)

### Bug fixes

- Default to HTTP if no storage type has been set

## 0.3.2 (2021-02-03)

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.3.1...v0.3.2)

### Bug fixes

- Allow lazy quotes ([#17](https://github.com/grafana/grafana-csv-datasource/issues/17))

## 0.3.1

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.3.0...v0.3.1)

### Enhancements

- Update grafana-plugin-sdk-go to v0.83.0

### Bug fixes

- Ignore empty custom HTTP headers
- Fix duplicate JSON tag for TLS skip verify

## 0.3.0

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.2.0...v0.3.0)

### Enhancements

- Add ARM support ([#13](https://github.com/grafana/grafana-csv-datasource/issues/13))

### Bug fixes

- Windows: Paths with backslashes don't work ([#14](https://github.com/grafana/grafana-csv-datasource/issues/14))

## 0.2.0

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.1.0...v0.2.0)

### Enhancements

- Add support for local CSV files ([#6](https://github.com/grafana/grafana-csv-datasource/issues/6))
- Add a default Accept header for text/csv

## 0.1.0

Initial release. Not fit for production use.
