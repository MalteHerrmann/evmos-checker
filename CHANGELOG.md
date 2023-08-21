# Changelog

## Unreleased

- Add methods to run a command in a subprocess
- Add UI to show the progress of the checks
    - use [termui](https://github.com/gizak/termui)
    - have a column per check or depending on windows have 3, 2 columns
      - `make lint`
      - `make check-licenses`
      - `make proto-all`
      - `make test-unit`
      - `make test-e2e`
- Have shortcuts for certain functionality (e.g. initial and target version for E2E tests)