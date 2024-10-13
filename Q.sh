#!/bin/bash

# Ensure HOME is set
if [ -z "$HOME" ]; then
    HOME=$(getent passwd $(id -u) | cut -d: -f6)
    if [ -z "$HOME" ]; then
        echo "Error: Unable to determine home directory."
        exit 1
    fi
fi

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check if running on macOS
if [[ "$(uname)" != "Darwin" ]]; then
    echo "This script is intended for macOS only."
    exit 1
fi

# Check if running on ARM-based Mac
if [[ "$(uname -m)" != "arm64" ]]; then
    echo "This script is intended for ARM-based Macs only."
    exit 1
fi

# Check if Go is installed
if ! command_exists go; then
    echo "Go is not installed. Installing Go..."
    
    GO_VERSION="1.23.2"  # Version from go.mod
    GO_PACKAGE="go${GO_VERSION}.darwin-arm64.pkg"
    
    curl -O "https://golang.org/dl/${GO_PACKAGE}"
    sudo installer -pkg "${GO_PACKAGE}" -target /
    rm "${GO_PACKAGE}"
    
    # Set up Go environment variables
    echo 'export PATH=$PATH:/usr/local/go/bin' >> "${HOME}/.zshrc"
    echo 'export GOPATH=$HOME/go' >> "${HOME}/.zshrc"
    echo 'export PATH=$PATH:$GOPATH/bin' >> "${HOME}/.zshrc"
    source "${HOME}/.zshrc"
    
    echo "Go has been installed successfully."
else
    echo "Go is already installed."
fi

# Create a temporary directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR" || exit 1

# Download the repository as a ZIP file
REPO_URL="https://github.com/storozhenko98/Q"
ZIP_URL="${REPO_URL}/archive/refs/heads/main.zip"
ZIP_FILE="Q-main.zip"

echo "Downloading project..."
curl -L -o "$ZIP_FILE" "$ZIP_URL"

# Unzip the downloaded file
unzip "$ZIP_FILE"
cd "Q-main" || exit 1

# Compile the project
go build -o Q

# Move the binary to /usr/local/bin
echo "Installing Q to /usr/local/bin..."
sudo mv Q /usr/local/bin

if [ $? -ne 0 ]; then
    echo "Error: Failed to move Q to /usr/local/bin."
    cd "$HOME" || exit 1
    rm -rf "$TEMP_DIR"
    exit 1
fi

# Set appropriate permissions
sudo chmod 755 /usr/local/bin/Q

# Clean up: Remove the temporary directory
cd "$HOME" || exit 1
rm -rf "$TEMP_DIR"

# Show completion message
echo "Project has been successfully compiled and installed as 'Q'."
echo "You can now use 'Q' from anywhere in the terminal."
echo "For help, use 'Q --help' or 'Q -help' or 'Q -h' or 'Q --h'"
