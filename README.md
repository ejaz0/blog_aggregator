# Blog Aggregator CLI (gator)

A command-line tool for aggregating and managing blog feeds.

## Prerequisites

- **Go 1.25+**
  - Verify: `go version`
- **PostgreSQL 14+** (or your installed version)
  - Verify: `psql --version`

## Installation

### Option 1: Install from GitHub
```bash
go install github.com/ejaz0/blog_aggreator@latest
```
Ensure `$GOPATH/bin` (or `go env GOPATH`/bin) is in your PATH so `gator` is runnable.

### Option 2: Clone and Build
```bash
git clone https://github.com/ejaz0/blog_aggreator.git
cd blog_aggreator
go build -o gator .
```

## Configuration

The app reads a JSON config at: `~/.gatorconfig.json`

### Required fields:
- `db_url`: PostgreSQL connection string
- `current_user_name`: set by the app after login/register

### Minimal example:
Create `~/.gatorconfig.json` with:
```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

## Database Setup

1. Create a database:
   ```bash
   createdb gator
   ```

2. Apply your schema/migrations as you normally do (your project includes SQL in `sql/schema`). If you're using goose manually, run goose against your DB. Otherwise, ensure the tables from the `sql/schema` directory exist before using the CLI.

## Usage

### Get Help
Running with no args will show usage:
```bash
gator <command> [args...]
```

### Authentication and Users
- **Register**: `gator register <username>`
- **Login** (sets current user in config): `gator login <username>`
- **List users**: `gator users`

### Feed Management
- **Add a feed** (requires login): `gator addfeed <name> <url>`
- **List feeds**: `gator feeds`
- **Follow a feed** (requires login): `gator follow <url>`
- **List following** (requires login): `gator following`
- **Unfollow** (requires login): `gator unfollow <url>`

### Aggregation and Browsing
- **Aggregate** (fetch/process): `gator agg`
- **Browse posts**: `gator browse`

### Admin/Reset
- **Reset database**: `gator reset` ⚠️ **Warning: drops data**

## Development vs Production

- **Development**: `go run .` from the repo root
- **Production**: use the installed binary `gator` after `go install`

## Troubleshooting

### "Usage: cli <command> [args…]"
You didn't pass a command. Run `gator users` or `gator --help` style commands.

### "error reading config"
Ensure `~/.gatorconfig.json` exists and is valid JSON with a correct `db_url`.

### DB connection failures
Verify the `db_url`, DB is running, network reachable, and credentials correct.

### "command not found" from shell
Ensure `$GOPATH/bin` is on PATH and the binary name you built/installed is what you're invoking.
