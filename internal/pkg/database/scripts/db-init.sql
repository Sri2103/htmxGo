CREATE TABLE "public".todos (
    title varchar(128) NOT NULL,
    description varchar(128),
    id numeric NOT NULL,
    CONSTRAINT pk_todos PRIMARY KEY (id)
)