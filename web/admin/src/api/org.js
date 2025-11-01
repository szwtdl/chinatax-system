import request from '@/utils/request'

export function fetchOrgList(query) {
  return request({
    url: '/admin/org/list',
    method: 'get',
    params: query
  })
}

export function addOrg(data) {
  return request({
    url: '/admin/org/add',
    method: 'post',
    data: data
  })
}

export function updateOrg(data) {
  return request({
    url: '/admin/org/update',
    method: 'post',
    data: data
  })
}

export function deleteOrg(query) {
  return request({
    url: '/admin/org/delete',
    method: 'get',
    params: query
  })
}

export function getMemberList(data) {
  return request({
    url: '/admin/org/users',
    method: 'post',
    data
  })
}

export function addMembers(data) {
  return request({
    url: '/admin/student/add_users',
    method: 'post',
    timeout: 10 * 1000,
    data
  })
}
