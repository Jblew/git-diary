<template>
  <div class="diary-publish">
    <textarea v-model="publishText" placeholder="Quomodo te habes?"></textarea>
    <div class="publish-button-container">
      <button :disabled="!publishButtonEnabled" @click="publish()">
        Serva!
      </button>
    </div>
    <loading v-if="publishing">Servabam...</loading>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, InjectReactive, Provide } from 'vue-property-decorator';
import { DiaryMachine } from './machine';
import { Loading } from '@/components';
import DiaryPreviewCoverButton from './DiaryPreviewCoverButton.vue';

@Component({
  components: {
    DiaryPreviewCoverButton,
    Loading,
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
.diary-publish {
  width: 100%;
}

.diary-publish textarea {
  width: 100%;
  box-sizing: border-box;
  min-height: 5rem;
}

.publish-button-container {
  width: 100%;
  text-align: right;
}

.publish-button-container button {
  font-size: 1.5em;
  padding-left: 2rem;
  padding-right: 2rem;
}
</style>
