# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- default order for `SecretProviderClass` resource kind
- new annotation `mia-platform.eu/apply-before-kinds` to override default resources application order
- support for jobs annotation `mia-platform.eu/await-completion` for waiting job completion after it has been applied on the cluster
- support for  resource`ExternalSecrets` for the annotation `mia-platform.eu/await-completion`

### Fixed

- fixed a bug in `createPatch` that caused all annotations to be deleted in the resulting patch if the target resource was annotated with `kubectl.kubernetes.io/last-applied-configuration`

## [1.1.0] - 2022-03-17

### Changed

- update to go 1.17

### Added

- ignore patches already in kustomization.yaml
- match as patch patch.ya?ml

### Fixes

- smart deploy don't force deploy pods on first update

## [1.0.3] - 2022-02-07

### Added

- support arm arch
- run image as the root user

## [1.0.2] - 2022-02-07

### Added

- improve error message when convert resource from yaml to json
- fix: splitting error on --- inside file

### Changed

- set QPS and burst for request to 500

## [1.0.1] - 2022-02-03

### Added

- feat: increase api-server throttling options

## [1.0.0] - 2021-12-28

- stable release

## [0.5.0] - 2021-08-24

## Added

- [BPSINS-27](https://makeitapp.atlassian.net/browse/BPSINS-27): add flag to skip namespace ensure when deploy

## [0.4.1] - 2021-03-30

### Fixed

- [BMP-940](https://makeitapp.atlassian.net/browse/BMP-940): fix annotation length by using an unique name, `mia-platform.eu/dependenciesChecksum`, for all dependencies and its value is a object of key-values of all the dependencies.

## [0.4.0] - 2021-03-17

### Added

- Add deploy type support, `smart deploy` or `deploy all`.

### Fixes

- [BMP-823](https://makeitapp.atlassian.net/browse/BMP-823): fix quote in configmap strings

## [0.3.2] - 2021-01-22

### Fixed

- [MPPS-57](https://makeitapp.atlassian.net/browse/MPPS-57): interpolation of variables inside single quotes

## [0.3.1] - 2020-11-25

### Added

- Add manual resource deletion

## [0.3.0] - 2020-11-02

### Added

- Add label `"app.kubernetes.io/managed-by": "mia-platform"`
- Unset original resource namespace
- Add resource deletion if no longer deployed with `mlp`

## [0.2.0] - 2020-10-20

### Added

- Add Job creation from CronJob

## [0.1.1] - 2020-10-14

### Changed

- Ignore unreadable or missing files passed as inputs to subcommands

## [0.1.0] - 2020-10-13

### Added

- Initial Release 🎉🎉🎉

[Unreleased]: https://github.com/mia-platform/mlp/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/mia-platform/mlp/compare/v1.0.3...v1.1.0
[1.0.3]: https://github.com/mia-platform/mlp/compare/v1.0.2...v1.0.3
[1.0.2]: https://github.com/mia-platform/mlp/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/mia-platform/mlp/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/mia-platform/mlp/compare/v1.0.0
