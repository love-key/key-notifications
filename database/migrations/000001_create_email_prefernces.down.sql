-- 000001_create_email_prefernces.down.sql

DROP INDEX IF EXISTS idx_email_preferences_deleted_at;

DROP TABLE IF EXISTS emailPreferences;
