CREATE TABLE IF NOT EXISTS problem_submits (
    id SERIAL NOT NULL,
    problem_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INTEGER NOT NULL,
    PRIMARY KEY (id)
);