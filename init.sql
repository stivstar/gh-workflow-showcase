-- public.users definition

DROP TABLE IF EXISTS public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	fname varchar NULL,
	lname varchar NULL,
	email varchar NULL
);

INSERT INTO public.users (fname, lname, email) VALUES
 ('Joe', 'Bob', 'joe@labstack'),
 ('Mary', 'Munn', 'mary@labstack'),
 ('Joe', 'Rogan', 'joel@labstack'),
 ('Joe', 'Madden', 'joe@invalid-domain'),
 ('Okay', 'You', 'got@us.lol');

SELECT * FROM users;
