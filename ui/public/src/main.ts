import Vue from 'vue';
import { App } from './components';
import { constructFirebaseApp } from './config';
import { HomePage } from '@/features/home';
import { TimetrackingPage } from '@/features/timetracking';
import { LoginPage, LogoutPage, installFirebaseAuthHooksToVueRouter } from '@/features/auth';
import VueRouter from 'vue-router';
import VueHotkey from 'v-hotkey';
import 'firebase/app';
import 'firebase/auth';
import 'firebase/functions';
import { TimeEntryRepositoryFirestore } from './repository/TimeEntryRepositoryFirestore';

Vue.config.productionTip = false;
Vue.use(VueRouter);
Vue.use(VueHotkey);

const firebase = constructFirebaseApp();

const routes = [
  { path: '/', component: HomePage, meta: { requiresAuth: true } },
  { path: '/track', component: TimetrackingPage, meta: { requiresAuth: true } },
  { path: '/login', component: LoginPage },
  { path: '/logout', component: LogoutPage },
];
const router = new VueRouter({ routes });
installFirebaseAuthHooksToVueRouter(router);

new Vue({
  provide() {
    return {
      firebase,
      timeEntryRepository: new TimeEntryRepositoryFirestore()
    };
  },
  router,
  render: (h) => h(App),
}).$mount('#app');
