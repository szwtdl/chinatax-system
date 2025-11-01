import request from '@/utils/request'

export function getConfigs(query) {
  return request({
    url: '/system/get',
    method: 'get',
    params: query
  })
}

export function updateConfigs(data) {
  return request({
    url: '/system/update',
    method: 'post',
    data
  })
}
