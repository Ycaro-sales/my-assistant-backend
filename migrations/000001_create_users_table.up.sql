CREATE TABLE IF NOT EXISTS users(
	user_id serial PRIMARY KEY,
	username text NOT NULL,
	password text NOT NULL,
	email text NOT NULL,
);
