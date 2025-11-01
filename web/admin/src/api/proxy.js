import request from '@/utils/request'

export function ProxyList(query) {
  return request({
    url: '/admin/proxy/list',
    method: 'get',
    params: query
  })
}

export function ProxyDelete(data) {
  return request({
    url: '/admin/proxy/delete',
    method: 'post',
    data
  })
}

export function ProxyRefresh(query) {
  return request({
    url: '/admin/proxy/refresh',
    method: 'GET',
    params: query
  })
}

export function ProxyTest(data) {
  return request({
    url: '/admin/proxy/test',
    method: 'post',
    data
  })
}

export function ProxyPing(data) {
  return request({
    url: '/admin/proxy/ping',
    method: 'post',
    data
  })
}

export function ProxyRegions(data) {
  return request({
    url: '/admin/proxy/regions',
    method: 'post',
    data
  })
}

export function ProxySwitch(data) {
  return request({
    url: '/admin/proxy/switch',
    method: 'post',
    data
  })
}

