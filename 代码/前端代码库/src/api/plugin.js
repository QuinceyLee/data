import request from '@/utils/request'


export function plugins() {
    return request({url: 'plugins', method: 'get'})
}
