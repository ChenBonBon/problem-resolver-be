CREATE TABLE IF NOT EXISTS problem_types (
    id SERIAL NOT NULL,
    name VARCHAR(63) NOT NULL,
    status VARCHAR(16) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMPTZ,
    updated_by INTEGER,
    PRIMARY KEY (id)
);

INSERT INTO problem_types(name, status, created_at, created_by) VALUES('其他', 'enabled', NOW(), 1), ('前端', 'enabled', NOW(), 1), ('后端', 'enabled', NOW(), 1), ('算法', 'enabled', NOW(), 1), ('数据库', 'enabled', NOW(), 1);
