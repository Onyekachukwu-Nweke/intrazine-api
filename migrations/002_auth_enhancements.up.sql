-- Create enum for user roles
CREATE TYPE user_role AS ENUM ('reader', 'author', 'editor', 'admin');

ALTER TABLE users
    ADD COLUMN role user_role DEFAULT 'reader',
    ADD COLUMN two_factor_secret VARCHAR(32),
    ADD COLUMN is_2fa_enabled BOOLEAN DEFAULT false,
    ADD COLUMN login_attempts INTEGER DEFAULT 0,
    ADD COLUMN locked_until TIMESTAMP WITH TIME ZONE;

CREATE TABLE sessions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_sessions_token ON sessions(token);
CREATE INDEX idx_sessions_user_id ON sessions(user_id);

-- Add comments explaining roles
COMMENT ON TYPE user_role IS 'Blog user roles:
- reader: Can read posts and comment
- author: Can create and manage their own posts
- editor: Can edit and manage all posts
- admin: Full system access'; 