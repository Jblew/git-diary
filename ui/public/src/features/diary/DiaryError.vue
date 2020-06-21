<template>
  <div>
    <error v-if="errorText">Error accideras: {{ errorText }}</error>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, InjectReactive, Provide } from 'vue-property-decorator';
import { DiaryMachine } from './machine';
import { Error } from '@/components';

@Component({
  components: {
    Error,
  },
})
export default class DiaryError extends Vue {
  @Inject('machine')
  public machine!: DiaryMachine;
  @InjectReactive('state')
  public state!: DiaryMachine['state'];

  get errorText(): string {
    return this.state.context.errorText;
  }
}
</script>

<style scoped>
</style>
