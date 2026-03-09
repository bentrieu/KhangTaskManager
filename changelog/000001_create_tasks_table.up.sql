CREATE TYPE progress_enum
AS
ENUM('done', 'in progress', 'todo');

CREATE TABLE IF NOT EXISTS tasks(
    task_id serial PRIMARY KEY,
    descript VARCHAR(255) NOT NULL,
    progress PROGRESS_ENUM NOT NULL,
    created_at timestamp default current_timestamp
);
