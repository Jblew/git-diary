<template>
  <div class="diary-preview-uncovered">
    <p v-for="lines in paragraphsLines" :key="lines">
      <template v-for="line in lines"> {{ line }}<br :key="line" /> </template>
    </p>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide, InjectReactive } from 'vue-property-decorator';

@Component
export default class DiaryPreviewUncovered extends Vue {
  @Prop({ required: true, type: String })
  public paragraphs!: string[];

  get paragraphsLines(): string[][] {
    return this.paragraphs.map((p) => p.split('\n'));
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
