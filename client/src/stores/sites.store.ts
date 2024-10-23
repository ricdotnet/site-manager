import type { TSIteResponse, TSite, TSitesResponse } from '@types';
import { defineStore } from 'pinia';
import { ref } from 'vue';
import { unwrap } from '@utils';
import { useRequest } from '@composables';
import { useRoute } from 'vue-router';

export const useSitesStore = defineStore('sites', () => {
  const sites = ref<TSite[]>([]);
  const site = ref<TSite>();

  const route = useRoute();

  const addSite = (site: TSite) => {
    sites.value?.push(site);
  };

  const removeSites = (ids: number[]): void => {
    sites.value = sites.value.reduce((filtered: TSite[], site: TSite) => {
      if (!ids.includes(site.ID)) filtered.push(site);
      return filtered;
    }, []);
  };

  const fetchSites = async (): Promise<Error | undefined> => {
    const { data, error } = await useRequest<TSitesResponse>({
      endpoint: '/site/all',
      needsAuth: true,
    });

    if (error) return error;

    if (data?.sites) {
      sites.value = [...data.sites];
    }
  };

  const fetchSite = async (): Promise<Error | undefined> => {
    const { data, error } = await useRequest<TSIteResponse>({
      endpoint: `/site/${route.params.id}`,
      needsAuth: true,
    });

    if (error) return error;

    if (data?.site) {
      site.value = data.site;
      site.value.config = data.site.config;
    }
  };

  const checkSite = (id: number): void => {
    sites.value.forEach((site) => {
      if (site.ID === id) {
        site.checked = !site.checked;
      }
    });
  };

  const checkAll = (checked: boolean): void => {
    sites.value = sites.value.map((site) => {
      site.checked = checked;
      return site;
    });
  };

  const getSite = (): TSite => {
    return <TSite>unwrap(site);
  };

  return {
    sites,
    addSite,
    fetchSites,
    fetchSite,
    checkSite,
    checkAll,
    removeSites,
    getSite,
  };
});
