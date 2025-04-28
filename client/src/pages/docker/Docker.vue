<template>
  <div v-if="loading">
    <p>Loading...</p>
  </div>
  <div v-else>
    <h1>Docker Containers</h1>
    <p>Docker containers:</p>
    <ul>
      <li v-for="container in containers" :key="container.id">
        {{ container.id }} - {{ container.names }} [{{ container.ports }}] ({{ container.state }})
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
  import { onMounted, ref } from 'vue';
  import { useRequest } from '@composables';

  const loading = ref(true);
  const containers = ref([]);

  onMounted(async () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const { data, error } = await useRequest<any>({
      endpoint: '/docker/containers',
    });

    if (error) {
      console.error('Error fetching containers:', error);
      return;
    }

    console.log('Containers:', data);

    containers.value = data.data;

    loading.value = false;
  });
</script>