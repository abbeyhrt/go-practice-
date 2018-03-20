CREATE TABLE homes
(
  id SERIAL PRIMARY KEY,
  object_id UUID DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL,
  description TEXT NOT NULL,
  avatar_url VARCHAR(256),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  UNIQUE(object_id)
);