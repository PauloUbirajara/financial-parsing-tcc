import type { Actions } from './$types';
import { redirect } from '@sveltejs/kit';

import constants from '../../../../constants'

async function getLoginBody(e): object {
	const formData = await e.request.formData();
	const loginBody = {
		email: formData.get('email'),
		password: formData.get('password'),
	}

	return loginBody
}

type LoginResult = {
	title: string
	success: boolean
	message?: string
}

export const actions: Actions = {
	default: async (e) => {
		let result: LoginResult = {
			title: "Acesso de usuário",
			success: false
		}

		// Get content from form
		const loginBody = await getLoginBody(e)

		// Send contents to backend for login
		const response = await fetch(constants.API_LOGIN_URL, {
			headers: {
				'Content-Type': 'application/json'
			},
			method: 'POST',
			body: JSON.stringify(loginBody)
		})
		const json = await response.json()

		// Parse response text
		try {
			const message = json['error']
			result.message = message
		} catch (e) {
			// Error is a string, or not valid JSON
			console.warn(response.statusText, e)
			result.message = `Houve um erro durante o acesso do usuário`
		}

		if (!response.ok) {
			return result
		}

		// Get JWT token
		e.cookies.set('token', json['token'], { path: '/' });

		result.message = "Acesso realizado com sucesso"
		result.success = true
		console.debug({ result, json })

		// return result
		throw redirect(303, '/api/dashboard');
	}
}
