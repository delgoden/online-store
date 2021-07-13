-- 
CREATE TABLE users (
	id BIGSERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	login TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	role TEXT NOT NULL DEFAULT 'CUSTOMER',
	created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- 
CREATE TABLE users_tokens (
	token TEXT NOT NULL UNIQUE,
	user_id INTEGER NOT NULL,
	created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	expire TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '1 hour'
);
-- 
CREATE TABLE categories (
	id BIGSERIAL PRIMARY KEY,
	name TEXT NOT NULL UNIQUE
);
-- 
CREATE TABLE photos (
	id BIGSERIAL PRIMARY KEY,
	name TEXT NOT NULL UNIQUE,
	product_id INTEGER NOT NULL
);
-- 
CREATE TABLE products (
	id BIGSERIAL PRIMARY KEY,
	name TEXT NOT NULL UNIQUE,
	category_id INTEGER NOT NULL,
	description TEXT NOT NULL,
	qty INTEGER NOT NULL,
	price INTEGER NOT NULL,
	active BOOLEAN NOT NULL DEFAULT true,
	created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
--
CREATE TABLE positions (
	id BIGSERIAL PRIMARY KEY,
	product_id INTEGER NOT NULL,
	qty INTEGER NOT NULL,
	price INTEGER NOT NULL
);
--
CREATE TABLE purchases (
	id BIGSERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	position_id INTEGER NOT NULL,
	created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);