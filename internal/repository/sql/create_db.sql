CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
    email varchar(255),
    hash_password varchar(255),
    refresh_token varchar(255)
)