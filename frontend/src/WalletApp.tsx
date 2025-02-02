import React, { useState, useEffect } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Plus, Key, Wallet } from 'lucide-react';

const WalletApp = () => {
  const [accounts, setAccounts] = useState([]);
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  const handlePasskeyAuth = async () => {
    try {
      if (!isAuthenticated) {
        const cred = await navigator.credentials.get({
          publicKey: {
            challenge: new Uint8Array(32),
            rpId: window.location.hostname,
            userVerification: 'required',
          }
        });
        
        // Verify with backend
        setIsAuthenticated(true);
      }
    } catch (error) {
      console.error('Authentication failed:', error);
    }
  };

  const createNewAccount = async () => {
    try {
      const response = await fetch('/api/wallet/account', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        }
      });
      
      const newAccount = await response.json();
      setAccounts([...accounts, newAccount]);
    } catch (error) {
      console.error('Failed to create account:', error);
    }
  };

  return (
    <div className="container mx-auto p-4">
      <Card className="mb-6">
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Wallet className="h-6 w-6" />
            HD Wallet with Passkey
          </CardTitle>
        </CardHeader>
        <CardContent>
          {!isAuthenticated ? (
            <div className="text-center">
              <Button 
                onClick={handlePasskeyAuth}
                className="flex items-center gap-2"
              >
                <Key className="h-4 w-4" />
                Authenticate with Passkey
              </Button>
            </div>
          ) : (
            <div className="space-y-4">
              <div className="flex justify-between items-center">
                <h3 className="text-lg font-medium">Your Accounts</h3>
                <Button
                  onClick={createNewAccount}
                  className="flex items-center gap-2"
                >
                  <Plus className="h-4 w-4" />
                  Add Account
                </Button>
              </div>
              
              <div className="space-y-2">
                {accounts.map((account, index) => (
                  <Card key={account.index} className="p-4">
                    <div className="flex justify-between items-center">
                      <div>
                        <p className="font-medium">Account {index + 1}</p>
                        <p className="text-sm text-gray-500 break-all">
                          {account.address}
                        </p>
                      </div>
                    </div>
                  </Card>
                ))}
              </div>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
};

export default WalletApp;