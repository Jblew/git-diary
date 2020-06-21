<template>
  <div>
    <span v-if="publishing">Servabam...</span>
    <textarea v-model="entry" placeholder="Quomodo te habes?"></textarea>
    <button @click="publish()">Serva!</button>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, InjectReactive, Provide } from 'vue-property-decorator';
import { DiaryMachine } from './machine';
import DiaryPreviewCoverButton from './DiaryPreviewCoverButton.vue';

@Component({
  components: {
    DiaryPreviewCoverButton,
  },
})
export default class DiaryPreview extends Vue {
  @Inject('machine')
  public machine!: DiaryMachine;
  @InjectReactive('state')
  public state!: DiaryMachine['state'];

  public entry: string = '';

  get diary(): string {
    return this.state.context.diary;
  }

  get covered(): boolean {
    return this.state.context.covered;
  }

  get publishing(): boolean {
    return this.state.matches('publishing');
  }

  public cover() {
    this.machine.send('COVER_DIARY');
  }

  public uncover() {
    this.machine.send('UNCOVER_DIARY');
  }
}
</script>

<style scoped>
</style>
