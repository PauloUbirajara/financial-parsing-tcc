export function load({ cookies }) {
    return {
        isLogged: Boolean(cookies.get("token", { path: '/' }))
    }
}
