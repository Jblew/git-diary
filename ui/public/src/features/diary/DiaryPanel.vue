<template>
  <div>
    <diary-error />
    <diary-publish />
    <diary-preview />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide, ProvideReactive } from 'vue-property-decorator';
import { functions } from 'firebase';
import DiaryPublish from './DiaryPublish.vue';
import DiaryPreview from './DiaryPreview.vue';
import DiaryError from './DiaryError.vue';
import { DiaryMachine } from './machine';
import { interpretMachine } from './interpretMachine';

@Component({
  components: {
    DiaryPublish,
    DiaryPreview,
    DiaryError,
  },
})
export default class DiaryPanel extends Vue {
  @Provide('machine')
  public machine: DiaryMachine = interpretMachine();

  @ProvideReactive('state')
  public state: DiaryMachine['state'] = this.machine.initialState;

  public beforeMount() {
    this.machine.onTransition((s) => {
      this.state = s;
      // tslint:disable no-console
      console.log('State', s);
    }).start();
  }
}
</script>

<style scoped>
</style>
