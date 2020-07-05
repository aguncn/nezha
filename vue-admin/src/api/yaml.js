import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/yaml/list',
    method: 'get',
    params: query
  })
}

export function fetchYaml(data) {
  return request({
    url: '/yaml/detail/' + data,
    method: 'get'
  })
}


export function createYaml(data) {
  return request({
    url: '/yaml/create',
    method: 'post',
    data
  })
}

export function updateYaml(data) {
  return request({
    url: '/yaml/update',
    method: 'put',
    data
  })
}

export function deleteYaml(data) {
  return request({
    url: '/yaml/delete/' + data,
    method: 'delete',
    data
  })
}
