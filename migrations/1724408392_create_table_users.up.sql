-- The user table 
CREATE TABLE user(
    id          INTEGER PRIMARY KEY, 
    username    TEXT, 
    password    TEXT, 
    name        TEXT,
    admin       BOOLEAN
);
