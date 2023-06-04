export const messages = {
  user: {
    username_not_found: 'A user with that username does not exit',
    missing_username: 'Username cannot be missing',
    username_exists: 'A user with that username already exists',
    invalid_username_length: 'Your username needs to be more than 5 characters and no more than 30',
    invalid_username: 'Your username contains invalid characters (only letters, numbers, _ and $ are allowed)',
    missing_email: 'Email address cannot be missing',
    email_exists: 'A user with that email already exists',
    invalid_email: 'The email entered is not a valid email address',
    incorrect_password: 'You entered the wrong password',
    missing_password: 'Password cannot be missing',
    passwords_not_match: 'The passwords do not match',
    invalid_password_length: 'Your password needs to be longer than 8 characters',
    missing_password_confirm: 'Please confirm your password',
    login_success: 'User logged in successfully',
    register_success: 'New user registered successfully',
    valid_token: 'A valid token was used to authorize',
  },
  site: {
    site_create_success: 'The new site was created with success',
    site_create_error: 'We were unable to create a new site',
  },
  generic: {
    not_implemented: 'This functionality is not fully working yet',
  },
} as const;
