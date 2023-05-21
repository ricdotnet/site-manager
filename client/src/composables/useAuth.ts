const useAuth = () => {
  // const api = import.meta.env.VITE_API;

  const login = (username: string, password: string) => {
    console.log(username);
    console.log(password);
  }

  const register = (data: {}) => {
    console.log(data);
  }

  return { login, register }
}

export { useAuth };
