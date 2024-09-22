import type { Validator } from './index.ts';

export const apiKeyValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Enter a key for your third party API',
  },
};

export const apiKeyValueValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Enter a value for your third party API key',
  },
};
