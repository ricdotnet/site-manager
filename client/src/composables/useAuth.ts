import { RegisterData } from "../types.ts";

const useAuth = () => {
  // const api = import.meta.env.VITE_API;

  // @ts-ignore
  const login = (username: string, password: string) => {
  }

  const register = (data: RegisterData): void | {} => {
    console.log(data);
  }

  return { login, register }
}

export { useAuth };
