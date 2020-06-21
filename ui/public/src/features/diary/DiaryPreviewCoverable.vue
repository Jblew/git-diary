<template>
  <div class="diary-preview-coverable">
    <p v-for="lines in paragraphsLines" :key="lines">
      <template v-for="line in lines"> {{ line }}<br :key="line" /> </template>
    </p>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide, InjectReactive } from 'vue-property-decorator';

@Component
export default class DiaryPreviewCoverable extends Vue {
  @Prop({ required: true, type: String })
  public paragraphs!: string[];

  @Prop({ required: true, type: Boolean })
  public cover!: boolean;

  get paragraphsLines(): string[][] {
    return this.paragraphs.map((p) => this.mapParagraph(p));
  }

  private mapParagraph(paragraph: string): string[] {
    return paragraph.split('\n').map((line) => this.cover ? this.coverLine(line) : line);
  }

  private coverLine(line: string): string {
    return line.replace(/\S/gi, '█');
  }
}
</script>

<style scoped>
.diary-preview-coverable {
  width: 100%;
  text-align: left;
  padding: 2rem;
  padding-top: 1rem;
}
</style>
