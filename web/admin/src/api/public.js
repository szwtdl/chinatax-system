import request from '@/utils/request'

export function lastUpdate(query) {
  return request({
    url: '/latest',
    method: 'get',
    params: query
  })
}
