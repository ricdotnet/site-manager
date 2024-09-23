import type { TToast } from '@types';
import { ref } from 'vue';

const toasts = ref<TToast[]>([]);

export const useToaster = () => {
  function addToast(type: TToast['type'], message: string): void;
  function addToast(type: TToast['type'], message: string, title?: string) {
    const id = Date.now();
    if (title) {
      toasts.value.push({ id, title, message, type });
    } else {
      toasts.value.push({ id, message, type });
    }
    setTimeout(() => {
      removeToast(id);
    }, 10000); // to be an env variable
  }

  const removeToast = (id: number) => {
    toasts.value = toasts.value.filter((toast) => toast.id !== id);
  };

  return { toasts, addToast, removeToast };
};
