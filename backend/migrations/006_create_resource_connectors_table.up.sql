-- Migration: 006_create_resource_connectors_table
CREATE TABLE IF NOT EXISTS resource_connectors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resource_id UUID NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
    connector_type VARCHAR(50) NOT NULL,
    target_host VARCHAR(500) NOT NULL,
    target_port INTEGER,
    tls_required BOOLEAN DEFAULT TRUE,
    auth_method VARCHAR(50) DEFAULT 'token',
    config JSONB DEFAULT '{}',
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_resource_connectors_resource_id ON resource_connectors(resource_id);
