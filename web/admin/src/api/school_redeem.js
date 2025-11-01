import request from '@/utils/request'

export function List(query) {
  return request({
    url: '/admin/school/redeem/list',
    method: 'GET',
    params: query
  })
}

export function add(data) {
  return request({
    url: '/admin/school/redeem/create',
    method: 'POST',
    data: data
  })
}

export function del(data) {
  return request({
    url: '/admin/school/redeem/delete',
    method: 'POST',
    data: data
  })
}

export function update(data) {
  return request({
    url: '/admin/school/redeem/update',
    method: 'POST',
    data: data
  })
}
