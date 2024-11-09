import type { BaseResponse, TActiveSessionsResponse, TSession } from '@types';
import { type Ref, ref } from 'vue';
import { defineStore } from 'pinia';
import { useRequest } from '@composables';

export const useActiveSessionsStore = defineStore('activeSessions', () => {
  const activeSessions: Ref<TSession[]> = ref([]);

  const fetchActiveSessions = async () => {
    const { data, error } = await useRequest<TActiveSessionsResponse>({
      endpoint: '/user/sessions',
    });

    if (error) return { error };

    if (data) {
      activeSessions.value = [...data.active_sessions];
    }
  };

  const deleteSession = async (id: number) => {
    const { error, data } = await useRequest<BaseResponse>({
      endpoint: `/user/sessions/${id}`,
      method: 'DELETE',
    });

    if (error) return { error, data: null };

    activeSessions.value = activeSessions.value.filter(
      (session) => session.id !== id,
    );

    return { error: null, data };
  };

  return { activeSessions, fetchActiveSessions, deleteSession };
});
