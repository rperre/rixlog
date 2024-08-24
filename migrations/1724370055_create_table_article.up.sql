-- The article table 
CREATE TABLE article (
    id      INTEGER PRIMARY KEY, 
    user_id INTEGER, 
    title   TEXT, 
    body    TEXT, 
    slug    TEXT
);
