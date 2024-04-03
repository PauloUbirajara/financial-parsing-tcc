import type { Handle } from '@sveltejs/kit'
import { json, redirect } from '@sveltejs/kit'
import { constants as httpConstants } from 'http2'

import constants from '../constants.ts'

function onUnauthorized(event) {
  console.warn('Usuário não autorizado')

  // Clear token cookie
  event.cookies.set('token', '', { path: '/' })
  event.cookies.set('username', '', { path: '/' })

  if (!event.url.pathname.startsWith('/api')) {
    return
  }
  // Permanent redirect to login
  throw redirect(308, '/auth/login')
}

function onAuthorized(event, userJson) {
  console.warn('Usuário autorizado')

  // Set username
  event.cookies.set('username', userJson['username'], { path: '/'} )

  if (!event.url.pathname.startsWith('/auth')) {
    return
  }

  // Permanent redirect to dashboard if trying to authenticate while logged in
  throw redirect(308, '/api/dashboard')
}

export async function handle({ event, resolve }) {
  // Get JWT stored in cookies
  const token = event.cookies.get('token')

  // Check user in backend to see if token still is valid
  const userResponse = await fetch(constants.API_USER_URL, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })

  let userJson = {}
  let authorized = false

  try {
    userJson = await userResponse.json()
    authorized = (
      userResponse.status !== httpConstants.HTTP_STATUS_UNAUTHORIZED &&
      ![undefined, "undefined"].includes(userJson['username'])
    )
  } catch (e) {
    console.warn("Error when getting user validation response as JSON")
  }

  console.debug({userJson})

  if (!authorized) {
    onUnauthorized(event)
  }

  if (authorized) {
    onAuthorized(event, userJson)
  }

  // Continue to next request handlers
  const response = await resolve(event);
  return response;
}
