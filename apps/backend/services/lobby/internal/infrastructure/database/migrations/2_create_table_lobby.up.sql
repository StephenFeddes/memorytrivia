CREATE TABLE lobby (
    id SERIAL PRIMARY KEY,
    size INTEGER NOT NULL,
    owner_id INTEGER NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES member (id) ON DELETE CASCADE ON UPDATE CASCADE
);