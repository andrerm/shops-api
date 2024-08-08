-- DROP SCHEMA public;

CREATE SCHEMA public AUTHORIZATION pg_database_owner;
-- public.paymenttypes definition

-- Drop table

-- DROP TABLE public.paymenttypes;

CREATE TABLE public.paymenttypes (
	payment_type_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	type_name varchar(50) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT paymenttypes_pkey PRIMARY KEY (payment_type_id)
);


-- public.roles definition

-- Drop table

-- DROP TABLE public.roles;

CREATE TABLE public.roles (
	role_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	role_name varchar(50) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT roles_pkey PRIMARY KEY (role_id)
);


-- public.stores definition

-- Drop table

-- DROP TABLE public.stores;

CREATE TABLE public.stores (
	store_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	store_name varchar(100) NOT NULL,
	"location" varchar(100) NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT stores_pkey PRIMARY KEY (store_id)
);


-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	user_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	"name" varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	"password" varchar(100) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);


-- public.products definition

-- Drop table

-- DROP TABLE public.products;

CREATE TABLE public.products (
	product_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	store_id uuid NULL,
	"name" varchar(100) NOT NULL,
	category varchar(50) NOT NULL,
	price numeric(10, 2) NOT NULL,
	stock_quantity int4 NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT products_pkey PRIMARY KEY (product_id),
	CONSTRAINT products_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(store_id)
);


-- public.transactions definition

-- Drop table

-- DROP TABLE public.transactions;

CREATE TABLE public.transactions (
	transaction_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NULL,
	product_id uuid NULL,
	store_id uuid NULL,
	payment_type_id uuid NULL,
	amount numeric(10, 2) NOT NULL,
	"date" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id),
	CONSTRAINT transactions_payment_type_id_fkey FOREIGN KEY (payment_type_id) REFERENCES public.paymenttypes(payment_type_id),
	CONSTRAINT transactions_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(product_id),
	CONSTRAINT transactions_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(store_id),
	CONSTRAINT transactions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id)
);


-- public.userroles definition

-- Drop table

-- DROP TABLE public.userroles;

CREATE TABLE public.userroles (
	user_role_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NULL,
	role_id uuid NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT userroles_pkey PRIMARY KEY (user_role_id),
	CONSTRAINT userroles_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(role_id),
	CONSTRAINT userroles_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id)
);


-- public.wallets definition

-- Drop table

-- DROP TABLE public.wallets;

CREATE TABLE public.wallets (
	wallet_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NULL,
	balance numeric(10, 2) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT wallets_pkey PRIMARY KEY (wallet_id),
	CONSTRAINT wallets_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id)
);


-- public.bills definition

-- Drop table

-- DROP TABLE public.bills;

CREATE TABLE public.bills (
	bill_id uuid DEFAULT uuid_generate_v4() NOT NULL,
	transaction_id uuid NULL,
	total_amount numeric(10, 2) NOT NULL,
	"date" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created_by uuid NULL,
	updated_by uuid NULL,
	CONSTRAINT bills_pkey PRIMARY KEY (bill_id),
	CONSTRAINT bills_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(transaction_id)
);



-- DROP FUNCTION public.uuid_generate_v1();

CREATE OR REPLACE FUNCTION public.uuid_generate_v1()
 RETURNS uuid
 LANGUAGE c
 PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v1$function$
;

-- DROP FUNCTION public.uuid_generate_v1mc();

CREATE OR REPLACE FUNCTION public.uuid_generate_v1mc()
 RETURNS uuid
 LANGUAGE c
 PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v1mc$function$
;

-- DROP FUNCTION public.uuid_generate_v3(uuid, text);

CREATE OR REPLACE FUNCTION public.uuid_generate_v3(namespace uuid, name text)
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v3$function$
;

-- DROP FUNCTION public.uuid_generate_v4();

CREATE OR REPLACE FUNCTION public.uuid_generate_v4()
 RETURNS uuid
 LANGUAGE c
 PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v4$function$
;

-- DROP FUNCTION public.uuid_generate_v5(uuid, text);

CREATE OR REPLACE FUNCTION public.uuid_generate_v5(namespace uuid, name text)
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v5$function$
;

-- DROP FUNCTION public.uuid_nil();

CREATE OR REPLACE FUNCTION public.uuid_nil()
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_nil$function$
;

-- DROP FUNCTION public.uuid_ns_dns();

CREATE OR REPLACE FUNCTION public.uuid_ns_dns()
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_ns_dns$function$
;

-- DROP FUNCTION public.uuid_ns_oid();

CREATE OR REPLACE FUNCTION public.uuid_ns_oid()
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_ns_oid$function$
;

-- DROP FUNCTION public.uuid_ns_url();

CREATE OR REPLACE FUNCTION public.uuid_ns_url()
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_ns_url$function$
;

-- DROP FUNCTION public.uuid_ns_x500();

CREATE OR REPLACE FUNCTION public.uuid_ns_x500()
 RETURNS uuid
 LANGUAGE c
 IMMUTABLE PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_ns_x500$function$
;