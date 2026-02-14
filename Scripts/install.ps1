#Requires -Version 3.0
<#
.SYNOPSIS
    BanglaCode Installer for Windows
.DESCRIPTION
    Downloads and installs the latest version of BanglaCode programming language.
    Automatically detects system architecture and configures PATH.
.EXAMPLE
    # Run directly from web:
    irm https://raw.githubusercontent.com/nexoral/BanglaCode/main/Scripts/install.ps1 | iex
    
    # Or download and run:
    .\install.ps1
.NOTES
    Requires PowerShell 3.0 or later (Windows 7+, Windows Server 2008 R2+)
    For Windows PowerShell 5.1 and PowerShell Core 6+/7+
#>

# Configuration
$REPO = "nexoral/BanglaCode"
$INSTALL_DIR = "$env:LOCALAPPDATA\BanglaCode"
$ErrorActionPreference = "Stop"

# Enable TLS 1.2 for older PowerShell versions (required for GitHub API)
[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

function Write-ColorOutput {
    param(
        [string]$Message,
        [string]$Color = "White",
        [switch]$NoNewline
    )
    
    # Fallback for older terminals without color support
    try {
        if ($NoNewline) {
            Write-Host $Message -ForegroundColor $Color -NoNewline
        } else {
            Write-Host $Message -ForegroundColor $Color
        }
    } catch {
        if ($NoNewline) {
            Write-Host $Message -NoNewline
        } else {
            Write-Host $Message
        }
    }
}

function Test-Administrator {
    $currentUser = [Security.Principal.WindowsIdentity]::GetCurrent()
    $principal = New-Object Security.Principal.WindowsPrincipal($currentUser)
    return $principal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
}

function Get-SystemArchitecture {
    # Check for ARM64 first
    if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") {
        return "arm64"
    }
    
    # Check if running on 64-bit Windows
    if ([Environment]::Is64BitOperatingSystem) {
        return "amd64"
    }
    
    return "386"
}

function Test-CommandExists {
    param([string]$Command)
    $oldPreference = $ErrorActionPreference
    $ErrorActionPreference = "Stop"
    try {
        if (Get-Command $Command -ErrorAction Stop) { return $true }
    } catch {
        return $false
    } finally {
        $ErrorActionPreference = $oldPreference
    }
}

function Get-LatestVersion {
    param([string]$Repository)
    
    $apiUrl = "https://api.github.com/repos/$Repository/releases/latest"
    
    try {
        # Try Invoke-RestMethod first (PowerShell 3.0+)
        $release = Invoke-RestMethod -Uri $apiUrl -UseBasicParsing
        return $release.tag_name -replace '^v', ''
    } catch {
        # Fallback: try with WebClient for very old systems
        try {
            $webClient = New-Object System.Net.WebClient
            $webClient.Headers.Add("User-Agent", "PowerShell")
            $json = $webClient.DownloadString($apiUrl)
            if ($json -match '"tag_name"\s*:\s*"v?([^"]+)"') {
                return $matches[1]
            }
        } catch {
            return $null
        }
    }
    return $null
}

function Download-File {
    param(
        [string]$Url,
        [string]$Destination
    )
    
    try {
        # Try Invoke-WebRequest first (better progress indication)
        if ($PSVersionTable.PSVersion.Major -ge 5) {
            Invoke-WebRequest -Uri $Url -OutFile $Destination -UseBasicParsing
        } else {
            # Fallback for older PowerShell
            $webClient = New-Object System.Net.WebClient
            $webClient.DownloadFile($Url, $Destination)
        }
        return $true
    } catch {
        return $false
    }
}

function Expand-ZipFile {
    param(
        [string]$ZipPath,
        [string]$Destination
    )
    
    try {
        # Try Expand-Archive first (PowerShell 5.0+)
        if ($PSVersionTable.PSVersion.Major -ge 5) {
            Expand-Archive -Path $ZipPath -DestinationPath $Destination -Force
        } else {
            # Fallback for older PowerShell using .NET
            Add-Type -AssemblyName System.IO.Compression.FileSystem
            [System.IO.Compression.ZipFile]::ExtractToDirectory($ZipPath, $Destination)
        }
        return $true
    } catch {
        # Last resort: use Shell.Application COM object
        try {
            $shell = New-Object -ComObject Shell.Application
            $zip = $shell.NameSpace($ZipPath)
            $dest = $shell.NameSpace($Destination)
            $dest.CopyHere($zip.Items(), 16) # 16 = overwrite
            return $true
        } catch {
            return $false
        }
    }
}

function Add-ToPath {
    param([string]$Directory)
    
    $currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
    
    if ($currentPath -notlike "*$Directory*") {
        try {
            [Environment]::SetEnvironmentVariable("Path", "$currentPath;$Directory", "User")
            $env:Path = "$env:Path;$Directory"
            return $true
        } catch {
            return $false
        }
    }
    return $true  # Already in PATH
}

function Show-Banner {
    Write-Host ""
    Write-ColorOutput "  ____                    _        ____          _      " "Cyan"
    Write-ColorOutput " | __ )  __ _ _ __   __ _| | __ _ / ___|___   __| | ___ " "Cyan"
    Write-ColorOutput " |  _ \ / _` | '_ \ / _` | |/ _` | |   / _ \ / _` |/ _ \" "Cyan"
    Write-ColorOutput " | |_) | (_| | | | | (_| | | (_| | |__| (_) | (_| |  __/" "Cyan"
    Write-ColorOutput " |____/ \__,_|_| |_|\__, |_|\__,_|\____\___/ \__,_|\___|" "Cyan"
    Write-ColorOutput "                    |___/                               " "Cyan"
    Write-Host ""
    Write-ColorOutput "        BanglaCode Installer for Windows" "White"
    Write-ColorOutput "        ================================" "DarkGray"
    Write-Host ""
}

