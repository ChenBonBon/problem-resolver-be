CREATE TABLE IF NOT EXISTS problem_examples(
    id SERIAL NOT NULL,
    problem_id INTEGER NOT NULL,
    input TEXT NOT NULL,
    output TEXT NOT NULL,
    explanation TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMPTZ,
    updated_by INTEGER,
    PRIMARY KEY (id)
)