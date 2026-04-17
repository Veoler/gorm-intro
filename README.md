# GORM: базовое знакомство

Полезно прочитать перед началом работы:

- Подключение к базе данных (PostgreSQL): https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
- Модели в GORM: https://gorm.io/docs/models.html
- Миграции (AutoMigrate): https://gorm.io/docs/migration.html#Auto-Migration
- CRUD (создание/чтение/обновление/удаление): https://gorm.io/docs/create.html, https://gorm.io/docs/query.html, https://gorm.io/docs/update.html, https://gorm.io/docs/delete.html

## Релиз 0 — Установка GORM

Установи GORM и драйвер для PostgreSQL в текущий проект Go следуя документации https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL

Проверка:

- Убедись, что зависимости скачались и проект компилируется.

## Релиз 1 — Модель Student

Создай структуру `Student`, в которой есть поле `ID` (первичный ключ). Добавь 1–2 любых простых поля (например, имя или возраст), чтобы было с чем работать далее.

Документация:

- Модели: https://gorm.io/docs/models.html

Проверка:

- Просто убедись, что структура корректно описана и проект компилируется.

## Релиз 2 — Переход на `gorm.Model`

Перепиши модель `Student`, встроив в неё `gorm.Model`. Обрати внимание, что `gorm.Model` включает `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`.

Документация:

- https://gorm.io/docs/models.html#gorm.Model

Проверка:

- Проект должен компилироваться, а структура содержать поля из `gorm.Model`.

## Релиз 3 — Подключение к PostgreSQL и AutoMigrate для `Student`

Настрой подключение к PostgreSQL и включи автоматическое создание таблицы для модели `Student` при запуске приложения.

- Используй традиционный API: `gorm.Open(...)` с драйвером PostgreSQL.
- Укажи строку подключения (DSN): хост, порт, пользователь, пароль, имя БД, `sslmode=disable`.
- Вызови `AutoMigrate(&Student{})` после успешного подключения.

Документация:

- Подключение (PostgreSQL): https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
- AutoMigrate: https://gorm.io/docs/migration.html#Auto-Migration

Проверка в pgAdmin:

- Запусти приложение, затем в pgAdmin убедись, что создана таблица, соответствующая `Student` (например, `students`), и у неё есть ожидаемые колонки (включая из `gorm.Model`).

## Релиз 4 — Модель `Group` и добавление в AutoMigrate

Создай вторую модель `Group`. Минимум — `ID` и, например, `Name`. Можно также встроить `gorm.Model`, чтобы автоматом получить стандартные поля.

- Добавь `Group` в автоматическую миграцию (дополни вызов `AutoMigrate`).
- На этом этапе не настраиваем связи — сфокусируйся на второй таблице и миграции.

Документация:

- Модели: https://gorm.io/docs/models.html
- AutoMigrate: https://gorm.io/docs/migration.html#Auto-Migration

Проверка в pgAdmin:

- Перезапусти приложение и проверь, что появилась таблица `groups` с ожидаемыми полями.

## Релиз 5 — Вспомогательные функции (CRUD-минимум)

Реализуй базовые хелперы для работы с БД.

- `func CreateGroup(db *gorm.DB, name string) error`
- `func CreateStudent(db *gorm.DB, name string, age int) error`
- `func DeleteStudent(db *gorm.DB, id uint) error`
- `func GetStudentByID(db *gorm.DB, id uint) (*Student, error)`
- `func UpdateStudentName(db *gorm.DB, id uint, newName string) error`

Подсказки:

- Создание записей: https://gorm.io/docs/create.html
- Поиск/чтение: https://gorm.io/docs/query.html
- Обновление: https://gorm.io/docs/update.html
- Удаление: https://gorm.io/docs/delete.html

Проверка в pgAdmin:

- После вызова `CreateGroup` и `CreateStudent` проверь, что новые записи появились в соответствующих таблицах.
- После вызова `DeleteStudent` проверь, что запись удалена (или помечена как удалённая, если используешь soft delete).
