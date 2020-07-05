import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/environment/list',
    method: 'get',
    params: query
  })
}

export function fetchEnvironment(data) {
  return request({
    url: '/api/v1/environment/detail/' + data,
    method: 'get'
  })
}


export function createEnvironment(data) {
  return request({
    url: '/api/v1/environment/create',
    method: 'post',
    data
  })
}

export function updateEnvironment(data) {
  return request({
    url: '/api/v1/environment/update',
    method: 'put',
    data
  })
}

export function deleteEnvironment(data) {
  return request({
    url: '/api/v1/environment/delete/' + data,
    method: 'delete',
    data
  })
}
