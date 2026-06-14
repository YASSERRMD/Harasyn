-- Migration: 003_create_devices_table
CREATE TABLE IF NOT EXISTS devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    fingerprint VARCHAR(512) NOT NULL,
    os VARCHAR(100) NOT NULL,
    os_version VARCHAR(100),
    device_type VARCHAR(100),
    manufacturer VARCHAR(255),
    model VARCHAR(255),
    status VARCHAR(50) DEFAULT 'active',
    trust_score INTEGER DEFAULT 50,
    last_seen_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(tenant_id, fingerprint)
);

CREATE INDEX idx_devices_tenant_id ON devices(tenant_id);
CREATE INDEX idx_devices_user_id ON devices(user_id);
CREATE INDEX idx_devices_fingerprint ON devices(fingerprint);
