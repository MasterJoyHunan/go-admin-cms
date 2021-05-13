const getters = {
    sidebar: state => state.app.sidebar,
    device: state => state.app.device,
    user: state => state.user,
    permission_routes: state => state.permission.routes,
    authKey: state => state.permission.authKey
}
export default getters
