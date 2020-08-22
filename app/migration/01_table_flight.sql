DROP TABLE IF EXISTS public.flight;

CREATE TABLE public.flight (
	id bigserial NOT NULL,
    flight_number varchar(10) NOT NULL UNIQUE,
    departure_port varchar(50) NOT NULL,
    arrival_port varchar(50) NOT NULL,
    departure_time timestamp NOT NULL,
    arrival_time timestamp NOT NULL,
	created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp NULL,
	CONSTRAINT flight_pkey PRIMARY KEY (id)
);