export * from './auth.validator.ts';
export * from './site.validator.ts';
export * from './apikey.validator.ts';
export * from './dns.validator.ts';

type Presence = {
  allowEmpty: boolean;
  message?: string;
};

type Match = {
  pattern: RegExp;
  message?: string;
};

type Length = {
  minimum?: number;
  maximum?: number;
  message?: string;
} & ({ minimum: number } | { maximum: number });

export type Validator = {
  presence?: Presence;
  match?: Match;
  length?: Length;
};

export function validate(validator: Validator | (() => void), data: string) {
  if (validator instanceof Function) {
    return console.log('run a custom validator....');
  }

  if (validator.presence && !validator.presence.allowEmpty && !data) {
    if (!validator.presence.message) return 'your data cannot be empty';
    return validator.presence.message;
  }

  if (validator.match?.pattern) {
    const pattern = new RegExp(validator.match.pattern);
    if (!pattern.test(data)) {
      if (!validator.match.message)
        return 'your data contains invalid characters';
      return validator.match.message;
    }
  }

  if (validator.length) {
    if (validator.length.minimum && data.length < validator.length.minimum) {
      if (!validator.length.message)
        return `your data cannot be of length shorter than ${validator.length.minimum} characters`;
      return validator.length.message;
    }

    if (validator.length.maximum && data.length > validator.length.maximum) {
      if (!validator.length.message)
        return `your data cannot be longer than ${validator.length.maximum} characters`;
      return validator.length.message;
    }
  }
}
