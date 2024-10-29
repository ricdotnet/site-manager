export { messages } from './messages.ts';

import { type Ref, isRef, ref } from 'vue';

type RefValue = string | number | object | boolean;

export function wrap(value: RefValue): Ref {
  if (isRef(value)) throw new Error('You cannot wrap a Ref');
  return ref<RefValue>(value);
}

export function unwrap<T extends RefValue>(ref: Ref | T): unknown {
  if (typeof ref === 'undefined') return null;
  return isRef(ref) ? ref.value : ref;
}

export function formatDate(date: string): string {
  return new Date(date).toDateString();
}

export const sleep = (ms: number) =>
  new Promise((r) => setTimeout(() => r(true), ms));
