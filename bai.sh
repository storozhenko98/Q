#!/bin/bash

# Set variables
GO_FILE="main.go"
BINARY_NAME="Q"
INSTALL_DIR="/usr/local/bin"

# Ensure the script is run from the directory containing the Go file
if [ ! -f "$GO_FILE" ]; then
    echo "Error: $GO_FILE not found in the current directory."
    exit 1
fi

# Compile the Go program
echo "Compiling $GO_FILE..."
go build -o "$BINARY_NAME" "$GO_FILE"

if [ $? -ne 0 ]; then
    echo "Error: Compilation failed."
    exit 1
fi

# Move the binary to the installation directory
echo "Installing $BINARY_NAME to $INSTALL_DIR..."
sudo mv "$BINARY_NAME" "$INSTALL_DIR"

if [ $? -ne 0 ]; then
    echo "Error: Failed to move $BINARY_NAME to $INSTALL_DIR."
    exit 1
fi

# Set appropriate permissions
sudo chmod 755 "$INSTALL_DIR/$BINARY_NAME"

echo "Installation complete. You can now use '$BINARY_NAME' from anywhere in the terminal."