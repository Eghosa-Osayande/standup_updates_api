-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SET TIMEZONE TO 'UTC';



Create table if not exists employees (
	id uuid primary key DEFAULT uuid_generate_v1mc() NOT NULL,
	created_at TIMESTAMPTZ DEFAULT timezone('utc', now()) NOT NULL,
	name text NOT NULL,
	password text NOT NULL,
	role text NOT NULL
);

Create table if not exists sprints (
	id uuid primary key DEFAULT uuid_generate_v1mc() NOT NULL,
	created_at TIMESTAMPTZ DEFAULT timezone('utc', now()) NOT NULL,
	name text NOT NULL,
	standup_start_time TIMESTAMPTZ NOT NULL,
	standup_end_time TIMESTAMPTZ NOT NULL
);



Create table if not exists updates (

	id uuid primary key DEFAULT uuid_generate_v1mc() NOT NULL,
	created_at TIMESTAMPTZ DEFAULT timezone('utc', now()) NOT NULL,
	employee_id uuid NOT NULL,
	FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
	sprint_id uuid NOT NULL,
	FOREIGN KEY (sprint_id) REFERENCES sprints(id) ON DELETE CASCADE,
	yesterday text NOT NULL,
	today text NOT NULL,
	blocked_by text[] NULL,
	breakaway text NOT NULL,
	check_in_time TIMESTAMPTZ NOT NULL,
	status text NOT NULL,
	tasks text[] NULL

	
);

INSERT INTO employees (name, password, role) VALUES ('admin', 'dummy_placeholder', 'admin') RETURNING *;



-- +goose Down
DROP TABLE IF EXISTS updates;
DROP TABLE IF EXISTS sprints;
DROP TABLE IF EXISTS employees;








