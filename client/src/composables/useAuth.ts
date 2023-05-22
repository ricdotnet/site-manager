import { RegisterData } from "../types.ts";
import axios from "axios";
import { ref } from "vue";

const useAuth = () => {
  const api = import.meta.env.VITE_API;

  // @ts-ignore
  const login = async (username: string, password: string) => {
    const data = ref<any>();
    const error = ref<any>();

    try {
      const response = await axios.post(`${api}/user/login`, {
        username: username, password: password,
      });
      data.value = response.data;
    } catch (err: any) {
      error.value = err.response.data;
    }

    return { error: error.value, data: data.value };
  }

  const register = (data: RegisterData): void | {} => {
    console.log(data);
  }

  return { login, register }
}

export { useAuth };
