CREATE TYPE oauth_provider AS ENUM ('facebook', 'google');

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  object_id UUID DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL,
  email VARCHAR(64) NOT NULL,
  avatar_url VARCHAR(256),
  provider oauth_provider NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  UNIQUE(object_id)
);