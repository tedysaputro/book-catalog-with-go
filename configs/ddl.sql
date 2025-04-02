-- public.publishers definition

-- Drop table

-- DROP TABLE public.publishers;

CREATE TABLE public.publishers (
	id bigserial NOT NULL,
	name varchar(100) NOT NULL,
	description text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT publishers_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_publishers_deleted_at ON public.publishers USING btree (deleted_at);


-- public.authors definition

-- Drop table

-- DROP TABLE public.authors;

CREATE TABLE public.authors (
	id bigserial NOT NULL,
	"name" varchar(100) NOT NULL,
	description text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT authors_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_authors_deleted_at ON public.authors USING btree (deleted_at);