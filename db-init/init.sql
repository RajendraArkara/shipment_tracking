-- Extension
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- 1. orders
CREATE TABLE orders (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    order_ref varchar(50) NOT NULL,
    buyer_name varchar(150) NOT NULL,
    product_summary text,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT orders_order_ref_key UNIQUE (order_ref)
);

-- 2. carriers
CREATE TABLE carriers (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(150) NOT NULL,
    contact_phone varchar(30),
    vehicle_type varchar(50),
    plate_number varchar(20),
    created_at timestamptz NOT NULL DEFAULT now()
);

-- 3. shipments (FK ke orders dan carriers ditambahkan di sini)
CREATE TABLE shipments (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id uuid REFERENCES orders(id),
    carrier_id uuid REFERENCES carriers(id),
    origin varchar(150) NOT NULL,
    destination varchar(150) NOT NULL,
    current_status varchar(30) NOT NULL DEFAULT 'created',
    eta timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

-- 4. shipment_events
CREATE TABLE shipment_events (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    shipment_id uuid NOT NULL REFERENCES shipments(id) ON DELETE CASCADE,
    status varchar(30) NOT NULL,
    location varchar(150),
    notes text,
    occurred_at timestamptz NOT NULL DEFAULT now(),
    created_at timestamptz NOT NULL DEFAULT now()
);

-- 5. webhook_subscriptions
CREATE TABLE webhook_subscriptions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    shipment_id uuid NOT NULL REFERENCES shipments(id) ON DELETE CASCADE,
    target_url text NOT NULL,
    secret_key varchar(100) NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now()
);

-- 6. api_keys
CREATE TABLE api_keys (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    key_hash varchar(255) NOT NULL,
    owner_name varchar(150) NOT NULL,
    rate_limit_per_min integer NOT NULL DEFAULT 60,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT api_keys_key_hash_key UNIQUE (key_hash)
);