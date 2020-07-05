import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/k8s/list',
    method: 'get',
    params: query
  })
}

export function fetchK8s(data) {
  return request({
    url: '/k8s/detail/' + data,
    method: 'get'
  })
}


export function createK8s(data) {
  console.log(data, "!!!!!!!!!!")
  return request({
    url: '/k8s/create',
    method: 'post',
    data
  })
}

export function updateK8s(data) {
  return request({
    url: '/k8s/update',
    method: 'put',
    data
  })
}

export function deleteK8s(data) {
  return request({
    url: '/k8s/delete/' + data,
    method: 'delete',
    data
  })
}
