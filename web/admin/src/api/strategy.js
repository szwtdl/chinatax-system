import request from '@/utils/request'

export function StrategyList(query) {
  return request({
    url: '/admin/strategy/list',
    method: 'get',
    params: query
  })
}

export function StrategyCreate(data) {
  return request({
    url: '/admin/strategy/create',
    method: 'post',
    data
  })
}

export function StrategyUpdate(data) {
  return request({
    url: '/admin/strategy/update',
    method: 'post',
    data
  })
}

export function StrategyDelete(data) {
  return request({
    url: '/admin/strategy/delete',
    method: 'post',
    data
  })
}

export function PlatformList() {
  return request({
    url: '/admin/strategy/platform',
    method: 'get'
  })
}

export function Status(data) {
  return request({
    url: '/admin/strategy/status',
    method: 'post',
    data
  })
}

export function Sort(data) {
  return request({
    url: '/admin/strategy/sort',
    method: 'post',
    data
  })
}
