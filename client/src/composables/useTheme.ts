import { ref } from 'vue';

const useTheme = (): { toggle: () => void } => {
  const currentTheme = ref<'light' | 'dark'>();

  const setDark = (set: boolean): void => {
    if (set) {
      document.documentElement.classList.add('dark');
      currentTheme.value = 'dark';
      localStorage.theme = 'dark';
    } else {
      document.documentElement.classList.remove('dark');
      currentTheme.value = 'light';
      localStorage.theme = 'light';
    }
  };

  if (localStorage.theme === 'dark' || (!('theme' in localStorage) &&
    window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    setDark(true);
  } else {
    setDark(false);
  }

  const toggle = (): void => {
    currentTheme.value === 'dark' ? setDark(false) : setDark(true);
  };

  return {
    toggle
  };
};

export { useTheme };
