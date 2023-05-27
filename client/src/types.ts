export type ButtonType = 'button' | 'reset' | 'submit';
export type ButtonSize = 'sm' | 'md' | 'lg';
export type ButtonColor = 'primary' | 'gray' | 'danger';

export type InputComponent = {
  getValue: () => string;
  setError: (bool: boolean, message?: string) => void;
}

export type RegisterData = {
  username: string;
  email: string;
  password: string;
  password_confirm: string;
}
