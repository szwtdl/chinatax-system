import request from '@/utils/request'

export function Welcome(query) {
  return request({
    url: '/dashboard/welcome',
    method: 'get',
    params: query
  })
}
