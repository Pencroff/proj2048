CREATE TABLE game_stat
(
    -- id should be persistent and depend from
    -- game_id, agent_id, mode
    id              bigint NOT NULL,
    game_id         int    NOT NULL,
    agent_id        text   NOT NULL,
    mode            text   NOT NULL,
    field_state     jsonb,
    score           integer,
    step_counter    integer,
    no_move_counter integer,
    max_tile        integer,
    PRIMARY KEY (id)
);

CREATE TABLE game_step
(
    game_stat_id bigint NOT NULL,
    step         bigint NOT NULL,
    score        bigint NOT NULL,
    noMove       boolean,
    field        jsonb  NOT NULL,
    direction    text   NOT NULL,
    PRIMARY KEY (game_stat_id, step),
    CONSTRAINT fk_game_stat_id
        FOREIGN KEY (game_stat_id)
            REFERENCES game_stat (id)
);