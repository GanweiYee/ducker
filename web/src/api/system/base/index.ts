import request from "@/utils/request";

export const login = (data: any) => {
    return request({
        url: '/api/system/base/login',
        method: 'post',
        data
    })
}

export const logout = (data: any) => {
    return request({
        url: '/api/system/base/logout',
        method: 'post',
        data
    })
}