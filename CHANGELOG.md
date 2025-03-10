# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

## [1.0.4] - 2025-01-22

- Dependency updates

## [1.0.3] - 2023-11-23

## [1.0.2] - 2023-04-17

### Changed

- Add hint about `values.yaml` being auto-generated so that developers do not try to edit manually

## [1.0.1] - 2023-03-22

- Fix module path.

## [1.0.0] - 2023-03-21

- Fix update action workflow trigger.

## [0.2.0] - 2023-03-09

- Improve update action workflow.
- Add a reusable composite GitHub action that calls `helm-values-gen`.
- Add version command.

## [0.1.0] - 2023-02-28

- Fix yaml marshaling.
- Dont add line breaks into long strings in yaml output.
- Print output to `stdout` instead of `stderr`.
- Add prototype.


[Unreleased]: https://github.com/giantswarm/helm-values-gen/compare/v1.0.4...HEAD
[1.0.4]: https://github.com/giantswarm/helm-values-gen/compare/v1.0.3...v1.0.4
[1.0.3]: https://github.com/giantswarm/helm-values-gen/compare/v1.0.2...v1.0.3
[1.0.2]: https://github.com/giantswarm/helm-values-gen/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/giantswarm/helm-values-gen/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/helm-values-gen/compare/v0.2.0...v1.0.0
[0.2.0]: https://github.com/giantswarm/helm-values-gen/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/helm-values-gen/releases/tag/v0.1.0
