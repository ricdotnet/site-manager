export const useDebounce = () => {
  let timer: NodeJS.Timeout | undefined;

  return (fn: (...args: any[]) => void, delay = 500) => {
    if (timer) clearTimeout(timer);
    timer = setTimeout(fn, delay);
  };
};
