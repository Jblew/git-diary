import Vue from 'vue';
import { App } from './components';
import { constructFirebaseApp } from './config';
import { HomePage } from '@/features/home';
import { LoginPage } from '@/features/auth';
import VueRouter from 'vue-router';

Vue.config.productionTip = false;

const firebase = constructFirebaseApp();

const routes = [
  { path: '/', component: HomePage },
  { path: '/login', component: LoginPage },
];

new Vue({
  provide() {
    return {
      firebase,
    };
  },
  router: new VueRouter({
    routes,
  }),
  render: (h) => h(App),
}).$mount('#app');
