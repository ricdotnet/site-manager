import axios from "axios";
import { Ref, ref } from "vue";
import { unwrap } from "../utils";

type RequestVerb = 'GET' | 'POST' | 'PATCH' | 'PUT' | 'DELETE';

type RequestOptions<T = any> = {
  endpoint: string;
  method?: RequestVerb;
  needsAuth?: boolean;
  payload?: T;
};

interface UseRequestResult<TResult, TError = any> {
  isLoading: boolean;
  error: TError;
  data: TResult;
}

export const useRequest = async <TResult>(options: RequestOptions): Promise<UseRequestResult<TResult | undefined, any>> => {
  const api = import.meta.env.VITE_API;

  const isLoading = ref(false);
  const error = ref<unknown>(null);
  const data = ref<TResult | null>(null) as Ref<TResult | null>;

  const headers = {
    ...options.needsAuth === true
      ? { authorization: 'Bearer ' + localStorage.getItem('token') } : {}
  }

  try {
    const response = await axios.request({
      url: `${api}${options.endpoint}`,
      method: options.method ?? 'GET',
      headers,
      data: options.payload,
    });

    data.value = response.data;
  } catch (e) {
    error.value = e;
  } finally {
    isLoading.value = false;
  }

  return { isLoading: unwrap(isLoading) as boolean, error: unwrap(error), data: unwrap(data) as TResult };
}
