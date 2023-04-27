<template>
  <Nav></Nav>

  <div class="content mx-auto px-4">
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts">
  import Nav from './components/Nav.vue';
</script>
