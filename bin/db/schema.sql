-- User
CREATE TABLE IF NOT EXISTS users (
    id                  serial PRIMARY KEY,
    email               text NOT NULL UNIQUE,
    name                text NOT NULL,
    contact_name        text default '',    -- For company only
    password            text NOT NULL,
    rg                  text default '',
    cpf                 text default '',
    cnpj                text default '',    -- For company only
    state_registration  text default '',    -- For company only
    mobile_number              text default '',
    permission          integer default 0,
    saved               boolean default true,
    created_at          timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    changed_at          timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_deleted          boolean default false
);
CREATE INDEX IF NOT EXISTS idx_user_email ON users(email);

-- Email confirmation
CREATE TABLE IF NOT EXISTS email_confirmations (
    email               text PRIMARY KEY,
    uuid                text NOT NULL,
    name                text default '',  -- Name must be empty when used for change email, instead to create a new user. 
    password            text NOT NULL,
    created_at          timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Password reset
CREATE TABLE IF NOT EXISTS password_resets (
    user_email          text PRIMARY KEY,
    uuid                text NOT NULL,
    created_at          timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_email) REFERENCES users(email)
);

-- Products
CREATE TABLE IF NOT EXISTS products (
    id                  serial PRIMARY KEY,
    code                text default '',
    title               text NOT NULL,
    category            text default '',
    maker               text default '',
    description         text default '',
    ean                 text default '',
    ncm                 text default '',
    warranty_month      integer default 0,
    length_mm           integer NOT NULL,
    width_mm            integer NOT NULL,
    height_mm           integer NOT NULL,
    weight_g            integer NOT NULL,
    is_active           boolean default false,   -- System defined
    is_available        boolean default false,   -- Dealer defined
    is_sell             boolean default false,   -- User defined
    stock_origin        text default '',
    stock_count         integer default 0,
    price_buy           integer default 0,
    price_sale          integer default 0,
    supplier_code       text default '',
    supplier_name       text default '',
    created_at          timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    changed_at          timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_deleted          boolean default false
);
