
-- +goose Up
CREATE TABLE todos
  (id serial PRIMARY KEY,
   text VARCHAR(255),
   completed BOOLEAN NOT NULL DEFAULT FALSE);

-- SQL in section 'Up' is executed when this migration is applied


-- +goose Down
DROP TABLE todos;
-- SQL section 'Down' is executed when this migration is rolled back

