<template>
  <div class="diary-preview-covered">
    <div
      v-for="(paragraphHeight, index) in paragraphHeights"
      :key="index"
      class="cover-paragraph"
      :style="{ height: paragraphHeight + 'px' }"
    >
      &nbsp;
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide, InjectReactive } from 'vue-property-decorator';

@Component
export default class DiaryPreviewCovered extends Vue {
  @Prop({ required: true, type: String })
  public paragraphs!: string[];

  get paragraphHeights(): number[] {
    return this.paragraphs.map((paragraph) => paragraph.length / 3 + 2).map(Math.ceil);
  }
}
</script>

<style scoped>
.diary-preview-covered {
  width: 100%;
}
.cover-paragraph {
  width: 95%;
  display: block;
  margin: 0 auto;
  margin-bottom: 0.5rem;
  background: #777;
  min-height: 2px;
}
</style>
