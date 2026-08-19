CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username VARCHAR(128) NOT NULL,
  email VARCHAR(256) NOT NULL UNIQUE,
  password_hash VARCHAR(256) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
);

CREATE INDEX IF NOT EXISTS idx_user_unique_email ON users(email);

CREATE TABLE IF NOT EXISTS tasks (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  title VARCHAR(128) NOT NULL,
  is_active BOOLEAN DEFAULT TRUE,
  description VARCHAR(512) NOT NULL,
  leetcode_url VARCHAR(256) NOT NULL,
  notification_count INTEGER DEFAULT 0,
  next_notification TIMESTAMPTZ NOT NULL DEFAULT NOW() + INTERVAL '1 day',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS notification_preferences (
  user_id UUID PRIMARY KEY,
  telegram_enabled BOOLEAN NOT NULL DEFAULT FALSE,
  telegram_chat_id BIGINT NOT NULL DEFAULT 0,

  FOREIGN KEY (user_id) REFERENCES users(id)
);