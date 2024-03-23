import * as monaco from 'monaco-editor';
import { ref } from 'vue';
import theme from '../monaco/theme.ts';
import { config, tokens } from '../monaco/syntax.ts';

export const useEditor = () => {
  const monacoRef = ref<HTMLDivElement | null>(null);

  monaco.languages.register({ id: 'nginx' });
  monaco.languages.setLanguageConfiguration('nginx', config);
  monaco.languages.setMonarchTokensProvider('nginx', tokens);
  monaco.editor.defineTheme('nginxTheme', theme);

  const buildEditor = (config: string) => {
    monaco.editor.create(monacoRef.value, {
      language: 'nginx',
      theme: 'vs-dark',
      tabSize: 2,
      lineNumbers: 'relative',
      scrollBeyondLastLine: false,
      lineNumbersMinChars: 4,
      padding: { top: 10, bottom: 10 },
      minimap: { enabled: false },
      value: config,
    });
  };

  return { monacoRef, buildEditor };
};