# ============== Main Installation Logic ==============

Show-Banner

# Check PowerShell version
$psVersion = $PSVersionTable.PSVersion
Write-ColorOutput "[+] PowerShell Version: $($psVersion.Major).$($psVersion.Minor)" "Gray"

if ($psVersion.Major -lt 3) {
    Write-ColorOutput "[!] Warning: PowerShell 3.0+ recommended. Some features may not work." "Yellow"
}

# Detect architecture
$ARCH = Get-SystemArchitecture
Write-ColorOutput "[+] Detected Architecture: windows-$ARCH" "Green"

# Check if running as admin (informational only)
if (Test-Administrator) {
    Write-ColorOutput "[+] Running as Administrator" "Green"
} else {
    Write-ColorOutput "[*] Running as standard user (recommended)" "Gray"
}

# Check for existing installation
if (Test-CommandExists "banglacode") {
    $existingPath = (Get-Command banglacode -ErrorAction SilentlyContinue).Source
    Write-ColorOutput "[*] Existing installation found: $existingPath" "Yellow"
    Write-ColorOutput "[*] Will upgrade to latest version" "Yellow"
}

# Get latest version
Write-ColorOutput "[*] Fetching latest version..." "White"
$VERSION = Get-LatestVersion -Repository $REPO

if (-not $VERSION) {
    Write-ColorOutput "[X] Failed to get latest version from GitHub" "Red"
    Write-ColorOutput "    Check your internet connection and try again." "Red"
    Write-Host ""
    Write-ColorOutput "Manual download: https://github.com/$REPO/releases" "Yellow"
    exit 1
}

Write-ColorOutput "[+] Latest Version: v$VERSION" "Green"

# Prepare download
$DOWNLOAD_FILE = "banglacode-windows-$ARCH.zip"
$DOWNLOAD_URL = "https://github.com/$REPO/releases/download/v$VERSION/$DOWNLOAD_FILE"
$TMP_FILE = Join-Path $env:TEMP $DOWNLOAD_FILE

Write-ColorOutput "[*] Downloading $DOWNLOAD_FILE..." "White"

# Download file
if (-not (Download-File -Url $DOWNLOAD_URL -Destination $TMP_FILE)) {
    Write-ColorOutput "[X] Download failed!" "Red"
    Write-ColorOutput "    URL: $DOWNLOAD_URL" "Red"
    Write-Host ""
    Write-ColorOutput "Try manual download: https://github.com/$REPO/releases" "Yellow"
    exit 1
}

Write-ColorOutput "[+] Download complete" "Green"

# Create install directory
Write-ColorOutput "[*] Installing to $INSTALL_DIR..." "White"

try {
    if (-not (Test-Path $INSTALL_DIR)) {
        New-Item -ItemType Directory -Force -Path $INSTALL_DIR | Out-Null
    }
} catch {
    Write-ColorOutput "[X] Failed to create install directory" "Red"
    Write-ColorOutput "    Try running as Administrator" "Yellow"
    exit 1
}

# Extract ZIP
if (-not (Expand-ZipFile -ZipPath $TMP_FILE -Destination $INSTALL_DIR)) {
    Write-ColorOutput "[X] Failed to extract archive" "Red"
    Remove-Item $TMP_FILE -Force -ErrorAction SilentlyContinue
    exit 1
}

# Cleanup temp file
Remove-Item $TMP_FILE -Force -ErrorAction SilentlyContinue

Write-ColorOutput "[+] Files extracted" "Green"

# Add to PATH
Write-ColorOutput "[*] Configuring PATH..." "White"

if (Add-ToPath -Directory $INSTALL_DIR) {
    Write-ColorOutput "[+] Added to user PATH" "Green"
} else {
    Write-ColorOutput "[!] Could not add to PATH automatically" "Yellow"
    Write-ColorOutput "    Please add manually: $INSTALL_DIR" "Yellow"
}

# Verify installation
$banglacodePath = Join-Path $INSTALL_DIR "banglacode.exe"
if (Test-Path $banglacodePath) {
    Write-Host ""
    Write-ColorOutput "=============================================" "Green"
    Write-ColorOutput "  BanglaCode installed successfully!" "Green"
    Write-ColorOutput "=============================================" "Green"
    Write-Host ""
    Write-ColorOutput "  Version:  v$VERSION" "White"
    Write-ColorOutput "  Location: $INSTALL_DIR" "White"
    Write-Host ""
    Write-ColorOutput "  IMPORTANT: Restart your terminal to use 'banglacode'" "Yellow"
    Write-Host ""
    Write-ColorOutput "  Quick Start:" "Cyan"
    Write-ColorOutput "    banglacode              - Start interactive REPL" "White"
    Write-ColorOutput "    banglacode hello.bang   - Run a BanglaCode file" "White"
    Write-ColorOutput "    banglacode --help       - Show help" "White"
    Write-Host ""
    Write-ColorOutput "  Documentation: https://github.com/$REPO" "Gray"
    Write-ColorOutput "  VS Code Extension: Search 'BanglaCode' in Extensions" "Gray"
    Write-Host ""
} else {
    Write-ColorOutput "[X] Installation verification failed" "Red"
    Write-ColorOutput "    Expected file not found: $banglacodePath" "Red"
    exit 1
}

# Offer to open new terminal
Write-ColorOutput "Press any key to exit..." "DarkGray"
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
