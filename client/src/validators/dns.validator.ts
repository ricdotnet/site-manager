import type { Validator } from '@validators';

export const hostnameValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter a host name',
  },
};

export const ipv4Validator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter an ip address',
  },
  match: {
    pattern: /^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$/g,
    message: 'The ip address is not valid',
  },
};

export const ipv6Validator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter an ip address',
  },
  match: {
    pattern:
      /(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))/g,
    message: 'The ip address is not valid',
  },
};

export const aliasValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter an alias',
  },
};

export const ttlValidator: Validator = {
  presence: {
    allowEmpty: false,
    message: 'Please enter a TTL',
  },
  match: {
    pattern: /^\d+$/g,
    message: 'The TTL must be a number in seconds',
  },
};
