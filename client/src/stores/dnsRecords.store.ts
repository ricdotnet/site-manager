import type {
  TDNSRecord,
  TDNSRecordFormProcess,
  TDNSRecordsResponse,
} from '@types';
import { defineStore } from 'pinia';
import { ref } from 'vue';
import { useRequest } from '@composables';
import { useRoute } from 'vue-router';

export const useDnsRecordsStore = defineStore('dnsRecords', () => {
  const dnsRecords = ref<TDNSRecord[]>([]);

  const route = useRoute();

  const fetchDnsRecords = async () => {
    const { data, error } = await useRequest<TDNSRecordsResponse>({
      endpoint: `/domains/dns/${route.params.domain}/${route.params.type}`,
      needsAuth: true,
    });

    if (error) return error;

    if (data) {
      dnsRecords.value = data.records.records ? [...data.records.records] : [];
    }
  };

  const addDnsRecord = async (data: TDNSRecordFormProcess['data']) => {
    if (!data) return;

    const nextId = dnsRecords.value.length + 1;

    const { error } = await useRequest({
      endpoint: `/domains/dns/${route.params.domain}/${route.params.type}`,
      method: 'PUT',
      needsAuth: true,
      payload: data,
    });

    if (error) {
      throw new Error(error);
    }

    const newDnsRecords = [...dnsRecords.value];
    newDnsRecords.push({
      id: nextId,
      host: data.host,
      value: data.value,
      ttl: data.ttl,
      status: 'Active',
    });

    dnsRecords.value = newDnsRecords.sort((a, b) => {
      return a.host.localeCompare(b.host);
    });
  };

  return { addDnsRecord, dnsRecords, fetchDnsRecords };
});
