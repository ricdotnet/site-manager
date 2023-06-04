import { Validator } from "./index.ts";

export const domainValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter a domain',
  },
  match: {
    pattern: /^[A-Za-z0-9.]+$/g,
    message: 'The domain is not a valid domain',
  },
};

export const configValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter a config filename',
  },
  match: {
    pattern: /^[A-Za-z0-9.]+$/g,
    message: 'The filename is not a valid filename',
  },
};
