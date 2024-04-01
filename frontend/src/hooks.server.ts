import type { Handle } from '@sveltejs/kit'
import { json, redirect } from '@sveltejs/kit'
import { constants as httpConstants } from 'http2'

import constants from '../constants.ts'

export async function handle({ event, resolve }) {
  if (!event.url.pathname.startsWith('/api')) {
    const response = await resolve(event);
    return response;
  }

  // Get JWT stored in cookies
  const token = event.cookies.get('token')
  console.debug({token})

  // Check user in backend to see if token still is valid
  const userResponse = await fetch(constants.API_USER_URL, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })

  if (userResponse.status === httpConstants.HTTP_STATUS_UNAUTHORIZED) {
    console.warn('Usuário não autorizado')

    // Clear token cookie
    event.cookies.set('token', '', { path: '/' })

    // Permanent redirect to login
    throw redirect(308, '/auth/login')
  }

  console.debug('Usuário autorizado')

  //Set token as bearer
  const response = await resolve(event);
  return response;
}
