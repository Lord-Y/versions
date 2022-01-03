import { RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    component: () => import('@/views/Contents.vue'),
    meta: {
      root: 'home',
      activeLink: '/',
    },
    props: {
      url: {
        api: {
          default: '/api/v1/versions/read/home',
        },
      },
    },
  },
  {
    path: '/dashboards/deployments',
    component: () => import('@/views/dashboards/Deployments.vue'),
    meta: {
      menu: 'dashboards',
      root: 'dashboards',
      activeLink: '/dashboards/deployments',
    },
    props: {
      url: {
        api: {
          default: '/api/v1/versions/stats/latest',
        },
      },
    },
  },
  {
    path: '/workload/:workload/platform/:platform',
    component: () => import('@/views/Contents.vue'),
    meta: {
      menu: 'platform',
      root: 'platform',
    },
    props: {
      url: {
        api: {
          default: '/api/v1/versions/read/platform',
        },
      },
    },
  },
  {
    path: '/workload/:workload/platform/:platform/:page',
    component: () => import('@/views/Contents.vue'),
    meta: {
      menu: 'platform',
      root: 'platform',
    },
    props: {
      url: {
        api: {
          default: '/api/v1/versions/read/platform',
        },
      },
    },
  },
  {
    path: '/workload/:workload/platform/:platform/environment/:environment',
    component: () => import('@/views/Contents.vue'),
    meta: {
      menu: 'platform',
      root: 'environment',
    },
    props: {
      url: {
        api: {
          default: '/api/v1/versions/read/environment',
        },
      },
    },
  },
  {
    path: '/workload/:workload/platform/:platform/environment/:environment/:page',
    component: () => import('@/views/Contents.vue'),
    meta: {
      menu: 'platform',
      root: 'environment',
    },
    props: {
      url: {
        api: {
          default: '/api/v1/versions/read/environment',
        },
      },
    },
  },
  {
    path: '/workload/:workload/platform/:platform/environment/:environment/raw/:version',
    component: () => import('@/views/Raw.vue'),
    meta: {
      root: 'raw',
    },
    props: {
      url: {
        api: {
          default: '/api/v1/versions/read/raw',
        },
      },
    },
  },
  {
    path: '/raw/:version',
    component: () => import('@/views/Raw.vue'),
    meta: {
      root: 'raw',
    },
    props: {
      url: {
        api: {
          default: '/api/v1/versions/read/raw',
        },
      },
    },
  },
  {
    path: '/:pathMatch(.*)*',
    component: () => import('@/views/pageStatus/404.vue'),
  },
  {
    path: '/404',
    component: () => import('@/views/pageStatus/404.vue'),
  },
  {
    path: '/500',
    component: () => import('@/views/pageStatus/500.vue'),
  },
]

export default routes