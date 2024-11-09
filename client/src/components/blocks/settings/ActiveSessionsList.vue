<template>
  <Empty
    message="Unable to load active sessions"
    v-if="activeSessions.length === 0"
  />
  <div
    class="flex justify-between w-full p-5 gap-5 border border-light-border dark:border-dark-border rounded-md [&:not(last)]:mb-5"
    v-for="(activeSession, i) of activeSessions"
    :key="i"
  >
    <ComputerDesktopIcon
      v-if="activeSession.parsed_user_agent.is_desktop"
      class="w-10"
    />
    <div class="flex flex-col flex-grow gap-2">
      <span class="block">
        {{ activeSession.parsed_user_agent.os }} -
        {{ activeSession.parsed_user_agent.name }}
      </span>
      <span class="block text-xs">
        {{ activeSession.ip_address }}
        <span v-if="activeSession.this_device" class="text-red-500 ml-2">
          This device
        </span>
      </span>
    </div>
    <div v-if="!activeSession.this_device">
      <Button
        name="delete"
        color="danger"
        @click="() => onClickDeleteSession(activeSession.id)"
        :is-actioning="isDeletingSession"
      >
        Logout
      </Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Button, Empty } from '@components';
import { ComputerDesktopIcon } from '@heroicons/vue/24/solid';
import { messages } from '@utils';
import { ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useActiveSessionsStore } from '@stores';
import { useToaster } from '@composables';

const { addToast } = useToaster();

const activeSessionsStore = useActiveSessionsStore();
const { activeSessions } = storeToRefs(activeSessionsStore);

await activeSessionsStore.fetchActiveSessions();

const isDeletingSession = ref(false);

const onClickDeleteSession = async (id: number) => {
  isDeletingSession.value = true;

  const { error, data } = await activeSessionsStore.deleteSession(id);

  if (error) {
    isDeletingSession.value = false;
    addToast('error', 'Could not delete session');
    return;
  }

  isDeletingSession.value = false;

  if (data?.message_code) {
    addToast('success', messages.user[data.message_code]);
  }
};
</script>
