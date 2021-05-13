import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'

const requestTimeOut = () => {
    if (process.env.NODE_ENV == 'development') {
        return 1000 * 60 * 5
    } else {
        return 1000 * 5
    }
}

const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
    timeout: requestTimeOut() // request timeout
})

service.interceptors.request.use(config => config,
    error => {
        console.log('request err ', error) // for debug
        return Promise.reject(error)
    }
)

service.interceptors.response.use(
    response => {
        const { data } = response
        // 正常情况
        if (data.code != 1) {
            Message({
                message: data.msg,
                type: 'error',
                duration: 5 * 1000
            })
            return Promise.reject(data.msg)
        }
        return data
    },
    async error => {
        const { status, data } = error.response
        switch (status) {
            case 500:
                Message({
                    message: '系统错误',
                    type: 'error',
                    duration: 5 * 1000
                })
                break
            case 403:
                Message({
                    message: '没有权限',
                    type: 'error',
                    duration: 5 * 1000
                })
                break
            case 401:
                await MessageBox.confirm(data.msg, '登录过期，请重新登录', {
                    confirmButtonText: '重新登录',
                    cancelButtonText: '取消',
                    type: 'warning'
                })
                await store.dispatch('user/resetToken')
                location.reload()
                break
            default:
                Message({
                    message: '未知错误',
                    type: 'error',
                    duration: 5 * 1000
                })
                break
        }
        return Promise.reject(error)
    }
)

export default service
