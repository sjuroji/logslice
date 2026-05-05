# logslice

Fast log file slicer and filter tool with time-range and pattern support.

## Installation

```bash
go install github.com/yourusername/logslice@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/logslice.git && cd logslice && go build ./...
```

## Usage

```bash
# Slice logs between two timestamps
logslice -from "2024-01-15 08:00:00" -to "2024-01-15 09:00:00" app.log

# Filter by pattern within a time range
logslice -from "2024-01-15 08:00:00" -to "2024-01-15 09:00:00" -pattern "ERROR" app.log

# Read from stdin
cat app.log | logslice -from "2024-01-15 08:00:00" -to "2024-01-15 09:00:00"

# Write output to a file
logslice -from "2024-01-15 08:00:00" -to "2024-01-15 09:00:00" -o output.log app.log
```

### Flags

| Flag | Description |
|------|-------------|
| `-from` | Start of time range (RFC3339 or common log formats) |
| `-to` | End of time range |
| `-pattern` | Regex pattern to filter log lines |
| `-o` | Output file (defaults to stdout) |
| `-tz` | Timezone for timestamp parsing (default: UTC) |

## Features

- Blazing fast line-by-line processing with minimal memory usage
- Supports common log timestamp formats automatically
- Regex pattern filtering combined with time-range slicing
- Reads from files or stdin, writes to files or stdout

## License

MIT — see [LICENSE](LICENSE) for details.