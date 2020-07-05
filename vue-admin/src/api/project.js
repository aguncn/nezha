import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/project/list',
    method: 'get',
    params: query
  })
}

export function fetchProject(data) {
  return request({
    url: '/api/v1/project/detail/' + data,
    method: 'get'
  })
}


export function createProject(data) {
  return request({
    url: '/api/v1/project/create',
    method: 'post',
    data
  })
}

export function updateProject(data) {
  return request({
    url: '/api/v1/project/update',
    method: 'put',
    data
  })
}

export function deleteProject(data) {
  return request({
    url: '/api/v1/project/delete/' + data,
    method: 'delete',
    data
  })
}
