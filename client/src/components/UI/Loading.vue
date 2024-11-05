<template>
  <div class="spinner">
    <span class="bullet"></span>
    <span class="bullet"></span>
    <span class="bullet"></span>
  </div>
</template>

<style scoped lang="scss">
.spinner {
  display: grid;
  grid-template-columns: repeat(3, 10px);
  align-items: center;
  justify-items: center;
  justify-content: space-between;
  width: 50px;
}

.bullet {
  @apply bg-cobalt-green-darker;
  width: 5px;
  height: 5px;
  border-radius: 100%;
  display: block;

  animation: bullet 0.6s infinite forwards ease-in-out;
}

.bullet:nth-child(2) {
  animation-delay: 0.2s;
}

.bullet:nth-child(3) {
  animation-delay: 0.4s;
}

@keyframes bullet {
  20% { width: 5px; height: 5px; }
  60% { width: 10px; height: 10px; }
  100% { width: 5px; height: 5px; }
}
</style>
