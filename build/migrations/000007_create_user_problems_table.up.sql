CREATE TABLE IF NOT EXISTS user_problems(
    id SERIAL NOT NULL,
    user_id INTEGER NOT NULL,
    problem_id INTEGER NOT NULL,
    status VARCHAR(16) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    PRIMARY KEY (id)
);