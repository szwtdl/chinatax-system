import request from '@/utils/request'

export function StudentLogoList(query) {
  return request({
    url: '/admin/student/logs/list',
    method: 'get',
    params: query
  })
}

