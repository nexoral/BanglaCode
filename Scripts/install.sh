#!/bin/bash
set -e

# BanglaCode Installer Script
# Usage: curl -fsSL https://raw.githubusercontent.com/AKN414-IND/BanglaCode/main/Scripts/install.sh | bash

REPO="AKN414-IND/BanglaCode"
INSTALL_DIR="/usr/local/bin"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 BanglaCode Installer${NC}"
echo "========================"

# Detect OS and Architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64|amd64) ARCH="amd64" ;;
    i386|i686) ARCH="386" ;;
    aarch64|arm64) ARCH="arm64" ;;
    armv7*|armv6*) ARCH="arm" ;;
    *) echo -e "${RED}❌ Unsupported architecture: $ARCH${NC}"; exit 1 ;;
esac

case "$OS" in
    linux) OS="linux" ;;
    darwin) OS="macos" ;;
    freebsd) OS="freebsd" ;;
    openbsd) OS="openbsd" ;;
    netbsd) OS="netbsd" ;;
    *) echo -e "${RED}❌ Unsupported OS: $OS${NC}"; exit 1 ;;
esac

echo -e "${GREEN}✓${NC} Detected: $OS-$ARCH"

# Get latest version
VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | sed -E 's/.*"v([^"]+)".*/\1/')
if [ -z "$VERSION" ]; then
    echo -e "${RED}❌ Failed to get latest version${NC}"
    exit 1
fi
echo -e "${GREEN}✓${NC} Latest version: v$VERSION"

# Determine download file
if [ "$OS" = "linux" ]; then
    # Try to detect package manager for native package
    if command -v apt-get &> /dev/null && [ "$ARCH" = "amd64" ]; then
        DOWNLOAD_FILE="banglacode-linux-amd64.deb"
        INSTALL_CMD="sudo dpkg -i"
    elif command -v apt-get &> /dev/null && [ "$ARCH" = "arm64" ]; then
        DOWNLOAD_FILE="banglacode-linux-arm64.deb"
        INSTALL_CMD="sudo dpkg -i"
    elif command -v dnf &> /dev/null && [ "$ARCH" = "amd64" ]; then
        DOWNLOAD_FILE="banglacode-linux-amd64.rpm"
        INSTALL_CMD="sudo dnf install -y"
    elif command -v dnf &> /dev/null && [ "$ARCH" = "arm64" ]; then
        DOWNLOAD_FILE="banglacode-linux-arm64.rpm"
        INSTALL_CMD="sudo dnf install -y"
    elif command -v yum &> /dev/null && [ "$ARCH" = "amd64" ]; then
        DOWNLOAD_FILE="banglacode-linux-amd64.rpm"
        INSTALL_CMD="sudo yum install -y"
    elif command -v yum &> /dev/null && [ "$ARCH" = "arm64" ]; then
        DOWNLOAD_FILE="banglacode-linux-arm64.rpm"
        INSTALL_CMD="sudo yum install -y"
    else
        DOWNLOAD_FILE="banglacode-${OS}-${ARCH}.tar.gz"
        INSTALL_CMD="tar"
    fi
else
    DOWNLOAD_FILE="banglacode-${OS}-${ARCH}.tar.gz"
    INSTALL_CMD="tar"
fi

DOWNLOAD_URL="https://github.com/$REPO/releases/download/v$VERSION/$DOWNLOAD_FILE"

echo -e "${YELLOW}⬇${NC}  Downloading $DOWNLOAD_FILE..."

TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

if ! curl -fsSL -o "$DOWNLOAD_FILE" "$DOWNLOAD_URL"; then
    echo -e "${RED}❌ Download failed${NC}"
    rm -rf "$TMP_DIR"
    exit 1
fi

echo -e "${YELLOW}📦${NC} Installing..."

if [ "$INSTALL_CMD" = "tar" ]; then
    tar -xzf "$DOWNLOAD_FILE"
    sudo mv banglacode "$INSTALL_DIR/"
    sudo chmod +x "$INSTALL_DIR/banglacode"
else
    $INSTALL_CMD "$DOWNLOAD_FILE"
fi

# Cleanup
cd /
rm -rf "$TMP_DIR"

# Verify installation
if command -v banglacode &> /dev/null; then
    echo -e "${GREEN}✅ BanglaCode installed successfully!${NC}"
    echo ""
    echo "Run 'banglacode' to start the REPL"
    echo "Run 'banglacode <file.bang>' to execute a file"
    echo ""
    banglacode --version 2>/dev/null || echo "Version: v$VERSION"
else
    echo -e "${RED}❌ Installation verification failed${NC}"
    exit 1
fi
