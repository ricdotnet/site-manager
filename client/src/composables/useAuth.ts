import type { VerifyCodeResponse} from '@types';
import { useRequest } from '@composables';
import { useUserStore } from '@stores';

const useAuth = () => {
  const userStore = useUserStore();

  const login = async (email: string) => {
    const { error } = await useRequest({
      endpoint: '/user/sign-in',
      method: 'POST',
      payload: {
        email,
      }
    });

    return { error };
  };

  const verifyCode = async (code: string) => {
    const { error, data } = await useRequest<VerifyCodeResponse>({
      endpoint: '/user/verify-code',
      method: 'POST',
      payload: { code },
    });

    if (error) {
      return { error };
    }

    if (data) {
      userStore.setIsAuthed(true);
      userStore.setUserId(data.id);
      userStore.setUsername(data.username);
    }

    return { error };
  };

  return { login, verifyCode };
};

export { useAuth };
