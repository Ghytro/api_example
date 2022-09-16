CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    image TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS liquids (
    id SERIAL PRIMARY KEY,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    image TEXT NOT NULL,
    availability BOOLEAN NOT NULL,
    brief_desc TEXT NOT NULL,
    desc TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS strenghts (
    liquid_id INT REFERENCES liquids(id) ON DELETE CASCADE,
    strength TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS volumes (
    liquid_id INT REFERENCES liquids(id) ON DELETE CASCADE,
    volume INT NOT NULL
);

CREATE TABLE IF NOT EXISTS doppings (
    liquid_id INT REFERENCES liquids(id) ON DELETE CASCADE,
    dopping TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS prices (
    liquid_id INT REFERENCES liquids(id) ON DELETE CASCADE,
    roubles INT NOT NULL,
    cents INT NOT NULL,
    strength TEXT NOT NULL,
    volume INT NOT NULL,
    dopping TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
    liquid_id INT REFERENCES liquids(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id),
    text TEXT,
    rate INT NOT NULL
);
