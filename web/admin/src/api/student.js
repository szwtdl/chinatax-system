import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/admin/student/list',
    method: 'GET',
    params: query
  })
}

export function getInfo(data) {
  return request({
    url: '/admin/student/info',
    method: 'GET',
    params: data
  })
}

export function ImportExcel(data) {
  return request({
    url: '/admin/student/import',
    method: 'POST',
    data
  })
}

export function createUser(data) {
  return request({
    url: '/admin/student/add',
    method: 'POST',
    data
  })
}

export function updateUser(data) {
  return request({
    url: '/admin/student/update',
    method: 'POST',
    data
  })
}

export function deleteUser(data) {
  return request({
    url: '/admin/student/delete',
    method: 'POST',
    data
  })
}

export function schoolList(data) {
  return request({
    url: '/admin/student/school/list',
    method: 'GET',
    params: data
  })
}

export function getCourseList(data) {
  return request({
    url: '/admin/student/course_list',
    method: 'POST',
    data
  })
}

// export function getCourseList(data) {
//   return request({
//     url: '/admin/student/course',
//     method: 'post',
//     data
//   })
// }

export function getAllCourse(data) {
  return request({
    url: '/admin/student/course_type',
    method: 'post',
    data
  })
}

export function Login(data) {
  return request({
    url: '/user/login',
    method: 'post',
    data
  })
}

export function VerifyImage() {
  return request({
    url: '/admin/student/verify_code',
    method: 'get'
  })
}

export function PlatformList() {
  return request({
    url: '/admin/strategy/platform',
    method: 'get'
  })
}

export function switchPlatform(data) {
  return request({
    url: '/admin/student/switch',
    method: 'POST',
    data
  })
}

export function status(data) {
  return request({
    url: '/admin/student/status',
    method: 'POST',
    data
  })
}
