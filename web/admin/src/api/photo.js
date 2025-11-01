import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/admin/photo/list',
    method: 'get',
    params: query
  })
}

export function upload(data) {
  return request({
    url: '/admin/photo/upload',
    method: 'post',
    data
  })
}

export function remove(data) {
  return request({
    url: '/admin/photo/delete',
    method: 'post',
    data
  })
}
