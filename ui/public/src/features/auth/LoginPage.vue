<template>
  <div>
    <h1>Login page</h1>
    <section id="firebaseui-auth-container"></section>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import firebase from 'firebase';
import * as firebaseui from 'firebaseui';
import 'firebaseui/dist/firebaseui.css';
import { projectConfig } from '@/config';

@Component({
  components: {
  },
})
export default class LoginPage extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App;

  public mounted() {
    this.initializeFirebaseUI();
  }

  private initializeFirebaseUI() {
    let ui = firebaseui.auth.AuthUI.getInstance();
    if (!ui) {
      ui = new firebaseui.auth.AuthUI(firebase.auth());
    }
    const uiConfig = {
      signInSuccessUrl: '/',
      signInOptions: projectConfig.firebaseAuth.signInOptions,
    };
    ui.start('#firebaseui-auth-container', uiConfig);
  }
}
</script>

<style scoped>
</style>
