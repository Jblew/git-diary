<template>
  <button v-if="covered" @click="$emit('uncover')">
    Uncover
  </button>
  <button v-else @click="$emit('cover')">
    Cover
  </button>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject, Provide } from 'vue-property-decorator';
import { DiaryMachine } from './machine';

@Component
export default class DiaryPreview extends Vue {
  @Prop({ required: true, type: Boolean })
  public covered!: boolean;
}
</script>

<style scoped>
</style>
