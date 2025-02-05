'use client'; // Mark as a Client Component

import { useState } from 'react';
import { sendTransaction } from '../api/auth';

export default function Transaction() {
	const [amount, setAmount] = useState('');
	const [recipient, setRecipient] = useState('');

	const handleSend = async () => {
		try {
			await sendTransaction({ amount, recipient });
			alert('Transaction sent successfully!');
		} catch (err) {
			alert('Transaction failed. Please try again.');
		}
	};

	return (
		<div>
			<h2>Send Transaction</h2>
			<input
				type="text"
				placeholder="Recipient Address"
				value={recipient}
				onChange={(e) => setRecipient(e.target.value)}
			/>
			<input
				type="number"
				placeholder="Amount"
				value={amount}
				onChange={(e) => setAmount(e.target.value)}
			/>
			<button onClick={handleSend}>Send</button>
		</div>
	);
}