# Go-scripts
  Scripts Go used in my daily life

# 1- Windows Event Log Exporter (Go)
This project is a small utility written in **Go** that collects the latest Windows Event Logs and exports them as JSON files.  
It is designed as a facilitator for less technical users who want a simple way to check recent system/application/security/setup events without manually browsing the Windows Event Viewer.
## Features
- Runs PowerShell commands directly from Go using `exec.Command`.
- Exports logs from:
  - **Application**
  - **Security**
  - **System**
  - **Setup**
- Generates JSON files with the following fields:
  - `TimeCreated`
  - `Id` / `EventId`
  - `ProviderName`
  - `Message`
- Admin permissions required
