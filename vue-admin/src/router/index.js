import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/403',
    component: () => import('@/views/403'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: '首页',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '首页', icon: 'dashboard' }
    }]
  },

  {
    path: '/application',
    component: Layout,
    redirect: '/application/list',
    name: '应用部署',
    meta: { title: '应用部署', icon: 'example' },
    children: [
      {
        path: 'list',
        name: 'Application',
        component: () => import('@/views/application/list'),
        meta: { title: '应用列表', icon: 'table' }
      },
      {
        path: 'create',
        hidden: true,
        component: () => import('@/views/application/create'),
        name: 'CreateApplicaiton',
        meta: { title: '新建Application', icon: 'create' }
      },
      {
        path: 'update/:id',
        hidden: true,
        component: () => import('@/views/application/update'),
        name: 'EditApplicaiton',
        meta: { title: '编辑Application', icon: 'edit' }
      },
      {
        path: 'detail/:id',
        hidden: true,
        component: () => import('@/views/application/detail'),
        name: 'DetailApplicaiton',
        meta: { title: '查看Application', icon: 'detail' }
      }
    ]
  },
  {
    path: '/k8s',
    redirect: '/k8s/list',
    name: 'K8s集群',
    component: Layout,
    meta: { title: 'K8s集群', icon: 'example'},
    children: [
      {
        path: 'list',
        name: 'K8s集群',
        component: () => import('@/views/k8s/list'),
        meta: {
          title: 'K8s集群',
          icon: 'example'
        }
      },
      {
        path: 'create',
        hidden: true,
        component: () => import('@/views/k8s/create'),
        name: 'CreateK8s',
        meta: { title: '新建K8s', icon: 'create' }
      },
      {
        path: 'update/:id',
        hidden: true,
        component: () => import('@/views/k8s/update'),
        name: 'EditK8s',
        meta: { title: '编辑K8s', icon: 'edit' }
      },
      {
        path: 'detail/:id',
        hidden: true,
        component: () => import('@/views/k8s/detail'),
        name: 'DetailK8s',
        meta: { title: '查看K8s', icon: 'detail' }
      }
    ]
  },
  {
    path: '/yaml',
    redirect: '/yaml/list',
    name: 'Yaml配置',
    component: Layout,
    meta: { title: 'Yaml配置', icon: 'nested'},
    children: [
      {
        path: 'list',
        name: 'Yaml配置',
        component: () => import('@/views/yaml/list'),
        meta: {
          title: 'Yaml配置',
          icon: 'nested'
        }
      },
      {
        path: 'create',
        hidden: true,
        component: () => import('@/views/yaml/create'),
        name: 'CreateYaml',
        meta: { title: '新建Yaml', icon: 'create' }
      },
      {
        path: 'update/:id',
        hidden: true,
        component: () => import('@/views/yaml/update'),
        name: 'EditYaml',
        meta: { title: '编辑Yaml', icon: 'edit' }
      },
      {
        path: 'detail/:id',
        hidden: true,
        component: () => import('@/views/yaml/detail'),
        name: 'DetailYaml',
        meta: { title: '查看Yaml', icon: 'detail' }
      }
    ]
  },
  {
    path: '/deploy',
    component: Layout,
    redirect: '/deploy',
    hidden: true,
    children: [{
      path: '/deploy/index/:id',
      name: '部署',
      component: () => import('@/views/deploy/index'),
      meta: { title: '应用部署', icon: 'dashboard' },
      hidden: true
    },{
      path: 'list/:id',
      hidden: true,
      component: () => import('@/views/deploy/deploylist'),
      name: 'ListDeploy',
      meta: { title: '查看发布单', icon: 'detail' },
      hidden: true
    },{
      path: 'historylist/:id',
      hidden: true,
      component: () => import('@/views/deploy/historylist'),
      name: 'ListRelease',
      meta: { title: '查看历史部署', icon: 'detail' },
      hidden: true
    },{
      path: 'podlist/:id',
      hidden: true,
      component: () => import('@/views/deploy/podlist'),
      name: 'ListPod',
      meta: { title: '查看线上POD', icon: 'detail' },
      hidden: true
    }]
  },
]

// 异步挂载的路由
// 动态需要根据权限加载的路由表
export const asyncRoutes = [
  {
    path: '/user',
    component: Layout,
    children: [
      {
        path: 'index',
        name: '用户管理',
        component: () => import('@/views/user/index'),
        meta: {
          title: '用户管理',
          icon: 'user',
          roles: ['admin']
        }
      }
    ]
  },
  {
    path: '/environment',
    component: Layout,
    redirect: '/environment/list',
    name: '环境管理',
    meta: {
      title: '环境管理',
      icon: 'tree',
      roles: ['admin']
    },
    children: [
      {
        path: 'list',
        name: 'Environment',
        component: () => import('@/views/environment/list'),
        meta: { title: '环境列表', icon: 'tree' }
      },
      {
        path: 'create',
        hidden: true,
        component: () => import('@/views/environment/create'),
        name: 'CreateEnvironment',
        meta: { title: '新建环境', icon: 'create' }
      },
      {
        path: 'update/:id',
        hidden: true,
        component: () => import('@/views/environment/update'),
        name: 'EditEnvironment',
        meta: { title: '编辑环境', icon: 'edit' }
      }
    ]
  },
  {
    path: '/project',
    component: Layout,
    redirect: '/project/list',
    name: '项目管理',
    meta: {
      title: '项目管理',
      icon: 'skill',
      roles: ['admin']
    },
    children: [
      {
        path: 'list',
        name: 'Project',
        component: () => import('@/views/project/list'),
        meta: { title: '项目列表', icon: 'skill' }
      },
      {
        path: 'create',
        hidden: true,
        component: () => import('@/views/project/create'),
        name: 'CreateProject',
        meta: { title: '新建Project', icon: 'create' }
      },
      {
        path: 'update/:id',
        hidden: true,
        component: () => import('@/views/project/update'),
        name: 'EditProject',
        meta: { title: '编辑Project', icon: 'edit' }
      }
    ]
  },
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
