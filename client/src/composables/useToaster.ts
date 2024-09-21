import { ref } from 'vue';

type Toast = {
    id: number;
    message: string;
    type: 'success' | 'error' | 'info' | 'warning';
};

const toasts = ref<Toast[]>([]);

export const useToaster = () => {
    const addToast = (message: string, type: Toast['type']) => {
        const id = Date.now();
        toasts.value.push({ id, message, type });
        setTimeout(() => {
            removeToast(id);
        }, 10000); // to be an env variable
    };

    const removeToast = (id: number) => {
        toasts.value = toasts.value.filter((toast) => toast.id !== id);
    };

    return { toasts, addToast, removeToast };
};