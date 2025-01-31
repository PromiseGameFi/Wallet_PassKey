# HD Wallet with Passkey Authentication - Documentation

## Overview
This is a minimal implementation of a Hierarchical Deterministic (HD) wallet with passkey authentication. The project consists of a Rust backend handling wallet operations and passkey authentication, and a React frontend providing the user interface.

## Project Structure
```
hd-wallet-passkey/
├── Cargo.toml
├── src/
│   └── lib.rs          # Backend implementation
└── frontend/
    ├── package.json
    └── src/
        └── WalletApp.tsx   # Frontend implementation
```

## Prerequisites
- Rust (1.70 or later)
- Node.js (18 or later)
- npm
- A WebAuthn-capable browser

## Backend Setup

1. Create a new Rust project:
```bash
cargo new hd-wallet-passkey
cd hd-wallet-passkey
```

2. Add dependencies to `Cargo.toml`:
```toml
[package]
name = "hd-wallet-passkey"
version = "0.1.0"
edition = "2021"

[dependencies]
tiny-bip39 = "1.0"
bitcoincore-lib = "0.27"
hex = "0.4"
serde = { version = "1.0", features = ["derive"] }
serde_json = "1.0"
tokio = { version = "1.0", features = ["full"] }
warp = "0.3"
webauthn-rs = "0.3"
thiserror = "1.0"
tracing = "0.1"
ring = "0.16"
```

3. Copy the provided backend code into `src/lib.rs`

## Frontend Setup

1. Create a new React project with Vite:
```bash
npm create vite@latest frontend -- --template react-ts
cd frontend
```

2. Install dependencies:
```bash
npm install @/components/ui/card @/components/ui/button lucide-react
npm install @radix-ui/react-icons
```

3. Install shadcn/ui:
```bash
npx shadcn-ui@latest init
```

4. Copy the provided frontend code into `src/WalletApp.tsx`

## Running the Application

### Backend
1. Start the Rust server:
```bash
cd hd-wallet-passkey
cargo run
```

The server will start on `http://localhost:3030`

### Frontend
1. Start the development server:
```bash
cd frontend
npm run dev
```

The application will be available at `http://localhost:5173`

## Features

### Backend Features
1. HD Wallet Implementation
   - BIP39 mnemonic generation
   - BIP32 hierarchical deterministic key derivation
   - Multiple account support
   - Secure key storage

2. Passkey Authentication
   - WebAuthn standard implementation
   - Secure user verification
   - Challenge-response authentication

3. API Endpoints
   - POST `/api/wallet/create` - Create new wallet
   - POST `/api/wallet/account` - Add new account
   - GET `/api/wallet/accounts` - List accounts

### Frontend Features
1. User Interface
   - Clean, modern design using shadcn/ui
   - Responsive layout
   - Dark/light mode support

2. Wallet Management
   - Account creation
   - Account list view
   - Address display

3. Authentication
   - Passkey registration
   - Passkey authentication
   - Session management

## Usage Guide

1. First-time Setup:
   - Open the application in your browser
   - Click "Authenticate with Passkey"
   - Follow browser prompts to create a new passkey
   - Your wallet will be created automatically

2. Creating Additional Accounts:
   - Click "Add Account" button
   - New account will be derived from master seed
   - Account address will be displayed

3. Managing Accounts:
   - View all accounts in the main interface
   - Each account shows its address
   - Copy addresses using the copy button

## Security Considerations

1. Private Keys
   - Never exposed to frontend
   - Securely derived using BIP32
   - Stored encrypted in backend

2. Authentication
   - WebAuthn standard ensures strong authentication
   - Biometric/hardware security where available
   - No password storage needed

3. Best Practices
   - Always use HTTPS in production
   - Keep your system and browser up to date
   - Don't share passkey credentials

## Error Handling

1. Backend Errors:
   - `MnemonicError`: Issues with seed phrase generation
   - `DerivationError`: Problems with key derivation
   - `PasskeyError`: Authentication failures

2. Frontend Errors:
   - Authentication failures
   - Network connectivity issues
   - Account creation failures

## Limitations
1. Current Implementation:
   - Single user per instance
   - No transaction support
   - Limited to one passkey per wallet

2. Security:
   - Development server not suitable for production
   - No rate limiting implemented
   - Basic error handling

## Future Improvements
1. Potential Enhancements:
   - Multi-user support
   - Transaction handling
   - Multiple passkey support
   - Enhanced error handling
   - Rate limiting
   - Activity logging

## Troubleshooting

1. Backend Issues:
   - Ensure Rust and dependencies are up to date
   - Check server logs for errors
   - Verify port 3030 is available

2. Frontend Issues:
   - Clear browser cache if authentication fails
   - Ensure WebAuthn is supported by your browser
   - Check console for error messages

## Support
For issues and feature requests, please create an issue in the repository with:
- Detailed description of the problem
- Steps to reproduce
- Expected vs actual behavior
- Environment details

## License
This project is licensed under the MIT License.
