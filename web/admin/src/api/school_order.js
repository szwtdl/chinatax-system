import request from '@/utils/request'

export function SchoolOrderList(query) {
  return request({
    url: '/admin/school/order/list',
    method: 'GET',
    params: query
  })
}

export function SchoolOrderSettlement(data) {
  return request({
    url: '/admin/school/order/settlement',
    method: 'POST',
    data
  })
}

export function SchoolOrderDelete(data) {
  return request({
    url: '/admin/school/order/delete',
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
