# Note Usage Guide

This CLI helps you manage your notes quickly without leaving the terminal.

## Storage Structure

Notes are stored using the following pattern:
`[storage_dir]/[DDMMYYYY]/note[HHMMSS]::[Title].[ext]`

Example: `~/notes/13012026/note103005::Learn-Go.md`

---

## Basic Commands

### 1. Create a New Note (`create`, `c`)
By default, running `note` will create a new note.

```bash
# Create a default note (txt)
note

# Create with a title and markdown format
note c -t "Learn Go" -e md

# Create without opening the editor immediately
note c -t "Important Idea" -c
```

**Flags:**
- `-t, --title`: The title of the note.
- `-e, --ext`: File extension/format (md, txt, json, html).
- `-c`: Create the file only, do not open the editor.

---

### 2. List Notes (`list`, `ls`)
View the list of notes created for a specific day.

```bash
# List today's notes
note ls

# List notes from 2 days ago
note ls --ago 2
```

---

### 3. Open a Note (`open`, `o`)
Open a note for editing.

```bash
# Open the latest note from today
note o

# Open a note by its index in the list
note o 1
# or use the flag
note o -i 1
```

---

### 4. Format-based Shortcuts
You can use file extensions as commands for quick creation:

```bash
# Equivalent to 'note c -e md'
note md -t "Quick note"

# Equivalent to 'note c -e json'
note json -t "Config data"
```

---

## History Management (`--ago`, `-a`)
Most commands support the `-a` flag to interact with notes from previous days.

```bash
# See notes from yesterday
note ls -a 1

# Open a note from the day before yesterday
note o 1 -a 2
```

---

## Automatic Title Update
If you create a note without a title (`Untitled`), the app will check the first line of content after you close the editor. If that line contains text, the app will automatically rename the file based on that content. Perfect for "write first, name later" workflows!
