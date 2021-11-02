<template>
  <div>
    <timetracking-add @added="refresh()" />
    <timetracking-list ref="timetrackingList" />
  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TimetrackingList from './TimetrackingList.vue';
import TimetrackingAdd from './TimetrackingAdd.vue';

@Component({
  components: {
    TimetrackingAdd,
    TimetrackingList,
  },
})
export default class TimetrackingPanel extends Vue {
  refresh() {
    (this.$refs.timetrackingList as any).refresh()
  }
}
</script>

<style scoped>
</style>
