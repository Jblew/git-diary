<template>
  <div>
    <authenticated-header />

    <diary-panel />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import { DiaryPanel } from '@/features/diary';
import { AuthenticatedHeader } from '@/components';
import HomeHeader from './HomeHeader.vue';

@Component({
  components: {
    DiaryPanel,
    AuthenticatedHeader,
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
