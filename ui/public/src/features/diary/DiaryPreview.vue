<template>
  <div>
    <div v-if="loading">Onerabas...</div>
    <div>Occulta={{ covered }}</div>
    <diary-preview-cover-button
      :covered="covered"
      @uncover="uncover()"
      @cover="cover()"
    />
    <button @click="reload()">Refice</button>
    <p>
      {{ diary }}
    </p>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide, InjectReactive } from 'vue-property-decorator';
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

  get diary(): string {
    return this.state.context.diary;
  }

  get covered(): boolean {
    return this.state.context.covered;
  }

  get loading(): boolean {
    return this.state.matches('loading');
  }

  public cover() {
    this.machine.send('COVER_DIARY');
  }

  public uncover() {
    this.machine.send('UNCOVER_DIARY');
  }

  public reload() {
    this.machine.send('RELOAD');
  }
}
</script>

<style scoped>
</style>
