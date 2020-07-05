import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/pod/list',
    method: 'get',
    params: query
  })
}

export function fetchDeploy(data) {
  return request({
    url: '/pod/detail/' + data,
    method: 'get'
  })
}

