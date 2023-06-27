CREATE TABLE public.employees
(
    id SERIAL,
    username character varying(20) COLLATE pg_catalog."default" NOT NULL,
    password_hash character varying(72) COLLATE pg_catalog."default" NOT NULL,
    name character varying(72) COLLATE pg_catalog."default" NOT NULL,
    telegram_user_id character varying(60) COLLATE pg_catalog."default",

    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,

    CONSTRAINT employees_pkey PRIMARY KEY (id),
    CONSTRAINT employees_password_hash_key UNIQUE (password_hash),
    CONSTRAINT employees_username_key UNIQUE (username)
)
