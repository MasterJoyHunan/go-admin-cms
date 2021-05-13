import router from './router'
import store from './store'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken } from '@/utils/auth' // get token from cookie
import getPageTitle from '@/utils/get-page-title'

NProgress.configure({ showSpinner: false }) // NProgress Configuration

const whiteList = ['/login'] // no redirect whitelist

router.beforeEach(async(to, from, next) => {
    // start progress bar
    NProgress.start()

    // set page title
    document.title = getPageTitle(to.meta.title)

    // determine whether the user has logged in
    const hasToken = getToken()

    if (hasToken) {
        if (to.path === '/login') {
            // if is logged in, redirect to the home page
            next({ path: '/' })
            NProgress.done()
        } else {
            if (store.getters.user.id == 0) {
                await store.dispatch('user/getUserInfo')
            }
            if (store.getters.user.id == 0) {
                await store.dispatch('user/resetToken')
                next(`/login?redirect=${to.path}`)
                NProgress.done()
                return
            }
            if (store.getters.permission_routes.length == 0) {
                const accessRoutes = await store.dispatch('permission/generateRoutes')
                router.addRoutes(accessRoutes)
                next({ ...to, replace: true })
                return
            }
            next()
        }
    } else {
        /* has no token*/
        if (whiteList.indexOf(to.path) !== -1) {
            // in the free login whitelist, go directly
            next()
        } else {
            // other pages that do not have permission to access are redirected to the login page.
            next(`/login?redirect=${to.path}`)
            NProgress.done()
        }
    }
})

router.afterEach(() => {
    // finish progress bar
    NProgress.done()
})
