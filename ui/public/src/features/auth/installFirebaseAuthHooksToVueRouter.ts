import type VueRouter from 'vue-router';
import 'firebase/auth';
import firebase from 'firebase/app';

export function installFirebaseAuthHooksToVueRouter(router: VueRouter) {
  router.beforeEach(async (to, _, next) => {
    const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);
    if (requiresAuth && !await getCurrentFirebaseUser()) {
      next('/login');
    } else {
      next();
    }
  });
}

function getCurrentFirebaseUser() {
  return new Promise((resolve, reject) => {
    const unsubscribe = firebase.auth().onAuthStateChanged((user) => {
      unsubscribe();
      resolve(user);
    }, reject);
  });
}
