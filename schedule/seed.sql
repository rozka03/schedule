INSERT INTO groups (name, year) VALUES
    ('ИС-123/9', 2023),
    ('ИС-223/9', 2023);

INSERT INTO teachers (full_name, email) VALUES
    ('Самарин Игорь Владимирович',  'samarin@mail.ru'),
    ('Лямина Ксения Сергеевна', 'lyamina@mail.ru');

INSERT INTO subjects (name) VALUES
    ('Базы данных'),
    ('Веб-разработка'),
    ('Алгоритмы');

INSERT INTO lessons (group_id, teacher_id, subject_id, room, starts_at, ends_at, status) VALUES
    (1, 1, 1, 'ауд. 410', '2026-09-28 09:00+05', '2026-09-28 10:30+05', 'planned'),
    (1, 2, 2, 'ауд. 415', '2026-09-28 10:40+05', '2026-09-28 12:10+05', 'planned'),
    (2, 1, 1, 'ауд. 205', '2026-09-29 09:00+05', '2026-09-29 10:30+05', 'planned'),
    (2, 2, 3, 'ауд. 412', '2026-09-29 10:40+05', '2026-09-29 12:10+05', 'done');