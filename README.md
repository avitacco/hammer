# jig

A Go-based reimplementation of the [Puppet Development Kit (PDK)](https://github.com/puppetlabs/pdk), built to be fast, self-contained, and free of Ruby runtime dependencies.

## Why jig?

PDK has been an essential tool for Puppet module authors for years. When 
Perforce moved PDK to a closed-source model, it created a real problem for teams
and individuals who depend on open tooling for their workflows. On top of that,
PDK carries a heavy Ruby runtime footprint, which adds friction to CI
environments and developer machines alike.

jig aims to replace the parts of PDK that matter most: scaffolding new modules,
building module packages, and cutting releases. It ships as a single static
binary with no external runtime required.

## Status

jig is under active development. The table below reflects the current state of
planned functionality.

| Command    | Subcommand     | Status     |
|------------|----------------|------------|
| `new`      | `module`       | ✅ Working  |
| `new`      | `class`        | 🔲 Planned |
| `new`      | `defined_type` | 🔲 Planned |
| `new`      | `fact`         | 🔲 Planned |
| `new`      | `function`     | 🔲 Planned |
| `new`      | `provider`     | 🔲 Planned |
| `new`      | `task`         | 🔲 Planned |
| `new`      | `test`         | 🔲 Planned |
| `new`      | `transport`    | 🔲 Planned |
| `build`    |                | 🔲 Planned |
| `release`  |                | 🔲 Planned |

## Installation

### Build from source

Requires Go 1.21 or later.

```bash
git clone https://github.com/avitacco/jig.git
cd jig
go build -o jig .
```

Move the resulting binary somewhere in your `$PATH`:

```bash
mv jig /usr/local/bin/
```

No other dependencies or runtimes needed.

## Usage

### `jig new module`

Scaffolds a new Puppet module with the standard directory structure and
metadata.

```
jig new module [flags]
```

jig will walk you through an interactive interview to collect module metadata.
If you have an author configured (see [Configuration](#configuration) below),
it will be used as the default.

**Flags:**

| Flag | Description                                                                                                                |
|---|----------------------------------------------------------------------------------------------------------------------------|
| `--force` | Overwrite an existing module directory. The existing directory is backed up with a timestamp before any files are written. |

**Module naming:** jig validates module names against Puppet's naming
conventions. Violations produce a warning but do not stop scaffolding.

### Templates

jig embeds default templates for all generated files. If you want to customize
them, place your own templates in `~/.config/jig/templates/` and jig will use
those instead, falling back to the built-in defaults for anything not
overridden.

## Configuration

jig looks for a config file at `~/.config/jig/config.toml`.

```toml
[author]
name = "yourname"
```

Setting `author.name` lets jig pre-fill the module author during the
`new module` interview. Additional configuration options will be documented here
as they are implemented.

## Contributing

Contributions are welcome. The project is in early stages, so the best place to
start is by opening an issue to discuss what you want to work on before sending
a PR.

### Project layout

```
.
├── main.go
├── commands/        # Cobra command definitions
└── internal/
    ├── build/
    ├── config/
    ├── forge/
    ├── module/      # Module metadata and validation
    ├── release/
    ├── scaffold/    # Scaffolding orchestration
    └── template/   # Template rendering with fallback logic
```

### Design notes for contributors

- Templates are embedded via `go:embed`. External templates in
- `~/.config/jig/templates/` take precedence.
- `--force` never deletes existing files outright. It creates a timestamped
- backup of the target directory first.
- Module name validation uses a `ValidationResult` type with an iota-based
- `Severity`. Violations at the `Warning` level do not halt execution.
- Config is handled with [Viper](https://github.com/spf13/viper).

## License

See [LICENSE](LICENSE).