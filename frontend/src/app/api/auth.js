export const loginWithPasskey = async () => {
	const response = await fetch('http://localhost:8080/auth/login/begin', {
		method: 'POST',
	});
	const options = await response.json();

	const credential = await navigator.credentials.get({
		publicKey: options,
	});

	await fetch('http://localhost:8080/auth/login/finish', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(credential),
	});
};

export const fetchWalletAddresses = async () => {
	const response = await fetch('http://localhost:8080/wallet/addresses');
	return response.json();
};

export const sendTransaction = async (transaction) => {
	const response = await fetch('http://localhost:8080/wallet/transactions', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(transaction),
	});
	return response.json();
};