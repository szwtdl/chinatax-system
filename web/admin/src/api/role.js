import request from '@/utils/request'

export function List() {
  return request({
    url: '/role/list',
    method: 'get'
  })
}

export function Create(data) {
  return request({
    url: '/role/add',
    method: 'post',
    data
  })
}

export function Update(data) {
  return request({
    url: `/role/update`,
    method: 'post',
    data
  })
}

export function Delete(id) {
  return request({
    url: `/role/delete`,
    method: 'post',
    params: { id }
  })
}

export function permissions(query) {
  return request({
    url: `/role/permissions`,
    method: 'get',
    params: query
  })
}
