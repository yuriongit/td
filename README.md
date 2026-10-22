# Td

Td, "today", is a simple and straight-forward command-line
tool for my task management.

_Td was built for my personal use for keep tracking of what I need done for
the day._

## Features

1. Initializes the global configuration directory `~/.td` with a prepared .env
   file
2. Interactive task creation
3. Interactive task updates
4. Task completion
5. Task retrieval (all tasks)
6. Task deletion
7. Structured output

## Configuration

td's configuration directory, `~/.td`, holds environment variables that td
relies on are stored in the configuration directory. This specifically allows
enables global usage.

- `~/.td` directory layout document (coming soon): [/docs/td.md](./docs/td.md)

## Infrastructure

| Layer   | Tool                  |
| ------- | --------------------- |
| Main    | Go, Cobra             |
| Tooling | golangci-lint, Go     |
| UI      | PostgreSQL (Supabase) |

## Requirements

- Go 1.27.1 or later

## Quick Start

Clone the repo & build the binary:

```bash
git clone https://github.com/yuriongit/td.git
cd td
go build
go install
```

Start:

```bash
td help
```

Init td and create a task:

```bash
td init
td new
```

## Images

To see a preview of td, view the [/docs/preview.md](./docs/preview.md) document.

_Images directory - [/.github/images](./github/images)_

## Docs

- [/docs/architecture.md](./docs/architecture.md)
- [/docs/planned.md](./docs/planned.md)

## License

MIT
