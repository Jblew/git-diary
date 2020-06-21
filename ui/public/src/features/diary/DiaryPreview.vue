<template>
  <div>
    <h1>Preview panel</h1>
    <div>loading={{ loading }}</div>
    <div>covered={{ covered }}</div>
    <diary-preview-cover-button
      :covered="covered"
      @uncover="uncover()"
      @cover="cover()"
    />
    <p>
      {{ diary }}
    </p>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide } from 'vue-property-decorator';
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
  @Inject('state')
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
}
</script>

<style scoped>
</style>
