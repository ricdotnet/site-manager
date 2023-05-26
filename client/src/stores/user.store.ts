import { defineStore } from "pinia";
import { ref } from "vue";
import axios from "axios";

export const useUserStore = defineStore('user', () => {
  const api = import.meta.env.VITE_API;

  const isAuthed = ref(false);
  const username = ref('');
  const isAdmin = ref(false);

  const tokenAuth = async () => {
    try {
      const response = await axios.get(`${api}/user/auth`, {
        headers: {
          'authorization': `Bearer ${localStorage.getItem('token')}`
        }
      });

      setIsAuthed(true);
      setUsername(response.data.username);
    } catch (error) {
      localStorage.removeItem('token');
    }
  }

  const setIsAuthed = (v: boolean) => {
    isAuthed.value = v;
  }

  const setUsername = (v: string) => {
    username.value = v;
  }

  return { isAuthed, username, isAdmin, tokenAuth, setIsAuthed, setUsername };
});
