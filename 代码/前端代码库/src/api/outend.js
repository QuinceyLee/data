import request from '@/utils/request'

export function list() {
    return request({url: 'outends', method: 'get'})
}
export function create(data) {
    return request({url: 'outends', method: 'post', data})
}

export function remove(uuid) {
    return request({
        method: 'delete',
        url: 'outends?uuid=' + uuid
    })
}

export function detail(uuid) {
    return request({
        method: 'get',
        url: 'outends?uuid=' + uuid
    })
}
