# Go-scripts
  Scripts Go used in my daily life
  
---

## 1. Windows Event Log Exporter

This project is a utility written in **Go** that collects the latest Windows Event Logs and exports them as JSON files.  
It is designed as a facilitator for less technical users who want a simple way to check recent **Application**, **System**, **Security**, or **Setup** events without manually browsing the Windows Event Viewer.

### Features
- Executes PowerShell commands directly from Go using `exec.Command`.
- Exports logs from:
  - Application  
  - Security  
  - System  
  - Setup  
- Generates JSON files with the following fields:
  - `TimeCreated`  
  - `Id` / `EventId`  
  - `ProviderName`  
  - `Message`  

### Requirements
- Windows OS  
- PowerShell available in the system  
- Run as **Administrator** to access all logs  

### Output Example
```json
[
  {
    "TimeCreated": "2025-09-25T02:45:12Z",
    "Id": 7036,
    "ProviderName": "Service Control Manager",
    "Message": "The Windows Update service entered the running state."
  }
]
