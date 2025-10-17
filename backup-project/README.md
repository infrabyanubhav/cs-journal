# Backup Script

A bash script that creates compressed backups of recently modified files from a target directory.

## Features

- Creates timestamped backup archives (backup-{timestamp}.tar.gz)
- Only backs up files modified within the last 24 hours
- Validates input directories before processing
- Moves backup files to specified destination directory

## Usage

```bash
./backup.sh <target_directory> <destination_directory>
```

### Example

```bash
# Backup recently modified files from /home/user/documents to /backups
./backup.sh /home/user/documents /backups
```

## Requirements

- Bash shell
- tar command (for creating compressed archives)
- Read access to target directory
- Write access to destination directory

## How it works

1. Validates that exactly 2 arguments are provided
2. Checks that both directories exist
3. Gets current timestamp for backup filename
4. Finds all files modified in the last 24 hours
5. Creates a compressed tar archive with those files
6. Moves the backup file to the destination directory

## Output

Creates a file named `backup-{timestamp}.tar.gz` in the destination directory containing all recently modified files from the target directory.