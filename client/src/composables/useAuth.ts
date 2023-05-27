import { RegisterData } from "../types.ts";
import axios from "axios";
import { ref } from "vue";
import { useUserStore } from "../stores/user.store.ts";

const useAuth = () => {
  const api = import.meta.env.VITE_API;
  const userStore = useUserStore();

  // @ts-ignore
  const login = async (username: string, password: string) => {
    const error = ref<any>();

    try {
      const { data } = await axios.post(`${api}/user/login`, {
        username: username, password: password,
      });

      localStorage.setItem("token", data.token);

      userStore.setIsAuthed(true);
      userStore.setUsername(data.username);
    } catch (err: any) {
      error.value = err.response.data;
    }

    return { error: error.value };
  }

  const register = async (registerData: RegisterData) => {
    const error = ref<any>();
    console.log(registerData);
    return { error: 'hi' };
    try {
      const { data } = await axios.post(`${api}/user/register`, registerData);

      console.log(data);
    } catch (err: any) {
      error.value = err.response.data;
    }

    return { error: error.value };
  }

  return { login, register }
}

export { useAuth };
