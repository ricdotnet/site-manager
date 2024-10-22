export type ButtonType = 'button' | 'reset' | 'submit';
export type ButtonSize = 'sm' | 'md' | 'lg';
export type ButtonColor = 'primary' | 'gray' | 'danger' | 'icon';

export type InputComponent = {
  getValue: () => string;
  setValue: (value: string) => void;
  setError: (bool: boolean, message?: string) => void;
  hasError: boolean;
};

export type RegisterData = {
  username: string;
  email: string;
  password: string;
  password_confirm: string;
};

export type TSite = {
  ID: number;
  domain: string;
  config_name: string;
  enabled: boolean;
  has_ssl: boolean;
  created_at: string;
  updated_at: string;
  config?: string;
  checked?: boolean;
};

export type TSitesResponse = {
  code: number;
  message: string;
  sites: TSite[];
};

export type TSIteResponse = {
  code: number;
  message: string;
  site: TSite;
};

export type TToast = {
  id: number;
  title?: string;
  message: string;
  type: 'success' | 'error' | 'info' | 'warning';
};

export type TApiKey = {
  ID: number;
  key: string;
  value: string;
  is_api_key: boolean;
  created_at: string;
  updated_at: string;
};

export type TDomain = {
  order_id: number;
  name: string;
  created_at: string;
  renewal_at: string;
};
