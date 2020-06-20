import Vue from 'vue';
import { App } from './components';
import { constructFirebaseApp } from './config';
import { HomePage } from '@/features/home';
import { LoginPage, LogoutPage, installFirebaseAuthHooksToVueRouter } from '@/features/auth';
import VueRouter from 'vue-router';
import 'firebase/app';
import 'firebase/auth';

Vue.config.productionTip = false;
Vue.use(VueRouter);

const firebase = constructFirebaseApp();

const routes = [
  { path: '/', component: HomePage, meta: { requiresAuth: true } },
  { path: '/login', component: LoginPage },
  { path: '/logout', component: LogoutPage },
];
const router = new VueRouter({ routes });
installFirebaseAuthHooksToVueRouter(router);

new Vue({
  provide() {
    return {
      firebase,
    };
  },
  router,
  render: (h) => h(App),
}).$mount('#app');
