<template>
  <ul class="menu-bar">
    <li v-for="item of items" :key="item.url">
      <router-link :to="item.url">
        {{ item.label }}
      </router-link>
    </li>
  </ul>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { menu } from '@/config'

@Component
export default class extends Vue {
  get items(): MenuItem[] {
    return menu
  }
}

interface MenuItem {
  label: string
  url: string
}
</script>

<style scoped>
.menu-bar {
  list-style: none;
  padding: 0;
  margin: 0;
}

.menu-bar li {
  display: block;
  width: 120px;
  float: left;
  margin-left: 2px;
  border: 1px solid #000;
}

.menu-bar li a {
  display: block;
  padding: 3px;
  text-decoration: none;
  background-color: #fff;
  color: #009;
}

.menu-bar li a.router-link-exact-active {
  background-color: #009;
  color: #fff;
}

.menu-bar li a:hover {
  background-color: #009;
  color: #fff;
}
</style>
