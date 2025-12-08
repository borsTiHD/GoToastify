package toast

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"time"
)

// DisplayToast shows a Windows toast notification using PowerShell
func DisplayToast(title, message string, duration int, imagePath string) error {
	// Build the toast template based on whether an image is provided
	var toastTemplate string
	if imagePath != "" {
		// Normalize Windows path to use forward slashes for XML
		imagePath = convertToAbsolutePath(imagePath)
		toastTemplate = fmt.Sprintf(`
<toast>
    <visual>
        <binding template="ToastImageAndText02">
            <image id="1" src="file:///%s"/>
            <text id="1">%s</text>
            <text id="2">%s</text>
        </binding>
    </visual>
</toast>
`, escapePath(imagePath), title, message)
	} else {
		toastTemplate = fmt.Sprintf(`
<toast>
    <visual>
        <binding template="ToastText02">
            <text id="1">%s</text>
            <text id="2">%s</text>
        </binding>
    </visual>
</toast>
`, title, message)
	}

	psCommand := fmt.Sprintf(`
[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.UI.Notifications.ToastNotification, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.Data.Xml.Dom.XmlDocument, System.Xml.XmlDocument, ContentType = WindowsRuntime] | Out-Null

$APP_ID = 'GoToastify'

$template = @"
%s
"@

$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($template)
$toast = New-Object Windows.UI.Notifications.ToastNotification $xml
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($APP_ID).Show($toast)
`, toastTemplate)

	cmd := exec.Command("powershell.exe", "-NoProfile", "-Command", psCommand)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to display toast notification: %w", err)
	}

	// Wait for the specified duration
	time.Sleep(time.Duration(duration) * time.Second)

	return nil
}

// DisplayToastAsync shows a toast notification without blocking
func DisplayToastAsync(title, message string, imagePath string) error {
	// Build the toast template based on whether an image is provided
	var toastTemplate string
	if imagePath != "" {
		imagePath = convertToAbsolutePath(imagePath)
		toastTemplate = fmt.Sprintf(`
<toast>
    <visual>
        <binding template="ToastImageAndText02">
            <image id="1" src="file:///%s"/>
            <text id="1">%s</text>
            <text id="2">%s</text>
        </binding>
    </visual>
</toast>
`, escapePath(imagePath), title, message)
	} else {
		toastTemplate = fmt.Sprintf(`
<toast>
    <visual>
        <binding template="ToastText02">
            <text id="1">%s</text>
            <text id="2">%s</text>
        </binding>
    </visual>
</toast>
`, title, message)
	}

	psCommand := fmt.Sprintf(`
[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.UI.Notifications.ToastNotification, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.Data.Xml.Dom.XmlDocument, System.Xml.XmlDocument, ContentType = WindowsRuntime] | Out-Null

$APP_ID = 'GoToastify'

$template = @"
%s
"@

$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($template)
$toast = New-Object Windows.UI.Notifications.ToastNotification $xml
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($APP_ID).Show($toast)
`, toastTemplate)

	cmd := exec.Command("powershell.exe", "-NoProfile", "-Command", psCommand)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to display toast notification: %w", err)
	}

	return nil
}

// convertToAbsolutePath converts a relative path to an absolute path
func convertToAbsolutePath(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		log.Printf("Failed to convert path to absolute: %v\n", err)
		return path
	}
	return abs
}

// escapePath escapes backslashes in Windows paths for use in XML
func escapePath(path string) string {
	escaped := ""
	for _, char := range path {
		if char == '\\' {
			escaped += "/"
		} else {
			escaped += string(char)
		}
	}
	return escaped
}
