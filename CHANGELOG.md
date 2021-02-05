# Changelog

## 0.3.3 (2021-02-05)

[Full changelog](https://github.com/marcusolsson/grafana-csv-datasource/compare/v0.3.2...v0.3.3)

### Bug fixes

- Default to HTTP if no storage type has been set

## 0.3.2 (2021-02-03)

[Full changelog](https://github.com/marcusolsson/grafana-csv-datasource/compare/v0.3.1...v0.3.2)

### Bug fixes

- Allow lazy quotes ([#17](https://github.com/marcusolsson/grafana-csv-datasource/issues/17))

## 0.3.1

[Full changelog](https://github.com/marcusolsson/grafana-csv-datasource/compare/v0.3.0...v0.3.1)

### Enhancements

- Update grafana-plugin-sdk-go to v0.83.0

### Bug fixes

- Ignore empty custom HTTP headers
- Fix duplicate JSON tag for TLS skip verify

## 0.3.0

[Full changelog](https://github.com/marcusolsson/grafana-csv-datasource/compare/v0.2.0...v0.3.0)

### Enhancements

- Add ARM support ([#13](https://github.com/marcusolsson/grafana-csv-datasource/issues/13))

### Bug fixes

- Windows: Paths with backslashes don't work ([#14](https://github.com/marcusolsson/grafana-csv-datasource/issues/14))

## 0.2.0

[Full changelog](https://github.com/marcusolsson/grafana-csv-datasource/compare/v0.1.0...v0.2.0)

### Enhancements

- Add support for local CSV files ([#6](https://github.com/marcusolsson/grafana-csv-datasource/issues/6))
- Add a default Accept header for text/csv

## 0.1.0

Initial release. Not fit for production use.
