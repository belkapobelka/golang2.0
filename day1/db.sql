-- Table: public."Equation"

-- DROP TABLE public."Equation";

CREATE TABLE public."Equation"
(
    "A" integer NOT NULL,
    "B" integer,
    "C" integer,
    "NRoots" integer,
    "Id" integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    CONSTRAINT "Equation_pkey" PRIMARY KEY ("Id")
)

    TABLESPACE pg_default;

ALTER TABLE public."Equation"
    OWNER to postgres;