import { Validator } from "./index.ts";

export const usernameValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter a username',
  },
  match: {
    pattern: /^[A-Za-z0-9_$]+$/g,
    message: 'Your username cannot contain invalid characters (only letters, numbers, _ and $ are allowed)',
  },
  length: {
    minimum: 5,
    message: 'Your username needs to be longer than or 5 characters',
  },
};

export const emailValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter an email',
  },
};

export const passwordValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter a password',
  },
  length: {
    minimum: 9,
    message: 'Your password needs to be longer than 8 characters',
  },
};

export const passwordConfirmValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please confirm your password',
  },
};
