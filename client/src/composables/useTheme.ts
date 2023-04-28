import { Ref, ref } from 'vue';

type ThemeString = 'light' | 'dark';

type TUseTheme = {
  toggleTheme: () => void;
  currentTheme: Ref<ThemeString>;
}

const useTheme = (): TUseTheme => {
  // prefer light first
  const currentTheme = ref<ThemeString>('light');

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

  const toggleTheme = (): void => {
    currentTheme.value === 'dark' ? setDark(false) : setDark(true);
  };

  return {
    toggleTheme,
    currentTheme,
  };
};

export { useTheme };
