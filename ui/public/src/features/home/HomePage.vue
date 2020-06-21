<template>
  <div>
    <h1>Ephemeris, {{ user.displayName }}</h1>
    <router-link to="/logout">Log out</router-link>
    <hr />

    <diary-panel />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import { DiaryPanel } from '@/features/diary';

@Component({
  components: {
    DiaryPanel,
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
