export type ButtonType = 'button' | 'reset' | 'submit';
export type ButtonSize = 'sm' | 'md' | 'lg';
export type ButtonColor = 'primary' | 'gray' | 'danger' | 'icon';

// components
export type InputComponent = {
  getValue: () => string;
  setValue: (value: string) => void;
  setError: (bool: boolean, message?: string) => void;
  hasError: boolean;
};

export type FormComponent = {
  setFormError: (error: string | undefined) => void;
  clearFormError: () => void;
};

// responses
export type BaseResponse = {
  code: number;
  message?: string;
  message_code?: string;
};

export type TSitesResponse = BaseResponse & {
  sites: TSite[];
};

export type TSiteResponse = BaseResponse & {
  config: string;
  site: TSite;
};

export type TDNSRecordsResponse = BaseResponse & {
  records: {
    records: TDNSRecord[];
  };
};

export type TActiveSessionsResponse = BaseResponse & {
  active_sessions: TSession[];
};

// more
export type RegisterData = {
  username: string;
  email: string;
  password: string;
  password_confirm: string;
};

export type TSite = {
  id: number;
  domain: string;
  config_name: string;
  enabled: boolean;
  has_ssl: boolean;
  created_at: string;
  updated_at: string;
  config?: string;
  checked?: boolean;
  config_only?: boolean;
};

export type TToast = {
  id: number;
  title?: string;
  message: string;
  type: 'success' | 'error' | 'info' | 'warning';
};

export type TApiKey = {
  id: number;
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

export type TDNSRecord = {
  id: number;
  host: string;
  value: string;
  ttl: string;
  status: string;
};

export type TDNSRecordFormProps = {
  onClearFormError: () => void;
};

export type TDNSRecordFormProcess = {
  hasError?: boolean;
  error?: string;
  data?: {
    host: string;
    value: string;
    ttl: string;
  };
};

export type TSession = {
  id: number;
  user_agent: string;
  last_active: string;
  this_device: boolean;
  ip_address: string;
  parsed_user_agent: {
    name: string;
    version: string;
    os: string;
    os_version: string;
    is_desktop: boolean;
    is_mobile: boolean;
    is_tablet: boolean;
  };
};
