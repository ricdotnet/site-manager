<template>
  <div class="button-group">
    <slot/>
  </div>
</template>

<style scoped lang="scss">
  .button-group {
    @apply flex space-x-3 py-3 justify-end;
  }
</style>
