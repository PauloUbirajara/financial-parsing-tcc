const baseAPIUrl = import.meta.env.VITE_API_BASE_URL

export default {
		API_LOGIN_URL: `${baseAPIUrl}/auth/login`,
		API_REGISTER_URL: `${baseAPIUrl}/auth/register`,
		API_USER_URL: `${baseAPIUrl}/auth/user`
}
