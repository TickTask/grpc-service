-- Создаем таблицу пользователей
CREATE TABLE Users
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT, -- Автоинкрементируемый первичный ключ
    login         TEXT NOT NULL,                     -- Логин пользователя
    name          TEXT NOT NULL,                     -- Имя пользователя
    hash_password BLOB NOT NULL                      -- Хэшированный пароль пользователя
);

-- Создаем таблицу задач
CREATE TABLE Tasks
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT,                        -- Автоинкрементируемый первичный ключ
    title          TEXT NOT NULL,                                            -- Заголовок задачи
    body           TEXT,                                                     -- Описание задачи
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                      -- Дата создания задачи
    task_user_id   INTEGER,                                                  -- Ссылка на пользователя, назначившего задачу
    task_status_id INTEGER,                                                  -- Ссылка на статус задачи
    FOREIGN KEY (task_user_id) REFERENCES Users (id) ON DELETE SET NULL,     -- Внешний ключ на таблицу Users
    FOREIGN KEY (task_status_id) REFERENCES Statuses (id) ON DELETE SET NULL -- Внешний ключ на таблицу Statuses
);

-- Создаем таблицу статусов
CREATE TABLE Statuses
(
    id     INTEGER PRIMARY KEY AUTOINCREMENT, -- Автоинкрементируемый первичный ключ
    status TEXT NOT NULL                      -- Название статуса задачи
);

-- Создаем таблицу сессий
CREATE TABLE Sessions
(
    id              BLOB PRIMARY KEY,                                     -- UUID как бинарные данные (16 байт)
    refresh_token   TEXT    NOT NULL,                                     -- Токен обновления
    device_id       INTEGER NOT NULL,                                     -- ID устройства
    session_user_id INTEGER,                                              -- Ссылка на пользователя сессии
    FOREIGN KEY (session_user_id) REFERENCES Users (id) ON DELETE CASCADE -- Внешний ключ на таблицу Users
);