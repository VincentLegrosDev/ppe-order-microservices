-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE SCHEMA orders AUTHORIZATION postgres;

CREATE TABLE  orders.processedEvents (
  id uuid,
  event_name varchar(255),
  processed_time timestamp NOT NULL, 
  PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE processedEvent; 
-- +goose StatementEnd

