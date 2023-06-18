import { defineStore } from "pinia";
import { TSite, TSitesResponse } from "../types.ts";
import { ref } from "vue";
import { useRequest } from "../composables";

export const useSitesStore = defineStore('sites', () => {

  const sites = ref<TSite[]>([]);
  const lastFetch = ref(0); // for caching purposes I guess

  const addSite = (site: TSite) => {
    sites.value?.push(site);
  }

  const setLastFetch = () => {
    lastFetch.value = Date.now();
  }

  const fetchSites = async (): Promise<void | Error> => {
    if (Date.now() - lastFetch.value < 30000) return;

    const { data, error } = await useRequest<TSitesResponse>({
      endpoint: '/sites/all',
      needsAuth: true,
    });

    if (error) return error;

    if (data && data.sites) {
      sites.value = [...data.sites];
    }

    setLastFetch();
  }

  const checkSite = (id: number): void => {
    const site = sites.value.find(site => site.ID === id);
    if (site) {
      site.checked = !site.checked;
    }
  }

  const checkAll = (checked: boolean): void => {
    sites.value.forEach((site) => site.checked = checked);
  }

  return { sites, lastFetch, addSite, fetchSites, setLastFetch, checkSite, checkAll };
});
