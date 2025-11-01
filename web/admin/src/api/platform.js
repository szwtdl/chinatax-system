import request from '@/utils/request'

export function List(query) {
  return request({
    url: '/admin/platform/list',
    method: 'get',
    params: query
  })
}

export function Create(data) {
  return request({
    url: '/admin/platform/create',
    method: 'post',
    data
  })
}

export function Update(data) {
  return request({
    url: '/admin/platform/update',
    method: 'post',
    data
  })
}

export function Delete(data) {
  return request({
    url: '/admin/platform/delete',
    method: 'post',
    data
  })
}

export function Status(data) {
  return request({
    url: '/admin/platform/status',
    method: 'post',
    data
  })
}

export function Debug(data) {
  return request({
    url: '/admin/platform/debug',
    method: 'post',
    data
  })
}

export function Sort(data) {
  return request({
    url: '/admin/platform/sort',
    method: 'post',
    data
  })
}

export function Auth(data) {
  return request({
    url: '/admin/platform/auth',
    method: 'post',
    data
  })
}

export function Refresh(data) {
  return request({
    url: '/admin/platform/refresh',
    method: 'post',
    data
  })
}
