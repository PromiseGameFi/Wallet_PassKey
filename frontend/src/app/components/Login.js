'use client'; // Mark as a Client Component

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { loginWithPasskey } from '../api/auth';

export default function Login() {
	const router = useRouter();
	const [error, setError] = useState('');

	const handleLogin = async () => {
		try {
			await loginWithPasskey();
			router.push('/dashboard');
		} catch (err) {
			setError('Login failed. Please try again.');
		}
	};

	return (
		<div>
			<button onClick={handleLogin}>Log in with Passkey</button>
			{error && <p style={{ color: 'red' }}>{error}</p>}
		</div>
	);
}