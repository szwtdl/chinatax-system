import request from '@/utils/request'

export function taskList(query) {
  return request({
    url: '/admin/task/list',
    method: 'get',
    params: query
  })
}

export function addTask(data) {
  return request({
    url: '/admin/task/add',
    method: 'post',
    data: data
  })
}

export function addSchoolTask(data) {
  return request({
    url: '/admin/task/addSchoolTask',
    method: 'post',
    data: data
  })
}

export function batchAddTask(data) {
  return request({
    url: '/admin/task/batch_task',
    method: 'post',
    data: data
  })
}

export function refreshProgress(data) {
  return request({
    url: '/admin/task/refresh',
    method: 'post',
    data: data
  })
}

export function proxySelectList(query) {
  return request({
    url: '/admin/task/proxy/list',
    method: 'GET',
    params: query
  })
}

export function strategyList(query) {
  return request({
    url: '/admin/task/strategy/list',
    method: 'GET',
    params: query
  })
}

export function strategyUpdate(data) {
  return request({
    url: '/admin/task/strategy/update',
    method: 'POST',
    data: data
  })
}

export function proxyUpdate(data) {
  return request({
    url: '/admin/task/proxy/update',
    method: 'POST',
    data: data
  })
}

export function deleteTask(data) {
  return request({
    url: '/admin/task/delete',
    method: 'POST',
    data: data
  })
}

export function clearTask(data) {
  return request({
    url: '/admin/task/delete',
    method: 'POST',
    data: data
  })
}

export function startTask(data) {
  return request({
    url: '/admin/task/start',
    method: 'post',
    data: data
  })
}

export function stopTask(data) {
  return request({
    url: '/admin/task/stop',
    method: 'post',
    data: data
  })
}

export function running(query) {
  return request({
    url: '/admin/task/running',
    method: 'get',
    params: query
  })
}
