# MGnify CLI Tool

A command-line interface for fetching and processing data from the MGnify API, with automatic pagination handling and CSV export capabilities.

## Features

- **API Interaction**: Retrieve data from complex MGnify endpoints
- **Pagination Handling**: Automatically follows "next" links for complete datasets
- **JSON Flattening**: Converts nested JSON structures to flat CSV format
- **Flexible Querying**:
  - Positional arguments for endpoint construction
  - Multiple query parameters supported via flags
  - Automatic URL encoding of special characters
- **CSV Export**: Configurable output to file or stdout
- **Modular Design**: Separated components for API client, data processing, and output

## Installation

### Using Go (1.16+ required)
```bash
go install github.com/yourusername/mgnify-cli@latest
```

### From Source
```bash
git clone https://github.com/yourusername/mgnify-cli.git
cd mgnify-cli
go build -o mgnify-cli main.go
sudo mv mgnify-cli /usr/local/bin/
```

## Usage

### Basic Command Structure
```bash
mgnify-cli [FLAGS] ENDPOINT_PARTS...
```

### Positional Arguments
| Argument        | Description                                  |
|-----------------|----------------------------------------------|
| `ENDPOINT_PARTS`| API endpoint components (joined with '/')    |
|                 | Example: `studies ERP009703 samples` → `studies/ERP009703/samples` |

### Flags
| Flag            | Description                                  |
|-----------------|----------------------------------------------|
| `-o, --output`  | Output file path (default: stdout)           |
| `-p`            | Query parameters (key=value format)          |
| `-h, --help`    | Show help message                            |
| `-v, --version` | Show version information                     |

## Examples

### Fetch biome studies to CSV file
```bash
mgnify-cli biomes MGYS00005292 studies -o biome_studies.csv
```

### Get samples with multiple parameters
```bash
mgnify-cli samples \
  -p experiment_type=metagenomic \
  -p biome_name=marine \
  -p page_size=500 \
  -o marine_samples.csv
```

### Search recent host-associated studies
```bash
mgnify-cli studies recent \
  -p lineage="root:Host-associated" \
  -p min_date=2023-01-01
```

## Configuration

### Base URL
Modify the API base URL in `main.go` (default: `https://api.mgnify.example/v1/`):
```go
const baseURL = "https://api.mgnify.example/v1/"
```

## License

MIT License - see [LICENSE](./LICENSE) for details

## Acknowledgements

- MGnify API team for providing biological data
- Go standard library contributors
- CSV format specification ([RFC 4180](https://tools.ietf.org/html/rfc4180))
