CREATE TABLE users
(
    id                  serial                                                  not null unique,
    name                varchar(255)                                            not null,
    username            varchar(255)                                            not null unique,
    password_hash       varchar(255)                                            not null
);

CREATE TABLE data_owner
(
    id                  serial                                                  not null unique,
    title               varchar(255)                                            not null,
    link_to_logo        varchar(255)
);

CREATE TABLE dataset
(
    id                  serial                                                  not null unique,
    data_owner_id       int references data_owner (id) on delete cascade        not null,
    title               varchar(255)                                            not null,
    description         varchar(1024)
);

CREATE TABLE specification
(
    id                  serial                                                  not null unique,
    dataset_id          int references dataset (id) on delete cascade           not null,
    title               varchar(255)                                            not null,
    description         varchar(255),
    is_primary          boolean                                                 not null default false,
    is_reference        boolean                                                 not null default false
);

CREATE TABLE link_type
(
    id                  serial                                                  not null unique,
    title               varchar(255)                                            not null
);

CREATE TABLE tag
(
    id                  serial                                                  not null unique,
    title               varchar(255)                                            not null
);

CREATE TABLE dataset_tag
(
    id                  serial                                                  not null unique,
    tag_id              int references tag (id) on delete cascade               not null,
    dataset_id          int references dataset (id) on delete cascade           not null
);

CREATE TABLE dataset_specification
(
    id                  serial                                                  not null unique,
    specification_id    int references specification (id) on delete cascade     not null,
    dataset_id          int references dataset (id)       on delete cascade     not null,
    link_type           int references link_type (id) on delete cascade         not null
);

CREATE TABLE specification_specification
(
    id                  serial                                                  not null unique,
    parent_id           int references specification (id) on delete cascade     not null,
    child_id            int references specification (id) on delete cascade     not null,
    link_type           int references link_type (id) on delete cascade         not null
);

INSERT INTO tag (id, title) VALUES (1, 'физические лица');
INSERT INTO tag (id, title) VALUES (2, 'юридические лица');

INSERT INTO link_type (id, title) VALUES (1, 'Достоверная');
INSERT INTO link_type (id, title) VALUES (2, 'Недостоверная');
INSERT INTO link_type (id, title) VALUES (3, 'Предположительная');

-- Test data

INSERT INTO data_owner (id, title, link_to_logo) VALUES (1, 'ПФР', ' ');
INSERT INTO data_owner (id, title, link_to_logo) VALUES (2, 'ФНС', ' ');
INSERT INTO data_owner (id, title, link_to_logo) VALUES (3, 'ФОМС', ' ');
INSERT INTO data_owner (id, title, link_to_logo) VALUES (4, 'МВД', ' ');

INSERT INTO dataset (id, data_owner_id, title, description) VALUES (1, 1, 'Зарегистрированные в ПФР лица', ' ');
INSERT INTO dataset (id, data_owner_id, title, description) VALUES (2, 1, 'База данных СНИЛС', ' ');
INSERT INTO dataset (id, data_owner_id, title, description) VALUES (3, 2, 'ЕГРЮЛ', ' ');
INSERT INTO dataset (id, data_owner_id, title, description) VALUES (4, 2, 'Лица на налоговом учете', ' ');
INSERT INTO dataset (id, data_owner_id, title, description) VALUES (5, 2, 'База данных ИНН', ' ');
INSERT INTO dataset (id, data_owner_id, title, description) VALUES (6, 3, 'База данных полисов ОМС', ' ');
INSERT INTO dataset (id, data_owner_id, title, description) VALUES (7, 3, 'Страховые организации', ' ');
INSERT INTO dataset (id, data_owner_id, title, description) VALUES (8, 4, 'База данных паспортов', ' ');
INSERT INTO dataset (id, data_owner_id, title, description) VALUES (9, 4, 'Регистрация по месту пребывания', ' ');
INSERT INTO dataset (id, data_owner_id, title, description) VALUES (10, 4, 'Органы внутренних дел', ' ');

INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (1, 1, 'Фамилия', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (2, 1, 'Имя', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (3, 1, 'Отчество', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (4, 1, 'Дата рождения', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (5, 1, 'СНИЛС', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (6, 1, 'ИНН', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (7, 1, 'Номер паспорта', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (8, 2, 'Номер', ' ', true, true);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (9, 3, 'Название ЮЛ', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (10, 3, 'ОГРН', ' ', true, true);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (11, 3, 'ИНН', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (12, 4, 'ИНН', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (13, 4, 'Фамилия', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (14, 4, 'Имя', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (15, 4, 'Отчество', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (16, 4, 'Дата рождения', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (17, 5, 'ИНН', ' ', true, true);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (18, 5, 'Дата выдачи', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (19, 6, 'Фамилия', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (20, 6, 'Имя', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (21, 6, 'Отчество', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (22, 6, 'Дата рождения', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (23, 6, 'Паспорт', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (24, 6, 'Страховая организация', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (25, 6, 'Дата постановки на учет', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (26, 7, 'Название', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (27, 7, 'ИНН', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (28, 7, 'ОГРН', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (30, 8, 'Фамилия', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (31, 8, 'Имя', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (32, 8, 'Отчество', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (33, 8, 'Дата рождения', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (34, 8, 'Место рождения', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (35, 8, 'Дата выдачи', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (36, 8, 'Кем выдан', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (37, 8, 'Код подразделения', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (29, 8, 'Номер паспорта', ' ', true, true);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (38, 9, 'Номер паспорта', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (39, 9, 'Дата регистрации', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (40, 9, 'Место регистрации', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (41, 9, 'Орган регистрации', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (42, 10, 'Название', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (43, 10, 'Территория обслуживания', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (44, 10, 'Адрес', ' ', false, false);
INSERT INTO specification (id, dataset_id, title, description, is_primary, is_reference) VALUES (45, 10, 'Телефон', ' ', false, false);

INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (1, 8, 5, 1);
INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (2, 17, 11, 1);
INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (3, 17, 12, 1);
INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (4, 17, 27, 3);
INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (5, 26, 24, 2);
INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (6, 29, 7, 3);
INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (7, 12, 6, 1);
INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (8, 10, 28, 3);
INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (9, 29, 38, 1);
INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (10, 42, 36, 1);
INSERT INTO specification_specification (id, parent_id, child_id, link_type) VALUES (11, 42, 41, 1);

