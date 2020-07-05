import { register, login, logout, getInfo, refreshToken } from '@/api/user'
import { getToken, setToken, removeToken, getTokenExpire, setTokenExpire, removeTokenExpire } from '@/utils/auth'
import { resetRouter } from '@/router'

const state = {
  userId: '',
  token: getToken(),
  userName: '',
  roles: [],
  tokenExpire: getTokenExpire()
}

const mutations = {
  SET_USERID: (state, userId) => {
    state.userId = userId
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_USERNAME: (state, userName) => {
    state.userName = userName
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_TOKENEXPIRE: (state, token) => {
    state.tokenExpire = token
  }
}

const actions = {
  // user register
  register({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      register({ username: username.trim(), password: password }).then(response => {
        const data = response
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  // user login
  login({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password }).then(response => {
        const data = response
        commit('SET_TOKEN', data.token)
        commit('SET_TOKENEXPIRE', data.expire)
        setToken(data.token)
        setTokenExpire(data.expire)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },

  refreshToken({ commit }) {
    return new Promise((resolve, reject) => {
      refreshToken().then(response => {
        const data = response
        commit('SET_TOKEN', data.token)
        commit('SET_TOKENEXPIRE', data.expire)
        setToken(data.token)
        setTokenExpire(data.expire)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // get user info
  getInfo({ commit }) {
    return new Promise((resolve, reject) => {
      getInfo().then(response => {
        const data = response.data

        if (!data) {
          reject('Verification failed, please Login again.')
        }
        const { UserId, UserName, UserType } = data
        if (UserType === 1) {
          var Roles = new Array('admin')
        } else {
          Roles = new Array('dev')
        }
        commit('SET_USERID', UserId)
        commit('SET_ROLES', Roles)
        commit('SET_USERNAME', UserName)
        data['Roles'] = Roles
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state }) {
    return new Promise((resolve, reject) => {
      logout(state.token).then(() => {
        commit('SET_TOKEN', '')
        commit('SET_TOKENEXPIRE', '')
        commit('SET_ROLES', '')
        commit('SET_USERNAME', '')
        commit('SET_USERID', '')
        removeToken()
        removeTokenExpire()
        resetRouter()
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      commit('SET_TOKENEXPIRE', '')
      removeToken()
      removeTokenExpire()
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
