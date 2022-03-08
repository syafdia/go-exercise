CREATE TABLE resources (
	id SERIAL PRIMARY KEY,
	name VARCHAR (64) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE resource_actions (
	id SERIAL PRIMARY KEY,
	name VARCHAR (64) NOT NULL,
    resource_id BIGINT NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE INDEX idx_resource_id ON resource_actions(resource_id);
CREATE INDEX idx_name ON resource_actions(name);

ALTER TABLE resource_actions ADD CONSTRAINT fk_resource_id FOREIGN KEY (resource_id) REFERENCES resources(id);

CREATE TYPE TARGET_TYPE AS ENUM ('legal_entity', 'industry');

CREATE TABLE target_resource_actions (
	id SERIAL PRIMARY KEY,
    resource_action_id BIGINT NOT NULL,
    target_id BIGINT NOT NULL,
    target_type TARGET_TYPE NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE INDEX idx_resource_action_id ON target_resource_actions(resource_action_id);
CREATE INDEX idx_target_id ON target_resource_actions(target_id);
CREATE INDEX idx_target_type ON target_resource_actions(target_type);

ALTER TABLE target_resource_actions ADD CONSTRAINT fk_resource_action_id FOREIGN KEY (resource_action_id) REFERENCES resource_actions(id);