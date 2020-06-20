<template>
  <div>
    <h1>Welcome {{ user.displayName }}</h1>
    <router-link to="/logout">Log out</router-link>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';

@Component({
  components: {
  },
})
export default class HomePage extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App;

  public user: firebase.User = this.firebase.auth().currentUser!;
}
</script>

<style scoped>
</style>
