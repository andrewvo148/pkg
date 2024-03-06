-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS "jobs";

CREATE TABLE jobs.companies (
    id uuid NOT NULL DEFAULT (gen_random_uuid()),
    title text NOT NULL,
    description text NOT NULL,
    create_time timestamp with time zone NOT NULL DEFAULT (now()),
    update_time timestamp with time zone,
    CONSTRAINT pk_companies PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE companies;
-- +goose StatementEnd