CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.product (
    ID    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Name  VARCHAR(255),
    Price NUMERIC
);