<template>
  <div class="loading">
    <loading-spinner />

    <div class="loading-text">
      <span><slot></slot></span>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import LoadingSpinner from './LoadingSpinner.vue';

@Component({
  components: {
    LoadingSpinner,
  },
})
export default class Loading extends Vue {
}
</script>
<style>
.loading {
  width: 100%;
  text-align: center;
  margin-top: 3rem;
  margin-bottom: 1rem;
  padding-bottom: 1rem;
  background: #eee;
}

.loading .loading-spinner {
  display: block;
  margin-top: -35px;
}

.loading .loading-text {
  display: block;
  margin-top: -25px;
  color: #888;
  font-size: 1.1em;
}
</style>
