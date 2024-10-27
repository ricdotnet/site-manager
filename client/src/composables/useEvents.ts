const _events: { [key: string]: (args: unknown[]) => unknown } = {};

export const useEvents = () => {
  const on = (key: string, callback: (...args: unknown[]) => unknown) => {
    _events[key] = callback;
  };

  const off = (key: string): unknown => {
    if (!_events[key]) {
      console.error(`Event ${key} does not exist`);
      return;
    }

    // eslint-disable-next-line @typescript-eslint/no-dynamic-delete
    delete _events[key];
  };

  const emit = (key: string, ...args: unknown[]) => {
    if (!_events[key]) {
      console.error(`Event ${key} does not exist`);
      return;
    }

    const eventFn = _events[key];
    if (!eventFn) return {};
    return eventFn(args);
  };

  return { on, off, emit };
};
