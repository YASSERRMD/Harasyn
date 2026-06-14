-- Migration: 013_create_trust_scores_table
CREATE TABLE IF NOT EXISTS trust_scores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    entity_type VARCHAR(50) NOT NULL,
    entity_id UUID NOT NULL,
    score INTEGER NOT NULL CHECK (score >= 0 AND score <= 100),
    components JSONB DEFAULT '{}',
    calculated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_trust_scores_tenant_id ON trust_scores(tenant_id);
CREATE INDEX idx_trust_scores_entity ON trust_scores(entity_type, entity_id);
