import axios from 'axios';
import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useUserStore = defineStore('user', () => {
  const api = import.meta.env.VITE_API;

  const isAuthed = ref(false);
  const userId = ref('');
  const username = ref('');
  const isAdmin = ref(false);

  const tokenAuth = async () => {
    try {
      const response = await axios.get(`${api}/user/auth`, {
        withCredentials: true,
      });

      setIsAuthed(true);
      setUserId(response.data.id);
      setUsername(response.data.username);
    } catch (_) {
      localStorage.removeItem('token');
    }
  };

  const logout = async () => {
    await axios.get(`${api}/user/logout`, {
      withCredentials: true,
    });
  };

  const setIsAuthed = (v: boolean) => {
    isAuthed.value = v;
  };

  const setUserId = (v: string) => {
    userId.value = v;
  };

  const setUsername = (v: string) => {
    username.value = v;
  };

  return {
    isAuthed,
    userId,
    username,
    isAdmin,
    tokenAuth,
    setIsAuthed,
    setUserId,
    setUsername,
    logout,
  };
});
