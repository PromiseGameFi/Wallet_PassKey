# Passkey HD Crypto Wallet

## Overview
A secure, hierarchical deterministic cryptocurrency wallet using passkey authentication instead of traditional seed phrases.

## Key Features
- Passkey-based authentication
- Hierarchical Deterministic (HD) wallet structure
- Rust implementation
- Cross-platform UI

## Architecture
- Backend: Rust
- Frontend: Web-based UI (Tauri/React)
- Authentication: WebAuthn Passkey 
- Key Derivation: BIP32/BIP44 standards

## Prerequisites
- Rust (latest stable version)
- Node.js
- Cargo
- WebAuthn-compatible browser

## Installation
```bash
git clone https://github.com/yourusername/passkey-hd-wallet.git
cd passkey-hd-wallet
cargo build
npm install
```



## Security Considerations
- Passkey stored securely using platform authenticators
- No seed phrase exposure
- Hardware-backed key storage
