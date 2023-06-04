import { defineStore } from "pinia";
import { TSite } from "../types.ts";
import { ref } from "vue";

export const useSitesStore = defineStore('sites', () => {

  const sites = ref<TSite[]>([]);
  const lastFetch = ref(0); // for caching purposes I guess

  const addSite = (site: TSite) => {
    sites.value?.push(site);
  }

  const setLastFetch = () => {
    lastFetch.value = Date.now();
  }

  return { sites, lastFetch, addSite, setLastFetch };
});
