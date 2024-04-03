export function load({ cookies }) {
  const token = cookies.get("token", { path: '/' })
  const username = cookies.get("username", { path: '/' })

  const isLogged = Boolean(token)

  return {
    username,
    isLogged 
  }
}
