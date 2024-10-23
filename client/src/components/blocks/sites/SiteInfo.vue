<template>
  <div class="site-info-container">
    <h1 class="site-info-container__domain">{{ getSite().domain }}</h1>
    <div class="site-info-container__grid">
      <div class="site-info-container__grid--cell">
        <div class="uppercase label">Status</div>
        {{ getSite().enabled ? 'enabled' : 'disabled' }}
      </div>
      <div class="site-info-container__grid--cell">
        <div class="uppercase label">Has SSL</div>
        {{ getSite().has_ssl ? 'yes' : 'no' }}
      </div>
      <div class="site-info-container__grid--cell">
        <div class="uppercase label">Created on</div>
        {{ createdAt }}
      </div>
      <div class="site-info-container__grid--cell">
        <div class="uppercase label">Last updated</div>
        {{ lastUpdatedAt }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { formatDate } from '@utils';
import { useSitesStore } from '@stores';

const sitesStore = useSitesStore();
const { getSite } = sitesStore;

const createdAt = computed(() => formatDate(getSite().created_at));
const lastUpdatedAt = computed(() => formatDate(getSite().updated_at));
</script>

<style scoped lang="scss">
.site-info-container {
  @apply px-10
  py-10
  border-b
  border-light-border
  dark:border-dark-border;

  &__domain {
    @apply text-4xl
    pb-3;
  }

  &__grid {
    @apply grid
    grid-rows-2
    grid-cols-2
    gap-y-5;

    &--cell {
      @apply h-10
      flex
      flex-col;

      .label {
        @apply text-dark-dim dark:text-light-dim font-extralight;
      }
    }
  }
}
</style>
