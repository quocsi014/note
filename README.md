# Note

A dead-simple CLI tool for managing personal notes through organized local files.

## Key Features

- **Date-based Organization**: Automatically creates and stores notes in folders organized by date (`DDMMYYYY`).
- **Smart Naming**: Automatically generates filenames with a time prefix (`noteHHMMSS`) and title.
- **Auto-Sync Title**: Automatically updates the filename's title based on the first line of the note content if no title was specified during creation.
- **Multi-format Support**: Supports Markdown (.md), Text (.txt), JSON (.json), and HTML (.html).
- **Time-Travel**: Easily view or create notes for previous days using the `--ago` flag.
- **Configurable**: Customize your favorite editor (Vim, VS Code, Nano...) via a simple configuration file.

## Installation

### Option 1: Using Go (Recommended)

If you have Go installed on your system:

```bash
go install .
```

### Option 2: Binary File

1. Download the `note` binary.
2. Grant execution permission: `chmod +x note`.
3. Move the binary to a directory in your `$PATH` (e.g., `/usr/local/bin` or `~/.local/bin`).

## Usage Guide

For detailed command usage, check out: **[USAGE.md](usage.md)**

Quick Command Summary:

- `note` or `note c`: Create a new note.
- `note ls`: List today's notes.
- `note o`: Open the latest note or by index.

## Configuration

The config file is located at: `~/.config/note/config.jsonc` (Linux).
You can customize:

- `editor`: The command to open your editor (default is `vi`).
- `storage_dir`: The directory where notes are stored (default is `~/notes`).

---

Built with ❤️ by a Golang Backend Developer.
