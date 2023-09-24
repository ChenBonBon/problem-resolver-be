CREATE TABLE IF NOT EXISTS problems (
    id SERIAL NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    difficulty VARCHAR(16) NOT NULL,
    status VARCHAR(16) NOT NULL,
    answer TEXT,
    types INTEGER[],
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMPTZ,
    updated_by INTEGER,
    PRIMARY KEY (id)
);