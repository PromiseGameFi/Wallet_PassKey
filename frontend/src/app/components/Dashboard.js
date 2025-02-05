'use client'; // Mark as a Client Component

import { useEffect, useState } from 'react';
import { fetchWalletAddresses } from '../api/auth';

export default function Dashboard() {
	const [addresses, setAddresses] = useState([]);

	useEffect(() => {
		fetchWalletAddresses()
			.then((data) => setAddresses(data))
			.catch((err) => console.error(err));
	}, []);

	return (
		<div>
			<h2>Wallet Addresses</h2>
			<ul>
				{addresses.map((address, index) => (
					<li key={index}>{address}</li>
				))}
			</ul>
		</div>
	);
}