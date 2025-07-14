GRANT USAGE, CREATE ON SCHEMA public TO go_test;

CREATE TABLE public.product (
    ID    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Name  VARCHAR(255),
    Price NUMERIC
);