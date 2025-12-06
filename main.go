package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	version = "1.0.0"
)

func main() {
	// Define subcommands
	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)

	// Flags for 'show' command
	title := showCmd.String("title", "Notification", "Title of the toast notification")
	message := showCmd.String("message", "", "Message content of the toast notification")
	duration := showCmd.Int("duration", 5, "Duration in seconds to display the notification")

	// Check if no arguments provided, show help
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	// Handle subcommands
	switch os.Args[1] {
	case "show":
		showCmd.Parse(os.Args[2:])
		if *message == "" {
			fmt.Println("Error: message flag is required")
			fmt.Println("\nUsage: go-toastify show -title \"Title\" -message \"Message\" [-duration 5]")
			os.Exit(1)
		}
		displayToast(*title, *message, *duration)

	case "version":
		versionCmd.Parse(os.Args[2:])
		fmt.Printf("go-toastify version %s\n", version)

	case "help", "-h", "--help", "":
		printHelp()

	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printHelp()
		os.Exit(1)
	}
}

func displayToast(title, message string, duration int) {
	// Use PowerShell to display toast notification on Windows
	// This is the native Windows way without external dependencies
	psCommand := fmt.Sprintf(`
[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.UI.Notifications.ToastNotification, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.Data.Xml.Dom.XmlDocument, System.Xml.XmlDocument, ContentType = WindowsRuntime] | Out-Null

$APP_ID = 'GoToastify'

$template = @"
<toast>
    <visual>
        <binding template="ToastText02">
            <text id="1">%s</text>
            <text id="2">%s</text>
        </binding>
    </visual>
</toast>
"@

$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($template)
$toast = New-Object Windows.UI.Notifications.ToastNotification $xml
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($APP_ID).Show($toast)
`, title, message)

	cmd := exec.Command("powershell.exe", "-NoProfile", "-Command", psCommand)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to display toast notification: %v\n", err)
	}

	// Wait for the specified duration
	time.Sleep(time.Duration(duration) * time.Second)

	fmt.Printf("Toast notification displayed: %s - %s\n", title, message)
}

func printHelp() {
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
  go-toastify version
  go-toastify help

Flags for 'show' command:
  -title string        Title of the notification (default: "Notification")
  -message string      Message content (required)
  -duration int        Duration in seconds to display (default: 5)

Note: Requires Windows 10+ with toast notifications enabled.`
	fmt.Println(help)
}
