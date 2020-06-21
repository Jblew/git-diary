<template>
  <button v-hotkey="btnKeymap" v-if="covered" @click="$emit('uncover')">
    Denuda
  </button>
  <button v-hotkey="btnKeymap" v-else @click="$emit('cover')">
    Tege
  </button>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide } from 'vue-property-decorator';
import { DiaryMachine } from './machine';

@Component
export default class DiaryPreview extends Vue {
  @Prop({ required: true, type: Boolean })
  public covered!: boolean;

  get btnKeymap() {
    return {
      'cmd+1': this.toggle,
      'alt+w': this.toggle,
    };
  }

  public toggle() {
    if (this.covered) {
      this.$emit('uncover');
    } else {
      this.$emit('cover');
    }
  }
}
</script>

<style scoped>
</style>
