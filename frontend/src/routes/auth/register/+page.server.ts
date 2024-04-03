import type { Actions } from './$types';

import constants from '../../../../constants'

async function getRegisterBody(e): object {
	const formData = await e.request.formData();
	const registerBody = {
		email: formData.get('email'),
		username: formData.get('username'),
		password: formData.get('password'),
		confirmPassword: formData.get('confirm-password'),
	}

	return registerBody
}

type RegisterResult = {
	title: string
	success: boolean
	message?: string
}

export const actions: Actions = {
	default: async (e): RegisterResult => {
		let result: RegisterResult = {
			title: "Cadastro de usuário",
			success: false
		}

		// Get content from form
		const registerBody = await getRegisterBody(e)

		// Send contents to backend for register
		const response = await fetch(constants.API_REGISTER_URL, {
			headers: {
				'Content-Type': 'application/json'
			},
			method: 'POST',
			body: JSON.stringify(registerBody)
		})

		const json = await response.json()

		// Parse response text
		try {
			const message = json['error']
			result.message = message
		} catch (e) {
			// Error is a string, or not valid JSON
			console.warn(response.statusText, e)
			result.message = `Houve um erro durante o cadastro do usuário`
		}

		if (!response.ok) {
			return result
		}

		result.message = "Cadastro realizado com sucesso"
		result.success = true
		console.debug({ result, json })

		return result
	}
};
