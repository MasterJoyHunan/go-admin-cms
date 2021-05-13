import request from '@/utils/request'

export function getList(params) {
    return request({
        url: '/adminUser',
        params
    })
}

export function getDetail(id) {
    return request({
        url: '/adminUser/' + id,
        method: 'get'
    })
}

export function add(data) {
    return request({
        url: '/adminUser',
        method: 'post',
        data
    })
}

export function edit(id, data) {
    return request({
        url: '/adminUser/' + id,
        method: 'put',
        data
    })
}

export function del(id) {
    return request({
        url: '/adminUser/' + id,
        method: 'delete'
    })
}
