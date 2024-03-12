CREATE TABLE IF NOT EXISTS cars (
    id uuid PRIMARY KEY DEFAULT ,
    name Varchar(50) NOT NULL,
    brand Varchar(20) NOT NULL,
    model Varchar(30) NOT NULL,
    year INTEGER NOT NUll,
    hourse_power INTEGER DEFAULT 0,
    colour VARCHAR(20) NOT NULL DEFAULT 'black',
    engine_cap DECIMAL(10,2) NOT NULL DEFAULT 1.0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);


CREATE TABLE IF NOT EXISTS customerss (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) UNIQUE,
    gmail VARCHAR(50) NOT NULL UNIQUE,--NEED VALIDATION
    phone VARCHAR(20) NOT NULL,--NEED VALIDATION
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
 
)

ALTER TABLE customerss
ADD CONSTRAINT cu_customerss UNIQUE (delete_at,phone);

cu_customerss

ALTER TABLE customerss
DROP CONSTRAINT cu_customerss;