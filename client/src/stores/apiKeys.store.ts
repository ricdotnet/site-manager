import type { TApiKey } from '@types';
import { defineStore } from 'pinia';
import { ref } from 'vue';
import { useRequest } from '@composables';

export const useApiKeysStore = defineStore('apiKeys', () => {
  const apiKeys = ref<TApiKey[]>([]);

  const fetchApiKeys = async () => {
    const { data, error } = await useRequest<TApiKey[]>({
      endpoint: '/settings',
    });

    if (error) return error;

    if (data) {
      apiKeys.value = [...data];
    }
  };

  const addApiKey = async (
    apiKey: Pick<TApiKey, 'key' | 'value' | 'is_api_key'>,
  ) => {
    const { data, error } = await useRequest<TApiKey>({
      endpoint: '/settings',
      method: 'PUT',
      payload: apiKey,
    });

    if (error) return error;

    if (data) {
      const existingKey = apiKeys.value.find((key) => key.key === data.key);
      if (existingKey) {
        apiKeys.value = apiKeys.value.map((key) =>
          key.key === data.key ? data : key,
        );
        return;
      }
      apiKeys.value = [...apiKeys.value, data];
    }
  };

  const deleteApiKey = async (apiKey: TApiKey) => {
    const { error } = await useRequest<TApiKey>({
      endpoint: `/settings/${apiKey.key}`,
      method: 'DELETE',
    });

    if (error) return error;

    apiKeys.value = apiKeys.value.filter((key) => key.id !== apiKey.id);
  };

  return { apiKeys, fetchApiKeys, addApiKey, deleteApiKey };
});
