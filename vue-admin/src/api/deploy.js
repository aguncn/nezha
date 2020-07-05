import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/deploy/list',
    method: 'get',
    params: query
  })
}

export function fetchDeploy(data) {
  return request({
    url: '/deploy/detail/' + data,
    method: 'get'
  })
}


export function createDeploy(data) {
  return request({
    url: '/deploy/create',
    method: 'post',
    data
  })
}

export function updateDeploy(data) {
  return request({
    url: '/deploy/update',
    method: 'put',
    data
  })
}

export function deleteDeploy(data) {
  return request({
    url: '/deploy/delete/' + data,
    method: 'delete',
    data
  })
}

export function submitDeploy(data) {
  return request({
    url: '/deploy/submit',
    method: 'post',
    data
  })
}

export function statusDeploy(data) {
  return request({
    url: '/deploy/status',
    method: 'post',
    data
  })
}

export function getDockerTag(query) {
  return request({
    url: '/deploy/dockerTag',
    method: 'get',
    params: query
  })
}

export function getOkErrorCount(query) {
  return request({
    url: '/deploy/okErrorCount',
    method: 'get',
    params: query
  })
}

export function getCycleCount(query) {
  return request({
    url: '/deploy/cycleCount',
    method: 'get',
    params: query
  })
}
