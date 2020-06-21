<template>
  <div>
    <h1>Exibaris...</h1>
    <a @click="backToHome()">Domi revertere</a>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import firebase from 'firebase/app';
import { projectConfig } from '@/config';

@Component
export default class LogoutPage extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App;

  public user: firebase.User | null = null;

  public mounted() {
    this.listenToAuth();
  }

  public backToHome() {
    this.$router.push('/');
  }

  private listenToAuth() {
    this.firebase.auth().onAuthStateChanged((user) => {
      this.user = user || null;
      this.logout();
    });
  }

  private logout() {
    this.firebase.auth().signOut().then(() => this.backToHome());
  }
}
</script>

<style scoped>
</style>
