<template>
  <div class="stack">
    <slot/>
  </div>
</template>

<style scoped lang="scss">
  .stack {
    @apply flex flex-col items-center mt-10;
  }
</style>
