import Vue from 'vue';
import App from './App.vue';
Vue.config.productionTip = false;
import createRouter from '@router';
import createStore from '@store';
import { sync } from 'vuex-router-sync';
import VueMeta from 'vue-meta';
Vue.use(VueMeta);

// register filters globally.
import * as filters from '@filters';
Object.keys(filters).forEach((key) => {
  Vue.filter(key, filters[key]);
});

import '@plugins/axios';

export default () => {
  const store = createStore();
  const router = createRouter();
  sync(store, router);
  return new Vue({
    router,
    store,
    render: (h) => h(App),
  });
};
