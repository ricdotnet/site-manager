import * as monaco from 'monaco-editor';
import { ref } from 'vue';
import { config, tokens } from '../monaco/syntax.ts';
import theme from '../monaco/theme.ts';

export const useEditor = () => {
  const monacoRef = ref<HTMLDivElement | null>(null);

  monaco.languages.register({ id: 'nginx' });
  monaco.languages.setLanguageConfiguration('nginx', config as monaco.languages.LanguageConfiguration);
  monaco.languages.setMonarchTokensProvider('nginx', tokens as monaco.languages.IMonarchLanguage);
  monaco.editor.defineTheme('nginxTheme', theme as monaco.editor.IStandaloneThemeData);

  const buildEditor = (config: string) => {
    if (!monacoRef.value) return;

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
      automaticLayout: true,
    });
  };

  return { monacoRef, buildEditor };
};
