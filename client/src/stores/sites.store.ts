import { ref } from "vue";
import { defineStore } from "pinia";
import { TSite, TSIteResponse, TSitesResponse } from "@types";
import { useRequest } from "@composables";
import { useRoute } from "vue-router";
import { unwrap } from "@utils";

export const useSitesStore = defineStore("sites", () => {
	const sites = ref<TSite[]>([]);
	const lastFetch = ref(0); // for caching purposes I guess
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

	const setLastFetch = () => {
		lastFetch.value = Date.now();
	};

	const fetchSites = async (): Promise<Error> => {
		if (Date.now() - lastFetch.value < 30000) return;

		const { data, error } = await useRequest<TSitesResponse>({
			endpoint: "/site/all",
			needsAuth: true,
		});

		if (error) return error;

		if (data?.sites) {
			sites.value = [...data.sites];
		}

		setLastFetch();
	};

	const fetchSite = async (): Promise<Error> => {
		const { data, error } = await useRequest<TSIteResponse>({
			endpoint: `/site/${route.params.id}`,
			needsAuth: true,
		});

		if (error) return error;

		if (data?.site) {
			site.value = data.site;
			site.value.config = data.config;
		}
	};

	const checkSite = (id: number): void => {
		const site = sites.value.find((site) => site.ID === id);
		if (site) {
			site.checked = !site.checked;
		}
	};

	const checkAll = (checked: boolean): void => {
		// biome-ignore lint/complexity/noForEach: we can allow here a for each... it is easier
		sites.value.forEach((site) => {
			site.checked = checked;
			return site;
		});
	};

	const getSite = (): TSite => {
		return <TSite>unwrap(site);
	};

	return {
		sites,
		lastFetch,
		addSite,
		fetchSites,
		fetchSite,
		checkSite,
		checkAll,
		removeSites,
		getSite,
	};
});
