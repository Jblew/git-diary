<template>
  <div id="app">
    <MainPanel />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import MainPanel from './components/MainPanel.vue';

@Component({
  components: {
    MainPanel,
  },
})
export default class App extends Vue { }
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
