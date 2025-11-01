import request from '@/utils/request'

export function getFileList(query) {
  return request({
    url: '/admin/media/list',
    method: 'get',
    params: query
  })
}

export function uploadFile(data) {
  return request({
    url: '/admin/media/upload',
    method: 'post',
    data
  })
}

export function deleteFile(data) {
  return request({
    url: '/admin/media/delete',
    method: 'post',
    data
  })
}

export function getFileInfo(data) {
  return request({
    url: '/admin/media/info',
    method: 'get',
    data
  })
}

export function cutVideo(data) {
  return request({
    url: '/admin/media/cutVideo',
    method: 'post',
    timeout: 15 * 1000,
    data
  })
}

// 分类列表
export function getMediaList(query) {
  return request({
    url: '/admin/media_type/list',
    method: 'get',
    params: query
  })
}

// 增加分类
export function addMediaType(data) {
  return request({
    url: '/admin/media_type/add',
    method: 'post',
    data
  })
}

// 更新分类
export function updateMediaType(data) {
  return request({
    url: '/admin/media_type/update',
    method: 'post',
    data
  })
}

// 删除分类
export function deleteMediaType(data) {
  return request({
    url: '/admin/media_type/delete',
    method: 'post',
    data
  })
}
