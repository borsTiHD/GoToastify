package gui

import (
	"fmt"
	"os/exec"

	"github.com/borsTiHD/go-toastify/cli"
)

const HELP_TEXT = `GoToastify - Windows Toast Notification Tool
Version: %s

A simple Go CLI application that creates Windows toast notifications.

Features:
• Display toast notifications from the command line
• Customize title, message, duration, and app ID
• Display images or icons in notifications
• Simple and intuitive CLI interface

Prerequisites:
• Go 1.21 or later
• Windows OS (for toast notifications)

CLI Usage Examples:
• Basic: go-toastify show -message "Hello, World!"
• With title: go-toastify show -title "Important" -message "Message"
• Full: go-toastify show -title "Alert" -message "Content" -duration 10

Command Options for 'show':
• -title string    Title of the notification (default: "Notification")
• -message string  Message content (required)
• -duration int    Duration in seconds (default: 5)
• -image string    Path to an image/icon file

Author: borsTiHD
License: GNU General Public License`

// Run starts the GUI application using Windows Forms via PowerShell
func Run() {
	helpContent := fmt.Sprintf(HELP_TEXT, cli.VERSION)

	psScript := fmt.Sprintf(`
Add-Type -AssemblyName System.Windows.Forms
Add-Type -AssemblyName System.Drawing

$form = New-Object System.Windows.Forms.Form
$form.Text = 'GoToastify'
$form.Size = New-Object System.Drawing.Size(550, 550)
$form.StartPosition = 'CenterScreen'
$form.FormBorderStyle = 'FixedSingle'
$form.MaximizeBox = $false

# Title Label
$titleLabel = New-Object System.Windows.Forms.Label
$titleLabel.Location = New-Object System.Drawing.Point(10, 10)
$titleLabel.Size = New-Object System.Drawing.Size(515, 30)
$titleLabel.Text = 'GoToastify - Windows Toast Notification Tool'
$titleLabel.Font = New-Object System.Drawing.Font('Segoe UI', 14, [System.Drawing.FontStyle]::Bold)
$titleLabel.TextAlign = 'MiddleCenter'
$form.Controls.Add($titleLabel)

# Version Label
$versionLabel = New-Object System.Windows.Forms.Label
$versionLabel.Location = New-Object System.Drawing.Point(10, 40)
$versionLabel.Size = New-Object System.Drawing.Size(515, 20)
$versionLabel.Text = 'Version: %s'
$versionLabel.TextAlign = 'MiddleCenter'
$form.Controls.Add($versionLabel)

# Help Section Label
$helpSectionLabel = New-Object System.Windows.Forms.Label
$helpSectionLabel.Location = New-Object System.Drawing.Point(10, 70)
$helpSectionLabel.Size = New-Object System.Drawing.Size(515, 20)
$helpSectionLabel.Text = 'Help && Information'
$helpSectionLabel.Font = New-Object System.Drawing.Font('Segoe UI', 10, [System.Drawing.FontStyle]::Bold)
$form.Controls.Add($helpSectionLabel)

# Help TextBox
$helpTextBox = New-Object System.Windows.Forms.TextBox
$helpTextBox.Location = New-Object System.Drawing.Point(10, 95)
$helpTextBox.Size = New-Object System.Drawing.Size(515, 200)
$helpTextBox.Multiline = $true
$helpTextBox.ReadOnly = $true
$helpTextBox.ScrollBars = 'Vertical'
$helpTextBox.Font = New-Object System.Drawing.Font('Consolas', 9)
$helpTextBox.Text = @"
%s
"@
$form.Controls.Add($helpTextBox)

# Test Notification Section
$testSectionLabel = New-Object System.Windows.Forms.Label
$testSectionLabel.Location = New-Object System.Drawing.Point(10, 305)
$testSectionLabel.Size = New-Object System.Drawing.Size(515, 20)
$testSectionLabel.Text = 'Test Notification'
$testSectionLabel.Font = New-Object System.Drawing.Font('Segoe UI', 10, [System.Drawing.FontStyle]::Bold)
$form.Controls.Add($testSectionLabel)

# Title Input Label
$titleInputLabel = New-Object System.Windows.Forms.Label
$titleInputLabel.Location = New-Object System.Drawing.Point(10, 335)
$titleInputLabel.Size = New-Object System.Drawing.Size(60, 20)
$titleInputLabel.Text = 'Title:'
$form.Controls.Add($titleInputLabel)

# Title Input
$titleInput = New-Object System.Windows.Forms.TextBox
$titleInput.Location = New-Object System.Drawing.Point(75, 332)
$titleInput.Size = New-Object System.Drawing.Size(450, 25)
$titleInput.Text = 'Test Notification'
$form.Controls.Add($titleInput)

# Message Input Label
$messageInputLabel = New-Object System.Windows.Forms.Label
$messageInputLabel.Location = New-Object System.Drawing.Point(10, 365)
$messageInputLabel.Size = New-Object System.Drawing.Size(60, 20)
$messageInputLabel.Text = 'Message:'
$form.Controls.Add($messageInputLabel)

# Message Input
$messageInput = New-Object System.Windows.Forms.TextBox
$messageInput.Location = New-Object System.Drawing.Point(75, 362)
$messageInput.Size = New-Object System.Drawing.Size(450, 60)
$messageInput.Multiline = $true
$messageInput.Text = 'This is a test toast notification from GoToastify GUI!'
$form.Controls.Add($messageInput)

# Send Button
$sendButton = New-Object System.Windows.Forms.Button
$sendButton.Location = New-Object System.Drawing.Point(10, 435)
$sendButton.Size = New-Object System.Drawing.Size(515, 35)
$sendButton.Text = 'Send Test Notification'
$sendButton.Font = New-Object System.Drawing.Font('Segoe UI', 10, [System.Drawing.FontStyle]::Bold)
$sendButton.BackColor = [System.Drawing.Color]::FromArgb(0, 120, 212)
$sendButton.ForeColor = [System.Drawing.Color]::White
$sendButton.FlatStyle = 'Flat'
$sendButton.Add_Click({
    $notifTitle = $titleInput.Text
    $notifMessage = $messageInput.Text
    
    if ([string]::IsNullOrWhiteSpace($notifMessage)) {
        [System.Windows.Forms.MessageBox]::Show('Message cannot be empty!', 'Error', 'OK', 'Error')
        return
    }
    
    if ([string]::IsNullOrWhiteSpace($notifTitle)) {
        $notifTitle = 'Notification'
    }
    
    try {
        [Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
        [Windows.UI.Notifications.ToastNotification, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
        [Windows.Data.Xml.Dom.XmlDocument, System.Xml.XmlDocument, ContentType = WindowsRuntime] | Out-Null

        $APP_ID = 'GoToastify'
        $template = @"
<toast>
    <visual>
        <binding template="ToastText02">
            <text id="1">$notifTitle</text>
            <text id="2">$notifMessage</text>
        </binding>
    </visual>
</toast>
"@

        $xml = New-Object Windows.Data.Xml.Dom.XmlDocument
        $xml.LoadXml($template)
        $toastNotif = New-Object Windows.UI.Notifications.ToastNotification $xml
        [Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($APP_ID).Show($toastNotif)
        
        [System.Windows.Forms.MessageBox]::Show('Toast notification sent successfully!', 'Success', 'OK', 'Information')
    } catch {
        [System.Windows.Forms.MessageBox]::Show("Failed to send notification: $_", 'Error', 'OK', 'Error')
    }
})
$form.Controls.Add($sendButton)

# Close Button
$closeButton = New-Object System.Windows.Forms.Button
$closeButton.Location = New-Object System.Drawing.Point(10, 480)
$closeButton.Size = New-Object System.Drawing.Size(515, 30)
$closeButton.Text = 'Close'
$closeButton.Add_Click({ $form.Close() })
$form.Controls.Add($closeButton)

[void]$form.ShowDialog()
`, cli.VERSION, helpContent)

	cmd := exec.Command("powershell.exe", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", psScript)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error launching GUI: %v\n", err)
	}
}
