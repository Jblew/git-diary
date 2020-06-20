import Vue from 'vue';
import App from './App.vue';
import { constructFirebaseApp } from './config';

Vue.config.productionTip = false;

const firebase = constructFirebaseApp();

new Vue({
  provide() {
    return {
      firebase,
    };
  },
  render: (h) => h(App),
}).$mount('#app');
