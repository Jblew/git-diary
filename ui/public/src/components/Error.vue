<template>
  <div class="error">
    <slot></slot>
  </div>
</template>
<style>
.error {
  width: 100%;
  border: 2px solid red;
  background: rgba(255, 0, 0, 0.3);
  padding: 1rem;
  text-align: center;
}
</style>
