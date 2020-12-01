import Vue from 'vue';
import Router from 'vue-router';
Vue.use(Router);

const Contents = () => import('@views/contents/Contents.vue');
const Raw = () => import('@views/contents/Raw.vue');
const NotFound = () => import('@views/pageStatus/404.vue');
const InternalServerError = () => import('@views/pageStatus/500.vue');

export default () => {
  const router = new Router({
    mode: 'history',
    base: '/',
    routes: [
      {
        path: '/',
        component: Contents,
        meta: {
          root: 'home',
        },
        props: {
          url: {
            api: {
              distinctWorkload: '/api/v1/versions/read/distinct/workload',
              endpoint: '/api/v1/versions/read/home',
            },
          },
        },
      },
      {
        path: '/workload/:workload/platform/:platform',
        component: Contents,
        meta: {
          root: 'platform',
        },
        props: {
          url: {
            api: {
              endpoint: '/api/v1/versions/read/platform',
            },
          },
        },
      },
      {
        path: '/workload/:workload/platform/:platform/:page',
        component: Contents,
        meta: {
          root: 'platform',
        },
        props: {
          url: {
            api: {
              endpoint: '/api/v1/versions/read/platform',
            },
          },
        },
      },
      {
        path: '/workload/:workload/platform/:platform/environment/:environment',
        component: Contents,
        meta: {
          root: 'environment',
        },
        props: {
          url: {
            api: {
              endpoint: '/api/v1/versions/read/environment',
            },
          },
        },
      },
      {
        path:
          '/workload/:workload/platform/:platform/environment/:environment/:page',
        component: Contents,
        meta: {
          root: 'environment',
        },
        props: {
          url: {
            api: {
              endpoint: '/api/v1/versions/read/environment',
            },
          },
        },
      },
      {
        path:
          '/workload/:workload/platform/:platform/environment/:environment/raw/:version',
        component: Raw,
        meta: {
          root: 'raw',
        },
        props: {
          url: {
            api: {
              endpoint: '/api/v1/versions/read/raw',
            },
          },
        },
      },
      {
        path: '*',
        component: NotFound,
        meta: {
          root: 'notFound',
        },
      },
      {
        path: '/404',
        component: NotFound,
        meta: {
          root: 'notFound',
        },
      },
      {
        path: '/500',
        component: InternalServerError,
        meta: {
          root: 'internalServerError',
        },
      },
    ],
  });
  return router;
};
