<template>
  <div>
    <diary-publish />
    <diary-preview />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide } from 'vue-property-decorator';
import { functions } from 'firebase';
import DiaryPublish from './DiaryPublish.vue';
import DiaryPreview from './DiaryPreview.vue';
import { DiaryMachine } from './machine';
import { interpretMachine } from './interpretMachine';

@Component({
  components: {
    DiaryPublish,
    DiaryPreview,
  },
})
export default class DiaryPanel extends Vue {
  @Provide('machine')
  public machine: DiaryMachine = interpretMachine();

  @Provide('state')
  public state: DiaryMachine['state'] = this.machine.initialState;

  public beforeMount() {
    this.machine.onTransition((s) => this.state = s).start();
  }
}
</script>

<style scoped>
</style>
