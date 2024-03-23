import { Validator } from './index.ts';
import { messages } from '../utils';

export const usernameValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: messages.user.missing_username,
  },
  match: {
    pattern: /^[A-Za-z0-9_$]+$/g,
    message: messages.user.invalid_username,
  },
  length: {
    minimum: 5,
    maximum: 30,
    message: messages.user.invalid_username_length,
  },
};

export const emailValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: messages.user.missing_email,
  },
  match: {
    pattern: /^[A-Za-z0-9._%+-]+@[A-Za-z0-9-.]+\.[A-Za-z]{2,}$/g,
    message: messages.user.invalid_email,
  },
};

export const passwordValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: messages.user.missing_password,
  },
  length: {
    minimum: 9,
    message: messages.user.invalid_password_length,
  },
};

export const passwordConfirmValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: messages.user.missing_password_confirm,
  },
};
