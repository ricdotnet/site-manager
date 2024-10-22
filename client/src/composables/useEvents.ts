const events = new Map<string, (args: unknown[]) => void>();

export const useEvents = () => {
  const addEvent = (event: string, callback: (args: unknown[]) => void) => {
    if (events.has(event)) {
      console.warn(
        `Event ${event} already exists. Clean up the event on unmount to register a new event with the same key.`,
      );
    }

    events.set(event, callback);
  };

  const removeEvent = (event: keyof typeof events.keys) => {
    if (!events.has(event)) {
      console.error(`Event ${event} does not exist`);
    }

    events.delete(event);
  };

  const emitEvent = (event: keyof typeof events.keys, ...args: unknown[]) => {
    if (!events.has(event)) {
      console.error(`Event ${event} does not exist`);
    }

    const eventFn = events.get(event);
    if (!eventFn) return;
    eventFn(args);
  };

  return { addEvent, removeEvent, emitEvent };
};
