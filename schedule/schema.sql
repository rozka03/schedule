-- Группы студентов
CREATE TABLE IF NOT EXISTS groups (
    id    BIGSERIAL PRIMARY KEY,
    name  TEXT NOT NULL UNIQUE,          -- "ПИ-21"
    year  INT  NOT NULL                  -- год набора
);

-- Преподаватели
CREATE TABLE IF NOT EXISTS teachers (
    id         BIGSERIAL PRIMARY KEY,
    full_name  TEXT NOT NULL,
    email      TEXT NOT NULL UNIQUE
);

-- Предметы
CREATE TABLE IF NOT EXISTS subjects (
    id    BIGSERIAL PRIMARY KEY,
    name  TEXT NOT NULL UNIQUE           -- "Базы данных"
);

-- Занятия = расписание
CREATE TABLE IF NOT EXISTS lessons (
    id          BIGSERIAL   PRIMARY KEY,
    group_id    BIGINT      NOT NULL REFERENCES groups(id)   ON DELETE CASCADE,
    teacher_id  BIGINT      NOT NULL REFERENCES teachers(id) ON DELETE RESTRICT,
    subject_id  BIGINT      NOT NULL REFERENCES subjects(id) ON DELETE RESTRICT,
    room        TEXT        NOT NULL,                         -- "ауд. 312"
    starts_at   TIMESTAMPTZ NOT NULL,
    ends_at     TIMESTAMPTZ NOT NULL,
    status      TEXT        NOT NULL DEFAULT 'planned'
                CHECK (status IN ('planned', 'done', 'cancelled')),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (ends_at > starts_at)
);

-- Индексы на внешние ключи — без них джойны и каскады тормозят
CREATE INDEX IF NOT EXISTS idx_lessons_group_id   ON lessons(group_id);
CREATE INDEX IF NOT EXISTS idx_lessons_teacher_id ON lessons(teacher_id);
CREATE INDEX IF NOT EXISTS idx_lessons_subject_id ON lessons(subject_id);

-- Индекс под частый запрос "расписание группы на дату"
CREATE INDEX IF NOT EXISTS idx_lessons_starts_at  ON lessons(starts_at);