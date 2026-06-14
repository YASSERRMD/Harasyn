-- Migration: 008_create_policy_conditions_table
CREATE TABLE IF NOT EXISTS policy_conditions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    policy_id UUID NOT NULL REFERENCES access_policies(id) ON DELETE CASCADE,
    condition_type VARCHAR(100) NOT NULL,
    operator VARCHAR(20) NOT NULL,
    value TEXT NOT NULL,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_policy_conditions_policy_id ON policy_conditions(policy_id);
