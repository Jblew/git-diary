<template>
  <div>
    <span v-if="errorText" class="error">{{ errorText }}</span>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide } from 'vue-property-decorator';
import { DiaryMachine } from './machine';

@Component
export default class DiaryError extends Vue {
  @Inject('machine')
  public machine!: DiaryMachine;
  @Inject('state')
  public state!: DiaryMachine['state'];

  get errorText(): string {
    return this.state.context.errorText;
  }
}
</script>

<style scoped>
.error {
  color: red;
}
</style>
