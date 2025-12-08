package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/borsTiHD/go-toastify/toast"
)

const (
	VERSION = "1.0.0"
)

// Run executes the CLI handler and returns true if a CLI command was processed
func Run(args []string) bool {
	if len(args) < 2 {
		return false
	}

	// Define subcommands
	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)

	// Flags for 'show' command
	title := showCmd.String("title", "Notification", "Title of the toast notification")
	message := showCmd.String("message", "", "Message content of the toast notification")
	duration := showCmd.Int("duration", 5, "Duration in seconds to display the notification")
	imagePath := showCmd.String("image", "", "Path to an image file to display in the notification")

	// Handle subcommands
	switch args[1] {
	case "show":
		showCmd.Parse(args[2:])
		if *message == "" {
			fmt.Println("Error: message flag is required")
			fmt.Println("\nUsage: go-toastify show -title \"Title\" -message \"Message\" [-duration 5] [-image \"path/to/image\"]")
			os.Exit(1)
		}
		err := toast.DisplayToast(*title, *message, *duration, *imagePath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		if *imagePath != "" {
			fmt.Printf("Toast notification displayed: %s - %s (with image: %s)\n", *title, *message, *imagePath)
		} else {
			fmt.Printf("Toast notification displayed: %s - %s\n", *title, *message)
		}
		return true

	case "version":
		versionCmd.Parse(args[2:])
		fmt.Printf("go-toastify version %s\n", VERSION)
		return true

	case "help", "-h", "--help":
		PrintHelp()
		return true

	default:
		fmt.Printf("Unknown command: %s\n\n", args[1])
		PrintHelp()
		os.Exit(1)
		return true
	}
}

// PrintHelp displays the CLI help message
func PrintHelp() {
	help := `go-toastify - Windows Toast Notification CLI

Usage:
  go-toastify <command> [options]

Commands:
  show      Display a toast notification
  version   Show version information
  help      Show this help message

Examples:
  go-toastify show -title "Hello" -message "This is a test notification"
  go-toastify show -title "Info" -message "Important message" -duration 10
  go-toastify show -title "Build Complete" -message "Your project built successfully"
  go-toastify show -title "Alert" -message "Check this!" -image "C:\\path\\to\\icon.png"
  go-toastify version
  go-toastify help

Flags for 'show' command:
  -title string        Title of the notification (default: "Notification")
  -message string      Message content (required)
  -duration int        Duration in seconds to display (default: 5)
  -image string        Path to an image/icon file to display in the notification

Note: Requires Windows 10+ with toast notifications enabled.`
	fmt.Println(help)
}
