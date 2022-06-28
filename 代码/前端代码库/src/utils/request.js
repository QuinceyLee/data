import axios from 'axios'
import {Message} from 'element-ui'
import store from '@/store'
import {getToken} from '@/utils/auth'

// create an axios instance
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API, // URL
    timeout: 9000 // 超时时间非常长，给9秒
})

// request interceptor
service.interceptors.request.use(config => {
    if (store.getters.token) {
        config.headers['X-Token'] = getToken()
        config.headers['cache-control'] = 'no-cache'
    }
    return config
}, error => {
    console.log(error)
    return Promise.reject(error)
})

// response interceptor
//
// 这里比较坑，axios这个库的状态码全部是在业务层判断的，所以不得不在业务层设计状态码，也就是说所有的请求都是200，然后code来区分
// 这里不遵守http的语义我觉得很拉垮,导致我不得不改自己的后端代码来配合这套东西。  ---- 2021-11-13@wwhai
//
service.interceptors.response.use(response => {
    const res = response.data
    if (res.code !== 200) {
        Message({
            message: res.msg || 'Error',
            type: 'error',
            duration: 5 * 1000
        })
        return Promise.reject(new Error(res.msg || 'Error'))
    }
    return res

}, error => {
    console.log('service.interceptors.response error: ' + error)
    Message({
        message: error,
        type: 'error',
        duration: 5 * 1000
    })
    return Promise.reject(error)
})

export default service
