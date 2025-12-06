# GoToastify

A simple Go CLI application that creates Windows toast notifications.

## Features

- Display toast notifications from the command line
- Customize title, message, duration, and app ID
- Display images or icons in notifications
- Simple and intuitive CLI interface

## Installation

### Prerequisites
- Go 1.21 or later
- Windows OS (for toast notifications)

### Build from Source

```bash
git clone https://github.com/borsTiHD/go-toastify.git
cd go-toastify
go mod download
go build -o go-toastify.exe
```

## Usage

### Display a notification

Basic usage with required message:
```bash
go-toastify show -message "Hello, World!"
```

With custom title:
```bash
go-toastify show -title "Important" -message "This is important!"
```

With all options:
```bash
go-toastify show -title "Alert" -message "Check this out" -duration 10 -image "C:\path\to\icon.png"
```

### Show version

```bash
go-toastify version
```

### Show help

```bash
go-toastify help
```

## Command Options

### show
Display a toast notification

Flags:
- `-title string` - Title of the notification (default: "Notification")
- `-message string` - Message content (required)
- `-duration int` - Duration in seconds to display (default: 5)
- `-image string` - Path to an image/icon file to display in the notification

## Examples

```bash
# Simple notification
go-toastify show -message "Build completed successfully"

# Notification with custom title
go-toastify show -title "Build Complete" -message "Your build finished at 3:45 PM"

# Longer duration notification
go-toastify show -title "Reminder" -message "Don't forget the meeting" -duration 15

# Notification with image/icon
go-toastify show -title "Alert" -message "Important update" -image "C:\path\to\icon.png"

# Notification with all options
go-toastify show -title "Success" -message "Task completed" -duration 10 -image "C:\icons\success.png"
```

## License

GNU GENERAL PUBLIC LICENSE - See LICENSE file for details

## Author

borsTiHD
