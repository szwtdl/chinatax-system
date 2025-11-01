import request from '@/utils/request'

export function List(query) {
  return request({
    url: '/permission/list',
    method: 'get',
    params: query
  })
}

export function Info(query) {
  return request({
    url: '/permission/info',
    method: 'get',
    params: query
  })
}

export function Create(data) {
  return request({
    url: '/permission/add',
    method: 'post',
    data: data
  })
}

export function Update(data) {
  return request({
    url: '/permission/update',
    method: 'post',
    data: data
  })
}

export function Delete(query) {
  return request({
    url: '/permission/delete',
    method: 'get',
    params: query
  })
}
