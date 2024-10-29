export const useDebounce = () => {
  let timer: number | undefined;

  return (fn: (...args: unknown[]) => void, delay = 500) => {
    if (timer) clearTimeout(timer);
    // @ts-expect-error node uses NodeJS timer but the browser uses a number for the timer
    timer = setTimeout(fn, delay);
  };
};
