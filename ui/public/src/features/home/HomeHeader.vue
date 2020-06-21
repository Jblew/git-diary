<template>
  <div class="home-header">
    <img src="/logo.svg" class="logo" />

    <h1>Ephemeris, {{ user.displayName }}</h1>
    <div class="logout-link">
      <router-link to="/logout">Obliviscere me</router-link>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';

@Component
export default class HomePage extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App;

  public user: firebase.User = this.firebase.auth().currentUser!;
}
</script>

<style scoped>
.home-header {
  text-align: left;
  padding: 1rem;
  background: #eee;
  margin-bottom: 2rem;
}

.home-header h1 {
  margin: 0;
}

.logout-link {
  text-align: left;
  width: 100%;
  font-size: 75%;
}

.logo {
  height: 3rem;
  float: right;
}
</style>
