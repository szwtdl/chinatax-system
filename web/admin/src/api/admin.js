import request from '@/utils/request'

export function AdminList(query) {
  return request({
    url: '/account/list',
    method: 'get',
    params: query
  })
}

export function addAdmin(data) {
  return request({
    url: '/account/create',
    method: 'post',
    data
  })
}

export function infoAdmin(query) {
  return request({
    url: '/account/info',
    method: 'get',
    params: query
  })
}

export function updateAdmin(data) {
  return request({
    url: '/account/update',
    method: 'post',
    data
  })
}

export function deleteAdmin(data) {
  return request({
    url: '/account/delete',
    method: 'post',
    data
  })
}
