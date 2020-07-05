import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/application/list',
    method: 'get',
    params: query
  })
}

export function recentList(query) {
  return request({
    url: '/application/recentList',
    method: 'get',
    params: query
  })
}

export function fetchApplication(data) {
  return request({
    url: '/application/detail/' + data,
    method: 'get'
  })
}


export function createApplication(data) {
  return request({
    url: '/application/create',
    method: 'post',
    data
  })
}

export function updateApplication(data) {
  return request({
    url: '/application/update',
    method: 'put',
    data
  })
}

export function deleteApplication(data) {
  return request({
    url: '/application/delete/' + data,
    method: 'delete',
    data
  })
}
