<template>
  <div>
    <span v-if="publishing">Servabam...</span>
    <textarea v-model="publishText" placeholder="Quomodo te habes?"></textarea>
    <button :disabled="!publishButtonEnabled" @click="publish()">Serva!</button>
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

  get publishText(): string {
    return this.state.context.publishText;
  }

  set publishText(text: string) {
    this.machine.send({ type: 'SET_PUBLISH_TEXT', text });
  }

  get publishing(): boolean {
    return this.state.matches('publishing');
  }

  get publishButtonEnabled(): boolean {
    return this.publishText.length > 0 && !this.publishing;
  }

  public publish() {
    this.machine.send('PUBLISH');
  }
}
</script>

<style scoped>
</style>
