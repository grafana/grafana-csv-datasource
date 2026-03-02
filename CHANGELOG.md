# Changelog

## 0.8.2

⚙️ Updated backend/go version to 1.26.0 from 1.25.6 which resolves CVE-2025-68121

⚙️ Updated frontend & backend dependencies

## 0.8.1

⚙️ Chore: Frontend and backend dependency updates

## 0.8.0

🚀 Update dependencies for react-19 upgrade

## 0.7.4

🐛 Update various dependency versions

## 0.7.3

🚀 Add smoke tests
🐛 Fix: Query editor broken because of dependency (in v0.7.2)

## 0.7.2

⚙️ Chore: Replaced custom timezone picker with Grafana UI component
⚙️ Chore: Updated frontend deps
⚙️ Chore: Updated backend plugin sdk `grafana-plugin-sdk-go`
⚙️ Chore: Update create plugin config
⚙️ Chore: Pinned github workflow actions

## 0.7.1

🐛 Fix: Add error source for http errors
🐛 Fix: Panics `index out of range [0] with length 0`
⚙️ Chore: Bump create plugin to 5.26.4

## 0.7.0

🚀 Refactor backend to handle error source
🚀 Bump frontend dependencies and bump grafana dependency

## 0.6.25

⚙️ Chore: Remove props mutations
⚙️ Chore: Update dependencies

## v0.6.24 - 2025-03-21

- 🛡️ **Security**: Update `grafana-plugin-sdk-go` to 0.274.0
- 🛡️ **Security**: Update `go` to 1.24.1
- 🛡️ **Security**: `golang.org/x/net` from 0.34.0 to 0.36.0
- 🛡️ **Security**: Update frontend dependencies

## v0.6.23 - 2025-02-07

- Update documentation about the deprecation

## v0.6.22 - 2025-01-17

- 🛡️ **Security**: `golang.org/x/net` from 0.30.0 to 0.33.0
- 🛡️ **Security**: Update `grafana-plugin-sdk-go` to 0.263.0
- 🛡️ **Security**: Build and publish pipelines uses latest go lang version `1.23.5` which includes security fixes to the `crypto/x509` and `net/http` packages ( CVE-2024-45341 and CVE-2024-45336 ). More details can be found [here](https://groups.google.com/g/golang-announce/c/sSaUhLA-2SI)

## v0.6.21 - 2024-10-30

- 🛡️ **Security**: Update `uplot` to 1.6.31

## v0.6.20 - 2024-10-18

- 🛡️ **Security**: Update `grafana-plugin-sdk-go` to 0.255
- 🛡️ **Security**: Bump `dompurify` from 3.1.0 to 3.1.6
- 🛡️ **Security**: Bump `path-to-regexp` from 1.8.0 to 1.9.0

## v0.6.19 - 2024-09-03

- 🛡️ **Security**: Bump `micromatch` from 4.0.5 to 4.0.8
- 🛡️ **Security**: Bump `webpack` from 5.86.0 to 5.94.0
- 🛡️ **Security**: Bump `fast-loops` from 1.1.3 to 1.1.4
- 🛡️ **Security**: Bump `ws` from 8.14.2 to 8.17.1
- 🛡️ **Security**: Bump `braces` from 3.0.2 to 3.0.3

## v0.6.18 - 2024-05-14

- ⚙️ **Chore**: Updated the eslint-plugin-prettier dependency

## v0.6.17 - 2024-04-18

- ⚙️ **Chore**: Bump grafana-plugin-sdk-go from `v0.220.0` to `v0.225.0`
- ⚙️ **Chore**: Updated grafana frontend runtime dependencies

## v0.6.16 - 2024-04-08

- ⚙️ **Chore**: Bump grafana-plugin-sdk-go from `v0.197.0` to `v0.220.0`

## v0.6.15 - 2024-03-07

- ⚙️ **Chore**: Build with go 1.22

## v0.6.14 - 2024-02-14

- ⚙️ **Chore**: Maintenance

## v0.6.13 - 2024-02-14

- 🛡️ **Security**: More robust URL handling: disallow changing the configured hostname in the query editor ( CVE-2023-5122 )

## v0.6.12 - 2023-12-19

- ⚙️ **Docs**: Documentation website moved from [github pages](https://grafana.github.io/grafana-csv-datasource) to [grafana.com/docs/plugins](https://grafana.com/docs/plugins/marcusolsson-csv-datasource/latest/) page
- ⚙️ **Chore**: Updated grafana-plugin-sdk-go from `v0.193.0` to `v0.197.0`

## v0.6.11 - 2023-11-21

- **Feature**: Update configuration page to follow best practices
- ⚙️ **Chore**: Upgrade grafana-plugin-sdk-go to latest
- ⚙️ **Chore**: Added lint github workflow
- ⚙️ **Chore**: Update legacy form styling
- ⚙️ **Chore**: Update readme and documentation

## v0.6.10 - 2023-10-24

- 🐛 **Fix**: More robust local file mode handling

## v0.6.9 - 2023-10-19

- 🐛 **Fix**: Correct query field behavior on older Grafana versions
- ⚙️ **Chore**: Upgrade dependencies

## v0.6.8 - 2023-10-18

- ⚙️ **Chore**: Upgrade dependencies

## v0.6.7 - 2023-10-09

- ⚙️ **Chore**: Upgrade dependencies
- Added feature tracking

## v0.6.6 - 2023-08-23

- 🐛 **Fix**: Consistently apply field names

## v0.6.5 - 2023-05-03

- ⚙️ **Chore**: backend binaries are now compiled with go 1.20.4

## v0.6.4 - 2023-04-19

- ⚙️ **Chore**: backend binaries are now compiled with go 1.20.3

## v0.6.3 - 2021-12-03

- ⚙️ **Chore**: backend binaries are now compiled with go 1.19.3
- ⚙️ **Chore**: frontend npm dependencies updated
- ⚙️ **Chore**: added spellcheck

## v0.6.2 - 2021-10-14

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.6.1...v0.6.2)

- Fixed the broken docs and links

## v0.6.1 - 2021-06-22

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.6.0...v0.6.1)

