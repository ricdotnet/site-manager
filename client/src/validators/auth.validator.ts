import { Validator } from "./index.ts";

export const usernameValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter a username',
  },
  match: {
    pattern: /^[A-z0-9]+$/g,
    message: 'Your username cannot contain invalid characters',
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
