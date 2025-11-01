import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/admin/task_record/list',
    method: 'get',
    params: query
  })
}