### Bug fixes

- allow_local_mode accepts any value

## v0.6.0 - 2021-06-21

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.5.0...v0.6.0)

### Enhancements

- Disable local mode by default. To use local mode, allow it in your grafana.ini:

  ```ini
  [plugin.marcusolsson-csv-datasource]
  allow_local_mode = true
  ```

## v0.5.0 - 2021-03-21

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.4.1...v0.5.0)

### Enhancements

- Improved query editor with support for HTTP params, headers, and body.
- Add support for relative paths ([#69](https://github.com/grafana/grafana-csv-datasource/issues/69))
- Add support for decimal separators ([#43](https://github.com/grafana/grafana-csv-datasource/issues/43))
- **EXPERIMENTAL:** Add support for regular expressions in field names ([#68](https://github.com/grafana/grafana-csv-datasource/issues/68)). Must be enabled in the Experimental tab in the query editor.

### Bug fixes

- **BREAKING:** Remove default Accept header ([#56](https://github.com/grafana/grafana-csv-datasource/issues/56)). If your data source expects `Accept: text/csv` on the request, you now need to add it yourself in the Params tab.

## v0.4.1 - 2021-03-21

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.4.0...v0.4.1)

### Bug fixes

- Wrong data format is detected

## v0.4.0 - 2021-03-21

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.3.3...v0.4.0)

### Enhancements

- Add support for annotation queries
- Add support for variables queries ([#30](https://github.com/grafana/grafana-csv-datasource/issues/30))
- Upgrade @grafana/\* packages
- Upgrade Grafana Go SDK

## v0.3.3 - 2021-02-05

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.3.2...v0.3.3)

### Bug fixes

- Default to HTTP if no storage type has been set

## v0.3.2 - 2021-02-03

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.3.1...v0.3.2)

### Bug fixes

- Allow lazy quotes ([#17](https://github.com/grafana/grafana-csv-datasource/issues/17))

## v0.3.1

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.3.0...v0.3.1)

### Enhancements

- Update grafana-plugin-sdk-go to v0.83.0

### Bug fixes

- Ignore empty custom HTTP headers
- 🐛 **Fix**: duplicate JSON tag for TLS skip verify

## v0.3.0

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.2.0...v0.3.0)

### Enhancements

- Add ARM support ([#13](https://github.com/grafana/grafana-csv-datasource/issues/13))

### Bug fixes

- Windows: Paths with backslashes don't work ([#14](https://github.com/grafana/grafana-csv-datasource/issues/14))

## v0.2.0

[Full changelog](https://github.com/grafana/grafana-csv-datasource/compare/v0.1.0...v0.2.0)

### Enhancements

- Add support for local CSV files ([#6](https://github.com/grafana/grafana-csv-datasource/issues/6))
- Add a default Accept header for text/csv

## v0.1.0

Initial release. Not fit for production use.
