CREATE TABLE users (
    id              int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    username        varchar(50) NOT NULL,
    password_hash   varchar(512) NOT NULL,
    state           int2 CONSTRAINT state_check CHECK (state > 0 and state <=5),
    created_at      timestamp NOT NULL,
    email           varchar(255) NULL
);

create unique index table_name_username_uindex
    on users (username);

create table workout
(
    id          int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    user_id     int NOT NULL CONSTRAINT workout_user_fk REFERENCES users,
    created_at  timestamp NOT NULL,
    description varchar(4096)
);

create table workout_type
(
    id        int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    parent_id int CONSTRAINT workout_type_workout_type_id_fk REFERENCES workout_type ON UPDATE CASCADE ON DELETE CASCADE,
    name      varchar(96) NOT NULL,
    type      smallint    NOT NULL
);

create unique index workout_type_name_uindex
    on workout_type (name);

create unique index workout_type_type_uindex
    on workout_type (type);

create table workout_value
(
    id              int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    workout_id      int NOT NULL CONSTRAINT workout_value_workout_id_fk REFERENCES workout ON UPDATE CASCADE ON DELETE CASCADE,
    workout_type_id int NOT NULL CONSTRAINT workout_value_workout_type_id_fk REFERENCES workout_type ON UPDATE CASCADE ON DELETE CASCADE,
    value           double precision null,
    unit            smallint         not null,
    started_at      timestamp,
    ended_at        timestamp
);
