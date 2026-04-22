INSERT INTO tasks (descript, progress, created_at)
SELECT
    'Task #' || i || ' - something important',
    (ARRAY[
        'todo'::progress_enum,
        'in progress'::progress_enum,
        'done'::progress_enum
    ])[floor(random()*3 + 1)],
    NOW() - (random() * interval '7 days')
FROM generate_series(1, 1000) AS s(i);
