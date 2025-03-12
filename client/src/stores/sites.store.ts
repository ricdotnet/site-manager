import type {
  BaseResponse,
  TCertificate,
  TCertificatesResponse,
  TSite,
  TSiteResponse,
  TSitesResponse,
} from '@types';
import { defineStore } from 'pinia';
import { ref } from 'vue';
import { unwrap } from '@utils';
import { useRequest } from '@composables';
import { useRoute } from 'vue-router';

export const useSitesStore = defineStore('sites', () => {
  const sites = ref<TSite[]>([]);
  const certificates = ref<TCertificate[]>([]);
  const site = ref<TSite>();

  const route = useRoute();

  const addSite = (site: TSite) => {
    sites.value?.push(site);
  };

  const removeSites = (ids: number[]): void => {
    sites.value = sites.value.reduce((filtered: TSite[], site: TSite) => {
      if (!ids.includes(site.id)) {
        filtered.push(site);
      }
      return filtered;
    }, []);
  };

  const fetchSites = async (): Promise<string | undefined> => {
    const { data, error } = await useRequest<TSitesResponse>({
      endpoint: '/site/all',
    });

    if (error) return error;

    if (data?.sites) {
      sites.value = [...data.sites];
    }
  };

  const fetchCertificates = async (): Promise<string | undefined> => {
    const { data, error } = await useRequest<TCertificatesResponse>({
      endpoint: '/site/certificates',
    });

    if (error) {
      return error;
    }

    if (data?.certificates) {
      certificates.value = [...data.certificates];
    }
  };

  const fetchSite = async (): Promise<string | undefined> => {
    const { data, error } = await useRequest<TSiteResponse>({
      endpoint: `/site/${route.params.id}`,
    });

    if (error) {
      return error;
    }

    if (data?.site) {
      site.value = data.site;
      site.value.config = data.config;
    }
  };

  const checkSite = (id: number): void => {
    sites.value.forEach((site) => {
      if (site.id === id) {
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

  const getSites = (): TSite[] => {
    return sites.value.filter((site) => !site.config_only);
  };

  const reloadNginx = async (): Promise<{
    error: string | undefined;
    data: string | undefined;
  }> => {
    const { error, data } = await useRequest<BaseResponse>({
      endpoint: '/site/reload',
    });

    return { error, data: data?.message_code };
  };

  return {
    sites,
    certificates,
    addSite,
    fetchSites,
    fetchSite,
    checkSite,
    checkAll,
    removeSites,
    getSite,
    getSites,
    reloadNginx,
    fetchCertificates,
  };
});
