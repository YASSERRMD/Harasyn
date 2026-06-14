-- Migration: 007_create_access_policies_table
CREATE TABLE IF NOT EXISTS access_policies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    policy_type VARCHAR(50) DEFAULT 'access',
    priority INTEGER DEFAULT 100,
    enabled BOOLEAN DEFAULT TRUE,
    effect VARCHAR(20) DEFAULT 'allow',
    resource_id UUID REFERENCES resources(id) ON DELETE SET NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_access_policies_tenant_id ON access_policies(tenant_id);
CREATE INDEX idx_access_policies_resource_id ON access_policies(resource_id);
