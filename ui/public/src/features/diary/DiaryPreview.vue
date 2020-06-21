<template>
  <div class="diary-preview">
    <div class="button-panel">
      <diary-preview-cover-button
        class="cover-button"
        :covered="covered"
        @uncover="uncover()"
        @cover="cover()"
      />
      <button class="reload-button" @click="reload()">Refice</button>
    </div>

    <loading v-if="loading">Onerabas...</loading>
    <diary-preview-covered v-else-if="covered" :diary="diary" />
    <diary-preview-uncovered v-else :diary="diary" />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide, InjectReactive } from 'vue-property-decorator';
import { DiaryMachine } from './machine';
import DiaryPreviewCoverButton from './DiaryPreviewCoverButton.vue';
import DiaryPreviewCovered from './DiaryPreviewCovered.vue';
import DiaryPreviewUncovered from './DiaryPreviewUncovered.vue';
import { Loading } from '@/components';

@Component({
  components: {
    DiaryPreviewCoverButton,
    DiaryPreviewCovered,
    DiaryPreviewUncovered,
    Loading,
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
.diary-preview {
  width: 100%;
}

.button-panel {
  width: 100%;
  margin-bottom: 1rem;
  background: #eee;
  padding: 0.5rem;
  box-sizing: border-box;
  text-align: left;
}

.button-panel .reload-button {
  float: right;
}
</style>
