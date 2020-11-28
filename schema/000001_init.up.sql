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
    link_type           int references specification (id) on delete cascade     not null
);

CREATE TABLE specification_specification
(
    id                  serial                                                  not null unique,
    parent_id           int references specification (id) on delete cascade     not null,
    child_id            int references specification (id) on delete cascade     not null,
    link_type           int references specification (id) on delete cascade     not null
);

INSERT INTO link_type (title) VALUES ('Достоверная');
INSERT INTO link_type (title) VALUES ('Недостоверная');
INSERT INTO link_type (title) VALUES ('Предположительная');