import request from '@/utils/request'

export function List(query) {
  return request({
    url: '/city_tax/list',
    method: 'get',
    params: query
  })
}

export function Create(data) {
  return request({
    url: '/city_tax/create',
    method: 'post',
    data
  })
}

export function Update(data) {
  return request({
    url: '/city_tax/update',
    method: 'post',
    data
  })
}

export function Delete(data) {
  return request({
    url: '/city_tax/delete',
    method: 'post',
    data
  })
}

export function Sort(data) {
  return request({
    url: '/city_tax/sort',
    method: 'post',
    data
  })
}
