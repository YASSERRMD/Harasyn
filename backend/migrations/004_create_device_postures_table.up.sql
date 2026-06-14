-- Migration: 004_create_device_postures_table
CREATE TABLE IF NOT EXISTS device_postures (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    device_id UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    encrypted BOOLEAN DEFAULT FALSE,
    jailbroken BOOLEAN DEFAULT FALSE,
    rooted BOOLEAN DEFAULT FALSE,
    patched BOOLEAN DEFAULT FALSE,
    antivirus_enabled BOOLEAN DEFAULT FALSE,
    firewall_enabled BOOLEAN DEFAULT FALSE,
    disk_encrypted BOOLEAN DEFAULT FALSE,
    os_patch_level VARCHAR(100),
    compliance_status VARCHAR(50) DEFAULT 'unknown',
    evaluated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_device_postures_device_id ON device_postures(device_id);
