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
    <diary-preview-coverable v-else :paragraphs="paragraphs" :cover="covered" />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide, InjectReactive } from 'vue-property-decorator';
import { DiaryMachine } from './machine';
import DiaryPreviewCoverButton from './DiaryPreviewCoverButton.vue';
import DiaryPreviewCoverable from './DiaryPreviewCoverable.vue';
import { Loading } from '@/components';

@Component({
  components: {
    DiaryPreviewCoverButton,
    DiaryPreviewCoverable,
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

  get paragraphs(): string[] {
    return this.diary.split('\n\n')
      .map((paragraph) => paragraph.trim())
      .filter((paragraph) => paragraph.length > 0)
      .reverse();
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

.button-panel button {
  font-size: 1.1em;
}

.button-panel .reload-button {
  float: right;
}
</style>
