CREATE TABLE game_stat
(
    -- id should be persistent and depend from game_id, agent_id, mode
    id              bigint PRIMARY KEY NOT NULL,
    game_id         int                NOT NULL,
    agent_id        text               NOT NULL,
    mode            text               NOT NULL,
    field_start_end jsonb,
    score           integer,
    step_counter    integer,
    max_tile        integer,
    no_move_counter integer
);

CREATE TABLE game_step
(
    -- id should be persistent and depend from game_stat_id and step
    id           bigint PRIMARY KEY NOT NULL,
    game_stat_id bigint             NOT NULL,
    step         bigint             NOT NULL,
    score        bigint             NOT NULL,
    noMove       boolean,
    field        jsonb              NOT NULL,
    direction    text               NOT NULL
);