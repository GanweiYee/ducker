import request from '@/utils/request'

export const login = (data) => {
    return request({
        url: '/system/base/login',
        method: 'post',
        data
    })
}

export const signup = (data) => {
    return request({
        url: '/system/base/signup',
        method: 'post',
        data
    })
}