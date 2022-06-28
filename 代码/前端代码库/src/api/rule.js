import request from '@/utils/request'

export function list() {
    return request({url: 'rules', method: 'get'})
}

export function validateLua() {
    return request({url: 'rules', method: 'get'})
}
export function create(data) {
    return request({url: 'rules', method: 'post', data})
}

export function remove(uuid) {
    return request({
        method: 'delete',
        url: 'rules?uuid=' + uuid
    })
}

export function detail(uuid) {
    return request({
        method: 'get',
        url: 'rules?uuid=' + uuid
    })
}

export function validateRule(data) {
    return request({url: 'validateRule', method: 'post', data})
}
