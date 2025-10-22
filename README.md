# 🖥️ Helix Bridge CLI

**Command-line interface for Helix Bridge with real JWT authentication and enterprise features**

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Release](https://img.shields.io/badge/Release-bridge--auth--v0.1.0-green.svg)](https://github.com/helixprojectai-code/helix-bridgectl/releases)

## 🚀 Overview

Helix Bridge CLI (`helix-bridgectl`) is a powerful command-line tool for interacting with the Helix Bridge API. It provides **real JWT authentication**, secure token management, and a comprehensive set of commands for managing your Helix resources.

### ✨ Key Features

- 🔐 **Real JWT Authentication** - Enterprise-grade cryptographic token authentication
- 💾 **Secure Token Caching** - Automatic JWT token storage with proper file permissions
- 🆔 **Identity Management** - View authenticated identity with token details and expiration
- 🌐 **API Connectivity** - Full connectivity testing and health checks
- 🏢 **Multi-tenant Support** - Complete tenant isolation and management
- 📊 **Rich Output** - Beautiful, informative command output with color support
- 🐳 **Docker Ready** - Container-friendly with environment variable support

## 📦 Installation

### Binary Download

```bash
# Download latest release
curl -L https://github.com/helixprojectai-code/helix-bridgectl/releases/latest/download/helix-bridgectl-linux-amd64 -o helix-bridgectl
chmod +x helix-bridgectl
sudo mv helix-bridgectl /usr/local/bin/
From Source
bash
git clone https://github.com/helixprojectai-code/helix-bridgectl.git
cd helix-bridgectl
make build
sudo make install
Docker
bash
docker run --rm -v ~/.helix:/root/.helix \
  helixprojectai-code/helix-bridgectl:latest \
  login --base-url "https://bridge.example.com" \
  --client-id "your-client" --client-secret "your-secret"
🏁 Quick Start
1. Test Connectivity
bash
helix-bridgectl ping --base-url "http://localhost:3000"
Output:

text
ping: OK
2. Authenticate
bash
helix-bridgectl login \
  --base-url "http://localhost:3000" \
  --client-id "demo-client" \
  --client-secret "super-secret" \
  --tenant-id "tenant-001"
Output:

text
Authenticating client: demo-client
🔐 CLIENT: Starting REAL authentication for client: demo-client
🚀 REAL AUTH: Calling JWT endpoint for client: demo-client
✅ REAL AUTH: Successfully obtained JWT token (length: 273)
✅ CLIENT: Authentication successful, token length: 273
login: OK (client=demo-client) — REAL JWT token cached at ~/.helix/token
3. Verify Identity
bash
helix-bridgectl whoami --base-url "http://localhost:3000"
Output:

text
🔑 Using cached JWT token (length: 273)...
🔐 Authenticated Identity (REAL JWT):
   Subject: demo-client
   Tenant:  tenant-001
   Scopes:  bridge:read
   Token:   eyJhbGci…
   Expires: 2025-10-21 23:05:49 (30m0s)
   Status:  ✅ Authenticated with REAL JWT
📋 Command Reference
Core Commands
Command	Description	Example
ping	Test connectivity to Bridge API	helix-bridgectl ping --base-url URL
login	Authenticate and cache JWT token	helix-bridgectl login --client-id ID --client-secret SECRET
whoami	Show authenticated identity	helix-bridgectl whoami
Global Flags
Flag	Description	Default
--base-url	Bridge API base URL	http://127.0.0.1:3000
--timeout	Request timeout duration	5s
--verbose	Enable verbose output	false
🔧 Configuration
Environment Variables
bash
# Set in your shell profile
export HELIX_BASE_URL="https://bridge.production.example.com"
export HELIX_CLIENT_ID="production-client"
export HELIX_CLIENT_SECRET="super-secret-production-key"
export HELIX_TENANT_ID="production-tenant"
export HELIX_TIMEOUT="30s"
Configuration File
Create ~/.helix/config.yaml:

yaml
base_url: "https://bridge.production.example.com"
client_id: "production-client"
tenant_id: "production-tenant"
timeout: "30s"
default_scope:
  - "bridge:read"
  - "bridge:write"
Token Storage
Location: ~/.helix/token

Permissions: 0600 (user read/write only)

Format: Raw JWT token string (273 characters)

Auto-refresh: Manual re-authentication required

🏗️ Architecture
Authentication Flow








Security Features
JWT Tokens: 273-character cryptographically signed tokens

Secure Storage: File permissions 0600 prevent unauthorized access

Token Expiration: 30-minute TTL with clear expiration display

No Secrets in Logs: Only token previews shown in output

TLS Support: Full HTTPS support for production environments

🚀 Advanced Usage
Scripting Integration
bash
#!/bin/bash

# Check authentication
if ! helix-bridgectl whoami >/dev/null 2>&1; then
    echo "Not authenticated. Logging in..."
    helix-bridgectl login \
        --client-id "$HELIX_CLIENT_ID" \
        --client-secret "$HELIX_CLIENT_SECRET" \
        --tenant-id "$HELIX_TENANT_ID"
fi

# Use in scripts
IDENTITY=$(helix-bridgectl whoami --output json)
SUBJECT=$(echo "$IDENTITY" | jq -r '.subject')
echo "Running as: $SUBJECT"
Docker Integration
dockerfile
FROM alpine:latest

# Install helix-bridgectl
ADD https://github.com/helixprojectai-code/helix-bridgectl/releases/latest/download/helix-bridgectl-linux-amd64 /usr/local/bin/helix-bridgectl
RUN chmod +x /usr/local/bin/helix-bridgectl

# Set up authentication
ENV HELIX_BASE_URL="https://bridge.example.com"
ENV HELIX_CLIENT_ID="container-client"

# Use in entrypoint
CMD ["helix-bridgectl", "whoami"]
CI/CD Integration
yaml
# GitHub Actions example
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Authenticate with Helix
        run: |
          curl -L https://github.com/helixprojectai-code/helix-bridgectl/releases/latest/download/helix-bridgectl-linux-amd64 -o helix-bridgectl
          chmod +x helix-bridgectl
          ./helix-bridgectl login \
            --base-url "${{ secrets.HELIX_BASE_URL }}" \
            --client-id "${{ secrets.HELIX_CLIENT_ID }}" \
            --client-secret "${{ secrets.HELIX_CLIENT_SECRET }}"
🔒 Security Best Practices
Production Deployment
bash
# Use HTTPS in production
export HELIX_BASE_URL="https://bridge.production.example.com"

# Use dedicated service accounts
export HELIX_CLIENT_ID="deployment-bot"
export HELIX_CLIENT_SECRET=$(vault read -field=secret secret/helix/deployment)

# Set appropriate timeouts
export HELIX_TIMEOUT="30s"
Token Management
bash
# Check token status
helix-bridgectl whoami

# Clear tokens when needed
rm ~/.helix/token

# Verify file permissions
ls -la ~/.helix/token
# Should show: -rw------- (600)
🧪 Testing
Local Development
bash
# Build from source
make build

# Test against local bridge
./helix-bridgectl ping --base-url "http://localhost:3000"

# Run all tests
make test
Integration Testing
bash
# Test full authentication flow
./test-auth-flow.sh

# Example test script
#!/bin/bash
set -e

echo "🧪 Testing authentication flow..."
./helix-bridgectl login \
  --base-url "http://localhost:3000" \
  --client-id "demo-client" \
  --client-secret "super-secret"

./helix-bridgectl whoami --base-url "http://localhost:3000"
echo "✅ All tests passed!"
🐛 Troubleshooting
Common Issues
Connection Refused

bash
# Check if bridge is running
helix-bridgectl ping --base-url "http://localhost:3000"

# Check network connectivity
curl -v http://localhost:3000/health
Authentication Failures

bash
# Verify credentials
echo "Client: $HELIX_CLIENT_ID"
echo "Base URL: $HELIX_BASE_URL"

# Clear cached token and retry
rm ~/.helix/token
helix-bridgectl login --client-id "id" --client-secret "secret"
Token Issues

bash
# Check token file
ls -la ~/.helix/
cat ~/.helix/token | head -c 20

# Validate token structure
token=$(cat ~/.helix/token)
echo "Token length: ${#token}"
Debug Mode
bash
# Enable verbose output
helix-bridgectl whoami --verbose

# Or set environment variable
export HELIX_DEBUG="true"
helix-bridgectl login --client-id "id" --client-secret "secret"
📚 Examples
Basic Workflow
bash
#!/bin/bash

# Test connection
if ! helix-bridgectl ping; then
    echo "❌ Cannot connect to Bridge API"
    exit 1
fi

# Authenticate if needed
if ! helix-bridgectl whoami >/dev/null 2>&1; then
    helix-bridgectl login \
        --client-id "$HELIX_CLIENT_ID" \
        --client-secret "$HELIX_CLIENT_SECRET"
fi

# Show identity
echo "🔐 Current identity:"
helix-bridgectl whoami
Production Script
bash
#!/bin/bash

set -euo pipefail

# Configuration
BRIDGE_URL="${HELIX_BASE_URL:-https://bridge.production.example.com}"
TIMEOUT="30s"

echo "🚀 Starting deployment..."

# Verify authentication
if ! helix-bridgectl whoami --base-url "$BRIDGE_URL" --timeout "$TIMEOUT" >/dev/null; then
    echo "🔐 Authenticating..."
    helix-bridgectl login \
        --base-url "$BRIDGE_URL" \
        --timeout "$TIMEOUT" \
        --client-id "$HELIX_CLIENT_ID" \
        --client-secret "$HELIX_CLIENT_SECRET"
fi

# Get identity for audit
IDENTITY=$(helix-bridgectl whoami --base-url "$BRIDGE_URL" --output json)
echo "👤 Deploying as: $(echo "$IDENTITY" | jq -r '.subject')"

# Continue with deployment...
🤝 Contributing
We welcome contributions! Please see our Contributing Guide for details.

Development Setup
bash
# Fork and clone
git clone https://github.com/your-username/helix-bridgectl.git
cd helix-bridgectl

# Build
make build

# Test
make test

# Install locally
make install
Release Process
bash
# Create new release
make release VERSION="v1.2.3"

# Build for all platforms
make release-all
📄 License
Licensed under the Apache License 2.0 - see the LICENSE file for details.

🆘 Support
📚 Documentation: GitHub Wiki

🐛 Bug Reports: GitHub Issues

💬 Discussions: GitHub Discussions

📧 Email: [email protected]

🙏 Acknowledgments
Built with ❤️ by the Helix AI Team

Uses helix-sdk-go for core authentication

Command-line framework powered by Cobra

<div align="center">
Helix Bridge CLI • Documentation • Changelog • Download Latest

</div>
