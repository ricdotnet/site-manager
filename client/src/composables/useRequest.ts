import { type Ref, ref } from 'vue';
import axios from 'axios';
import { unwrap } from '@utils';

type RequestVerb = 'GET' | 'POST' | 'PATCH' | 'PUT' | 'DELETE';

type RequestOptions<T = unknown> = {
  endpoint: string;
  method?: RequestVerb;
  payload?: T;
  useCache?: boolean;
};

type RequestCache<T = unknown> = {
  duration: number;
  data: T;
};

interface UseRequestResult<TResult, TError = string | undefined> {
  error: TError;
  data: TResult;
  isLoading: boolean;
}

const requestCache = new Map<string, RequestCache>();

export const useRequest = async <TResult>(
  options: RequestOptions,
): Promise<UseRequestResult<TResult | undefined>> => {
  const api = import.meta.env.VITE_API;

  const isLoading = ref(true);
  const error = ref<string | undefined>();
  const data = ref<TResult | null>(null) as Ref<TResult | null>;

  if (options.useCache) {
    const cached = requestCache.get(options.endpoint);
    if (cached && Date.now() - cached.duration < 30000) {
      data.value = cached.data as TResult;
      isLoading.value = false;
      return {
        error: unwrap(error) as unknown as string,
        data: unwrap(data) as TResult,
        isLoading: unwrap(isLoading) as boolean,
      };
    }
  }

  try {
    const response = await axios.request({
      url: `${api}${options.endpoint}`,
      method: options.method ?? 'GET',
      data: options.payload,
      withCredentials: true,
    });

    data.value = response.data;
    isLoading.value = false;
  } catch (e) {
    if (axios.isAxiosError(e) && e.response) {
      error.value = e.response.data.message_code ?? 'generic_error';
    } else if (e instanceof Error) {
      error.value = e.message ?? 'generic_error';
    } else {
      error.value = 'Something went wrong';
    }
    isLoading.value = false;
  }

  if (options.useCache && !error.value) {
    requestCache.set(options.endpoint, {
      duration: Date.now() + 30000,
      data: data.value,
    });
  }

  return {
    error: unwrap(error) as unknown as string,
    data: unwrap(data) as TResult,
    isLoading: unwrap(isLoading) as boolean,
  };
};
