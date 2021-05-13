
import request from '@/utils/request'

export function getRoleTree(id = 0) {
    return request({
        url: '/auth/role',
        params: {
            'id': id
        }
    })
}
