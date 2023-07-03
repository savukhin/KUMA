CREATE TABLE public.cnc_checkers
(
    id SERIAL,
    username character varying(20) COLLATE pg_catalog."default" NOT NULL,
    password_hash character varying(72) COLLATE pg_catalog."default" NOT NULL,
    status cnc_status,
    
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,

    CONSTRAINT cnc_checkers_pkey PRIMARY KEY (id),
    CONSTRAINT cnc_checkers_password_hash_key UNIQUE (password_hash),
    CONSTRAINT cnc_checkers_username_key UNIQUE (username)
)

TABLESPACE pg_default;
