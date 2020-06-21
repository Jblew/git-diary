<template>
  <div class="diary-preview-covered">
    <div
      v-for="(paragraphHeight, index) in paragraphHeights"
      :key="index"
      class="cover-paragraph"
      :style="'height:' + paragraphHeight + ';'"
    ></div>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide, InjectReactive } from 'vue-property-decorator';

@Component
export default class DiaryPreviewCovered extends Vue {
  @Prop({ required: true, type: String })
  public diary!: string;

  get paragraphs(): string[] {
    return this.diary.split('\n\n');
  }

  get paragraphHeights(): number[] {
    const ps = this.paragraphs.map((paragraph) => paragraph.length / 5);
    console.log(ps);
    return ps;
  }
}
</script>

<style scoped>
.diary-preview-covered {
  width: 100%;
}
.cover-paragraph {
  width: 90%;
  margin: 0 auto;
  margin-bottom: 0.5rem;
}
</style>
