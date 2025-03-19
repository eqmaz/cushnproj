#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
VERSION_FILE="${SCRIPT_DIR}/version.json"

# Read current version info using jq
MAJOR=$(jq -r '.major' $VERSION_FILE)
MINOR=$(jq -r '.minor' $VERSION_FILE)
PATCH=$(jq -r '.patch' $VERSION_FILE)
BUILD=$(jq -r '.build' $VERSION_FILE)

# Increment build number
NEW_BUILD=$((BUILD + 1))

# Update the version.json file
jq --argjson newBuild "$NEW_BUILD" '.build = $newBuild' $VERSION_FILE > tmp.json && mv tmp.json $VERSION_FILE

# Output the full version string
echo "${MAJOR}.${MINOR}.${PATCH}.${NEW_BUILD}"
