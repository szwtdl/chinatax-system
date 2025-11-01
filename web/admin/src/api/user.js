import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/public/login',
    method: 'post',
    data
  })
}

export function captcha() {
  return request({
    url: '/public/captcha',
    method: 'get',
    params: { }
  })
}

export function getInfo(token) {
  return request({
    url: '/public/info',
    method: 'get',
    params: { }
  })
}

export function resetToken() {
  return request({
    url: '/public/resetToken',
    method: 'post'
  })
}

export function logout() {
  return request({
    url: '/public/logout',
    method: 'post'
  })
}
