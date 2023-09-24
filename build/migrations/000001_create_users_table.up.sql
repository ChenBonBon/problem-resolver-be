CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255),
    source VARCHAR(255),
    status VARCHAR(255) NOT NULL,
    role INTEGER NOT NULL,
    avatar TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    PRIMARY KEY (id)
);

INSERT INTO users(name, email, password, source, status, role) VALUES('BonBon','per_cherry@163.com',md5('chenroc1001'),'initial','enabled',1), ('Jimmy','jimmy.jiang@basebit.ai',md5('jimmy.jiang'),'initial','enabled',1);