import { Validator } from "./index.ts";

export const usernameValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter a username',
  },
  match: {
    pattern: /^[A-Za-z0-9_$]+$/g,
    message: 'Your username contains invalid characters (only letters, numbers, _ and $ are allowed)',
  },
  length: {
    minimum: 5,
    maximum: 30,
    message: 'Your username needs to be more than 5 characters and no more than 30',
  },
};

export const emailValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter an email',
  },
  match: {
    pattern: /^[A-Za-z0-9._%+-]+@[A-Za-z0-9-.]+\.[A-Za-z]{2,}$/g,
    message: 'The email entered is not a valid email address'
  }
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
