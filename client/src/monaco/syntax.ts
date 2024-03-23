const keywords = ['server', 'location', 'if', 'else', 'break', 'set', 'rewrite', 'return'];

const identifiers = [
  'include',
  'default_type',
  'listen',
  'server_name',
  'proxy_pass',
  'root',
  'ssl_certificate',
  'ssl_certificate_key',
  'proxy_http_version',
  'proxy_set_header',
];

const keywordsRegex = new RegExp(`(\\b)(${keywords.join('|')})(\\b)`);
const identifiersRegex = new RegExp(`(\\s)(${identifiers.join('|')})(\\s)`);

export const config = {
  comments: {
    lineComment: '#',
  },
  brackets: [
    ['{', '}'],
    ['[', ']'],
  ],
  autoClosingPairs: [
    { open: '{', close: '}' },
    { open: '[', close: ']' },
    { open: '"', close: '"' },
  ],
};

export const tokens = {
  tokenizer: {
    root: [
      [/#.*/, 'comment'],
      [keywordsRegex, 'keyword'],
      [/"(?:\\.|[^"\\])*?"/, 'string'],
      [identifiersRegex, ['white', 'type', 'white']],
    ],
  },
};
