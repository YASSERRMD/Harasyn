-- Migration: 012_create_risk_signals_table
CREATE TABLE IF NOT EXISTS risk_signals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    device_id UUID REFERENCES devices(id) ON DELETE SET NULL,
    signal_type VARCHAR(100) NOT NULL,
    severity VARCHAR(20) DEFAULT 'low',
    score INTEGER DEFAULT 0,
    source VARCHAR(100),
    details JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_risk_signals_tenant_id ON risk_signals(tenant_id);
CREATE INDEX idx_risk_signals_user_id ON risk_signals(user_id);
CREATE INDEX idx_risk_signals_device_id ON risk_signals(device_id);
CREATE INDEX idx_risk_signals_type ON risk_signals(signal_type);
