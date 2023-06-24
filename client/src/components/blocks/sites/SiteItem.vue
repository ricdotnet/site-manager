<template>
  <div class="p-5">
    <h1 class="text-2xl mb-5">VHosts config</h1>
    <div>Here you can edit your vhosts config file linked to this domain.
      When saving your site will be updated with the new configs.</div>

<!--    {{ getSite() }}-->
  </div>
</template>

<script setup lang="ts">
  import { useSitesStore } from "@stores";

  const sitesStore = useSitesStore();
  const { getSite } = sitesStore;
</script>
