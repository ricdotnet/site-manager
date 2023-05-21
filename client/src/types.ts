export type ButtonType = 'button' | 'reset' | 'submit';
export type ButtonSize = 'sm' | 'md' | 'lg';
export type ButtonColor = 'primary' | 'gray' | 'danger';

export type InputComponent = {
  getValue: () => void;
  setError: () => void;
}
