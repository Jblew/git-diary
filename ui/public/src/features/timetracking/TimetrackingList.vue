<template>
  <div>
    <p>List</p>
  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TimetrackingList from './TimetrackingList.vue';
import TimetrackingAdd from './TimetrackingAdd.vue';

@Component
export default class extends Vue {
  refresh() {

  }
}
</script>

<style scoped>
</style>
