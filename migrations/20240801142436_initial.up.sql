-- АГЕНТЫ -- АГЕНТЫ -- АГЕНТЫ -- АГЕНТЫ --
create table if not exists agents
(
    id    serial
        constraint agents_pk
            primary key,
    login varchar(25) not null,
    fio   text        not null
);
INSERT INTO agents (login, fio) VALUES ('Kirill', 'Kirill Lagutin Andreevich');
INSERT INTO agents (login, fio) VALUES ('Andrey', 'Andrey Lagutin Kirillov');
INSERT INTO agents (login, fio) VALUES ('Ivan', 'Ivan Ivanov Ivanovich');

-- СТАТУСЫ ЗВОНКОВ -- СТАТУСЫ ЗВОНКОВ --
create table if not exists statuses
(
    id     serial
        constraint "call_states _pk"
            primary key,
    title text not null
        constraint "call_states _pk_2"
            unique
);
INSERT INTO statuses (title) VALUES ('Успешный');
INSERT INTO statuses (title) VALUES ('Не успешный');
INSERT INTO statuses (title) VALUES ('Unknown');

-- ЗВОНКИ -- ЗВОНКИ -- ЗВОНКИ -- ЗВОНКИ --
create table if not exists calls
(
    id          serial
        constraint calls_pk
            primary key,
    caller_id   text      not null,
    agent_id    integer   not null
        constraint calls_agents_id_fk
            references agents,
    call_start  timestamp not null,
    call_end    timestamp not null,
    status_id   integer   not null
        constraint "calls_call_states _id_fk"
            references statuses,
    call_notes  text
);