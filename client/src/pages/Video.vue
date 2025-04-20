<template>
  <video
    src="https://files.ricr.dev/media/cebe8952e101a038-teleport-demo_comp.mp4"
    controls
  />
</template>

<script setup lang="ts">
import { useSeoMeta } from '@unhead/vue';

useSeoMeta({
  title: 'Site-Manager Video Sample',
  description: 'Site-Manager Video Sample Description',
  ogTitle: 'Site-Manager Video Sample',
  ogDescription: 'Site-Manager Video Sample Description',
  ogImage:
    'https://files.ricr.dev/media/t/cebe8952e101a038-teleport-demo_comp.mp4-thumbnail.png',
  ogVideo:
    'https://files.ricr.dev/media/cebe8952e101a038-teleport-demo_comp.mp4',
  ogVideoType: 'video/mp4',
  ogVideoWidth: 1280,
  ogVideoHeight: 720,
  ogType: 'video.other',
  twitterCard: 'summary_large_image',
});
</script>
