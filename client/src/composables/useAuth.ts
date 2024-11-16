import type { AxiosError } from 'axios';
import type { RegisterData } from '@types';
import axios from 'axios';
import { ref } from 'vue';
import { useUserStore } from '@stores';

const useAuth = () => {
  const api = import.meta.env.VITE_API;
  const userStore = useUserStore();

  const login = async (email: string) => {
    const error = ref<unknown>();

    try {
      await axios.post(
        `${api}/user/sign-in`,
        {
          email,
        },
        { withCredentials: true },
      );
    } catch (err: unknown) {
      error.value = (err as AxiosError).response?.data;
    }

    return { error: error.value };
  };

  const register = async (registerData: RegisterData) => {
    const error = ref<unknown>();
    try {
      const { data } = await axios.post(`${api}/user/register`, registerData);

      console.log(data);
    } catch (err: unknown) {
      error.value = (err as AxiosError).response?.data;
    }

    return { error: error.value };
  };

  const verifyCode = async (code: string) => {
    const error = ref<unknown>();

    try {
      const { data } = await axios.post(
        `${api}/user/verify-code`,
        {
          code,
        },
        { withCredentials: true },
      );

      userStore.setIsAuthed(true);
      userStore.setUserId(data.id);
      userStore.setUsername(data.username);
    } catch (err: unknown) {
      error.value = (err as AxiosError).response?.data;
    }

    return { error: error.value };
  };

  return { login, register, verifyCode };
};

export { useAuth };
