CREATE TABLE public.todos (
	id uuid NOT NULL,
	title varchar(255) NOT NULL,
	completed boolean NOT NULL,
	CONSTRAINT todos_pk PRIMARY KEY (id)
);


