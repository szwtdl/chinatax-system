import request from '@/utils/request'

export function List(query) {
  return request({
    url: '/partner/list',
    method: 'get',
    params: query
  })
}

export function Create(data) {
  return request({
    url: '/partner/create',
    method: 'post',
    data
  })
}

export function Update(data) {
  return request({
    url: '/partner/update',
    method: 'post',
    data
  })
}

export function Delete(data) {
  return request({
    url: '/partner/delete',
    method: 'post',
    data
  })
}

export function Status(data) {
  return request({
    url: '/partner/status',
    method: 'post',
    data
  })
}
