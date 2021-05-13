import store from '@/store'

/**
 * @param {String} value
 * @returns {Boolean}
 * @example see @/views/permission/directive.vue
 */
export default function checkPermission(value) {
    if (value && value instanceof String) {
        const hasAuth = store.getters && store.getters.authKey
        return hasAuth.includes(value)
    } else {
        console.error(`need auth-keys! Like v-permission="admin-user-add"`)
        return false
    }
}
