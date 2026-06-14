-- Migration: 009_create_access_sessions_table
CREATE TABLE IF NOT EXISTS access_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    device_id UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    resource_id UUID NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
    policy_id UUID REFERENCES access_policies(id) ON DELETE SET NULL,
    token VARCHAR(512) NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    risk_score INTEGER DEFAULT 0,
    granted_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    last_revalidated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    revoked_at TIMESTAMP WITH TIME ZONE,
    revoke_reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_access_sessions_tenant_id ON access_sessions(tenant_id);
CREATE INDEX idx_access_sessions_user_id ON access_sessions(user_id);
CREATE INDEX idx_access_sessions_device_id ON access_sessions(device_id);
CREATE INDEX idx_access_sessions_token ON access_sessions(token);
CREATE INDEX idx_access_sessions_status ON access_sessions(status);
