<template>
  <div>
    <h1>Publishing panel</h1>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide } from 'vue-property-decorator';
import { DiaryMachine } from './machine';

@Component
export default class DiaryPublish extends Vue {
  @Inject('machine')
  public machine!: DiaryMachine;
  @Inject('state')
  public state!: DiaryMachine['state'];
}
</script>

<style scoped>
</style>
