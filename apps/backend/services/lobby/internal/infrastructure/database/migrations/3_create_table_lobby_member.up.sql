CREATE TABLE lobby_member (
    lobby_id INTEGER NOT NULL,
    member_id INTEGER NOT NULL,
    is_ready BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (lobby_id) REFERENCES lobby (id) ON DELETE CASCADE,
    FOREIGN KEY (member_id) REFERENCES member (id) ON DELETE CASCADE,
    PRIMARY KEY (lobby_id, member_id)
);