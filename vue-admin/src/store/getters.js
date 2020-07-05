const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  userId: state => state.user.userId,
  userName: state => state.user.userName,
  roles: state => state.user.roles,
  permission_routes: state => state.permission.routes
}
export default getters
