CREATE TABLE IF NOT EXISTS auth_data (
    token CHAR(50) PRIMARY KEY,
    login TEXT UNIQUE NOT NULL,
    password TEXT UNIQUE NOT NULL,
    user_id INT REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT UNIQUE NOT NULL,
    vk_link TEXT UNIQUE,
    tg_link TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS contacts (
    user_id INT REFERENCES users(id),
    phone TEXT NOT NULL,
    vk TEXT,
    telegram TEXT,
    email TEXT
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),

    -- если нет user_id указываем контакты
    phone TEXT NOT NULL,
    vk TEXT,
    telegram TEXT,
    email TEXT
);
