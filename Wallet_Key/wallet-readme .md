# Passkey Hierarchical Deterministic Wallet

## Overview
A secure cryptocurrency wallet using passkey authentication and hierarchical deterministic address generation.

## Key Features
- WebAuthn Passkey Authentication
- Hierarchical Deterministic (HD) Wallet
- Secure Address Derivation
- Cross-Platform Support

## Architecture
- Backend: Go
- Frontend: React
- Authentication: WebAuthn
- Key Derivation: BIP32/BIP44 Compatible

## Security Principles
- No seed phrase exposure
- Hardware-backed key storage
- Cryptographically secure key generation
- Platform authenticator integration

## Wallet Lifecycle
1. Wallet Creation
   - Generate master key from passkey
   - Create initial derivation path
   - Securely store master key

2. Address Management
   - Derive child addresses 
   - Support multiple cryptocurrency networks
   - Manage address lifecycle

3. Authentication
   - Passkey-based login
   - Biometric/Hardware key support
   - Secure key retrieval

## To run the project:
1. Install Go and Node.js
2. Backend setup:

# In backend directory
 - go mod init passkey-wallet
 - go mod tidy
 - go run main.go
