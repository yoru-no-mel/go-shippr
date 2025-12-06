CREATE TYPE user_type AS ENUM ('FUJO', 'TACHI', 'NEKO');
ALTER TABLE users
ADD COLUMN type user_type DEFAULT 'FUJO' NOT NULL,
ADD COLUMN bio text,
ADD COLUMN avatar text;