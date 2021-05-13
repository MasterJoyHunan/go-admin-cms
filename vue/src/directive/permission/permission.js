import store from '@/store'

export default {
    inserted(el, { value }, vnode) {
        const hasAuth = store.getters && store.getters.authKey
        if (value && typeof value == 'string') {
            if (!hasAuth.includes(value)) {
                el.parentNode && el.parentNode.removeChild(el)
            }
        } else {
            throw new Error(`need auth-keys! Like v-permission="admin-user-add"`)
        }
    }
}
