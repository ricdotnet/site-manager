export type ButtonType = 'button' | 'reset' | 'submit';
export type ButtonSize = 'sm' | 'md' | 'lg';
export type ButtonColor = 'primary' | 'gray' | 'danger';

export type InputComponent = {
  getValue: () => string;
  setValue: (value: string) => void;
  setError: (bool: boolean, message?: string) => void;
}

export type RegisterData = {
  username: string;
  email: string;
  password: string;
  password_confirm: string;
}

export type TSite = {
  ID: number;
  domain: string;
  config_name: string;
  enabled: boolean;
  has_ssl: boolean;
  created_at: string;
  updated_at: string;
  checked?: boolean;
}

export type TSitesResponse = {
  code: number;
  message: string;
  sites: TSite[];
};

export type TSIteResponse = {
  code: number;
  message: string;
  site: TSite;
}
