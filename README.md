# Testris

A tool for splitting Go tests across multiple CI parallel jobs.

## Installation

```bash
go install github.com/mcncl/testris/cmd/testris@latest
```

## Usage

```bash
testris -index=0 -total=2 -dir=./
```

### In Buildkite

```yaml
steps:
  - label: "go test"
    command: |
      TEST_PATTERN=$(testris -index=$$BUILDKITE_PARALLEL_JOB -total=$$BUILDKITE_PARALLEL_JOB_COUNT)
      go test ./... -run "$TEST_PATTERN"
    parallelism: 2
```

## Flags

- `-index`: Current parallel index (0-based)
- `-total`: Total number of parallel runs
- `-dir`: Root directory to scan for tests (default: ".")
