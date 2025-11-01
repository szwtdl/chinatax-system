import request from '@/utils/request'

export function SchoolList(query) {
  return request({
    url: '/admin/school/list',
    method: 'GET',
    params: query
  })
}

export function SchoolCreate(data) {
  return request({
    url: '/admin/school/create',
    method: 'post',
    data
  })
}

export function SchoolInfo(query) {
  return request({
    url: '/admin/school/info',
    method: 'GET',
    params: query
  })
}

export function SchoolUpdate(data) {
  return request({
    url: '/admin/school/update',
    method: 'post',
    data
  })
}

export function status(data) {
  return request({
    url: '/admin/school/status',
    method: 'post',
    data
  })
}

export function SchoolDelete(data) {
  return request({
    url: '/admin/school/delete',
    method: 'post',
    data
  })
}

export function SchoolGenerate(data) {
  return request({
    url: '/admin/school/generate',
    method: 'post',
    data
  })
}

export function GenerateToken(data) {
  return request({
    url: '/admin/school/token',
    method: 'post',
    data
  })
}
