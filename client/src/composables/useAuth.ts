import type { AxiosError } from 'axios';
import type { RegisterData } from '@types';
import axios from 'axios';
import { ref } from 'vue';
import { useUserStore } from '@stores';

const useAuth = () => {
  const api = import.meta.env.VITE_API;
  const userStore = useUserStore();

  const login = async (username: string, password: string) => {
    const error = ref<unknown>();

    try {
      const { data } = await axios.post(
        `${api}/user/login`,
        {
          username: username,
          password: password,
        },
        { withCredentials: true },
      );

      // localStorage.setItem('token', data.token);

      userStore.setIsAuthed(true);
      userStore.setUserId(data.id);
      userStore.setUsername(data.username);
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

  return { login, register };
};

export { useAuth };
