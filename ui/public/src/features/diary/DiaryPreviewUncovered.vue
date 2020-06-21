<template>
  <div class="diary-preview-uncovered">
    <pre
      >{{ formattedText }}
    </pre>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide, InjectReactive } from 'vue-property-decorator';

@Component
export default class DiaryPreviewCovered extends Vue {
  @Prop({ required: true, type: String })
  public paragraphs!: string[];

  get formattedText(): string {
    return this.paragraphs.join('\n\n');
  }
}
</script>

<style scoped>
.diary-preview-uncovered {
  width: 100%;
  text-align: left;
  padding: 2rem;
  padding-top: 1rem;
}
</style>
