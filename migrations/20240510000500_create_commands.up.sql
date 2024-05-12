CREATE TABLE commands (
    id bigserial not null primary key,
    command_name varchar not null,
    result varchar not null,
    date_time varchar not null
);