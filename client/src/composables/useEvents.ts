export const useEvents = () => {
  const events = new Map<string, Function>()

  const addEvent = (event: string, callback: Function) => {
    if (events.has(event)) {
      throw new Error(`Event ${event} already exists`)
    }

    events.set(event, callback)
  }

  const removeEvent = (event: string) => {
    if (!events.has(event)) {
      throw new Error(`Event ${event} does not exist`)
    }

    events.delete(event)
  }

  const emitEvent = (event: string, ...args: any) => {
    if (!events.has(event) || !events.get(event)) {
      throw new Error(`Event ${event} does not exist`)
    }

    events.get(event)!(...args);
  }

  return { addEvent, removeEvent, emitEvent };
}