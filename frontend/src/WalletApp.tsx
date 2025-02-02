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
    <div
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