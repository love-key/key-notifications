-- 000001_create_email_prefernces.up.sql

CREATE TABLE emailPreferences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    userId INT NOT NULL,
    category VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    isEnabled BOOLEAN DEFAULT TRUE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deletedAt TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
);

-- Optionally, create an index on deletedAt to optimize soft-deletes.
CREATE INDEX idx_email_preferences_deleted_at ON emailPreferences(deletedAt);
