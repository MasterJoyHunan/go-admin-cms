import { asyncRoutes, constantRoutes } from '@/router'

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
function filterAsyncRoutes(routes, authKey) {
    const res = []
    for (const v of routes) {
        if (!authKey.includes(v.name)) {
            continue
        }
        if (v.children) {
            v.children = filterAsyncRoutes(v.children, authKey)
        }
        res.push(v)
    }
    return res
}

const state = {
    // 所有路由
    routes: [],
    // 新增路由
    addRoutes: [],
    // 权限 -- 判断路由的 name 是否包含在内
    authKey: []
}

const mutations = {
    SET_ROUTES: (state, routes) => {
        state.addRoutes = routes
        state.routes = constantRoutes.concat(routes)
    },
    SET_AUTH_KEY: (state, authKey) => {
        state.authKey = authKey
    }
}

const actions = {
    generateRoutes({ commit, rootGetters }) {
        const { auth } = rootGetters['user']
        const authKey = []
        for (const v of auth) {
            authKey.push(v.Ext)
        }
        commit('SET_AUTH_KEY', authKey)
        const accessedRoutes = filterAsyncRoutes(asyncRoutes, authKey)
        accessedRoutes.push({ path: '*', redirect: '/404', hidden: true })
        commit('SET_ROUTES', accessedRoutes)
        return Promise.resolve(accessedRoutes)
    }
}

export default {
    namespaced: true,
    state,
    mutations,
    actions
}
