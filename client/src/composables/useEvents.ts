export const useEvents = () => {
  const events = new Map<string, (args: unknown[]) => void>();

  const addEvent = (event: string, callback: (args: unknown[]) => void) => {
    if (events.has(event)) {
      throw new Error(`Event ${event} already exists`);
    }

    events.set(event, callback);
  };

  const removeEvent = (event: string) => {
    if (!events.has(event)) {
      throw new Error(`Event ${event} does not exist`);
    }

    events.delete(event);
  };

  const emitEvent = (event: string, ...args: unknown[]) => {
    if (!events.has(event)) {
      throw new Error(`Event ${event} does not exist`);
    }

    const eventFn = events.get(event);
    if (!eventFn) return;
    eventFn(args);
  };

  return { addEvent, removeEvent, emitEvent };
};
