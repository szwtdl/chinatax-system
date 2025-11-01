import request from '@/utils/request'

export function List(query) {
  return request({
    url: '/admin/school/token/list',
    method: 'GET',
    params: query
  })
}

export function update(data) {
  return request({
    url: '/admin/school/token/update',
    method: 'POST',
    data: data
  })
}

export function del(data) {
  return request({
    url: '/admin/school/token/delete',
    method: 'POST',
    data: data
  })
}

export function create(data) {
  return request({
    url: '/admin/school/token/create',
    method: 'POST',
    data: data
  })
}
