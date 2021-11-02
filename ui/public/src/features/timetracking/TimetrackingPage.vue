<template>
  <div>
    <authenticated-header />

    <timetracking-panel />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import TimetrackingPanel from './TimetrackingPanel.vue';
import { AuthenticatedHeader } from '@/components';
import HomeHeader from './HomeHeader.vue';

@Component({
  components: {
    TimetrackingPanel,
    AuthenticatedHeader,
  },
})
export default class TimetrackingPage extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App;

  public user: firebase.User = this.firebase.auth().currentUser!;
}
</script>

<style scoped>
</style>
