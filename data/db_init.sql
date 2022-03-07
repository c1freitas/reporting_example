CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS flight_summary
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    external_id text,
    external_mission_id text,
    hardware_id text,
    nickname text,
    meta  jsonb not null default '{}'::jsonb,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
	CONSTRAINT flight_summary_prikey PRIMARY KEY (id)
);
insert into flight_summary (id, external_id, external_mission_id, hardware_id, nickname, meta, created_at, updated_at) values (uuid_generate_v4(), '123', 'abc', uuid_generate_v4(), 'flight1', '{ "pilot": {"name": "John Smith", "state_licensed": "IL"}, "weather": {"temperature": 50, "unit": "F"}  }', now(), now());
insert into flight_summary (id, external_id, external_mission_id, hardware_id, nickname, meta, created_at, updated_at) values (uuid_generate_v4(), '234', 'bcd', uuid_generate_v4(), 'flight2', '{ "pilot": {"name": "Olive Yew", "state_licensed": "NC"}, "weather": {"temperature": 10, "unit": "C"}  }', now(), now());
insert into flight_summary (id, external_id, external_mission_id, hardware_id, nickname, meta, created_at, updated_at) values (uuid_generate_v4(), '345', 'cde', uuid_generate_v4(), 'flight3', '{ "pilot": {"name": "Maureen Biologist", "state_licensed": "CA"}, "weather": {"temperature": 30, "unit": "C"}  }', now(), now());
insert into flight_summary (id, external_id, external_mission_id, hardware_id, nickname, meta, created_at, updated_at) values (uuid_generate_v4(), '456', 'def', uuid_generate_v4(), 'flight4', '{ "pilot": {"name": "Aida Bugg", "state_licensed": "WA"}, "weather": {"temperature": 25, "unit": "C"}  }', now(), now());
insert into flight_summary (id, external_id, external_mission_id, hardware_id, nickname, meta, created_at, updated_at) values (uuid_generate_v4(), '567', 'efg', uuid_generate_v4(), 'flight5', '{ "pilot": {"name": "Aida Cortez", "state_licensed": "NY"}, "weather": {"temperature": 22, "unit": "C"}  }', now(), now());