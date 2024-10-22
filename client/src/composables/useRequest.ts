import { unwrap } from '@utils';
import axios from 'axios';
import { type Ref, ref } from 'vue';

type RequestVerb = 'GET' | 'POST' | 'PATCH' | 'PUT' | 'DELETE';

type RequestOptions<T = unknown> = {
  endpoint: string;
  method?: RequestVerb;
  needsAuth?: boolean;
  payload?: T;
  useCache?: boolean;
};

type RequestCache<T = unknown> = {
  duration: number;
  data: T;
};

interface UseRequestResult<TResult, TError = unknown> {
  error: TError;
  data: TResult;
}

const requestCache = new Map<string, RequestCache>();

export const useRequest = async <TResult>(
  options: RequestOptions,
): Promise<UseRequestResult<TResult | undefined, unknown>> => {
  const api = import.meta.env.VITE_API;

  const error = ref<unknown>(null);
  const data = ref<TResult | null>(null) as Ref<TResult | null>;

  const headers = {
    ...(options.needsAuth === true ? { authorization: `Bearer ${localStorage.getItem('token')}` } : {}),
  };

  if (options.useCache) {
    const cached = requestCache.get(options.endpoint);
    if (cached && Date.now() - cached.duration < 30000) {
      data.value = cached.data as TResult;
      return { error: unwrap(error), data: unwrap(data) as TResult };
    }
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
  }

  if (options.useCache && !error.value) {
    requestCache.set(options.endpoint, {
      duration: Date.now() + 30000,
      data: data.value,
    });
  }

  return { error: unwrap(error), data: unwrap(data) as TResult };
};
