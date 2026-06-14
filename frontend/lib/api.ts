const API_BASE = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export interface Device {
  id: string;
  tenant_id: string;
  user_id: string;
  name: string;
  fingerprint: string;
  os: string;
  os_version: string;
  device_type: string;
  manufacturer: string;
  model: string;
  status: string;
  trust_score: number;
  last_seen_at: string;
  created_at: string;
}

export interface User {
  id: string;
  tenant_id: string;
  email: string;
  username: string;
  display_name: string;
  mfa_enabled: boolean;
  status: string;
  last_login_at: string;
  created_at: string;
}

export interface Resource {
  id: string;
  tenant_id: string;
  name: string;
  description: string;
  resource_type: string;
  endpoint: string;
  port: number;
  protocol: string;
  sensitivity: string;
  status: string;
  created_at: string;
}

export interface Policy {
  id: string;
  tenant_id: string;
  name: string;
  description: string;
  policy_type: string;
  priority: number;
  enabled: boolean;
  effect: string;
  created_at: string;
}

export interface Session {
  id: string;
  tenant_id: string;
  user_id: string;
  device_id: string;
  resource_id: string;
  token: string;
  status: string;
  risk_score: number;
  granted_at: string;
  expires_at: string;
  created_at: string;
}

export interface AccessRequest {
  id: string;
  tenant_id: string;
  user_id: string;
  device_id: string;
  resource_id: string;
  request_type: string;
  justification: string;
  status: string;
  duration_minutes: number;
  requested_at: string;
  created_at: string;
}

export interface AuditEvent {
  id: string;
  tenant_id: string;
  event_type: string;
  actor_id: string;
  actor_type: string;
  resource_type: string;
  resource_id: string;
  action: string;
  status: string;
  created_at: string;
}

class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string = API_BASE) {
    this.baseUrl = baseUrl;
  }

  private async request<T>(path: string, options?: RequestInit): Promise<T> {
    const response = await fetch(`${this.baseUrl}${path}`, {
      headers: {
        'Content-Type': 'application/json',
        ...options?.headers,
      },
      ...options,
    });

    if (!response.ok) {
      throw new Error(`API error: ${response.status} ${response.statusText}`);
    }

    return response.json();
  }

  async getDevices(tenantId?: string, userId?: string): Promise<Device[]> {
    const params = new URLSearchParams();
    if (tenantId) params.append('tenant_id', tenantId);
    if (userId) params.append('user_id', userId);
    return this.request<Device[]>(`/api/v1/devices?${params.toString()}`);
  }

  async getDevice(id: string): Promise<Device> {
    return this.request<Device>(`/api/v1/devices?id=${id}`);
  }

  async getResources(tenantId: string): Promise<Resource[]> {
    return this.request<Resource[]>(`/api/v1/resources?tenant_id=${tenantId}`);
  }

  async getResource(id: string): Promise<Resource> {
    return this.request<Resource>(`/api/v1/resources?id=${id}`);
  }

  async getSessions(tenantId?: string, userId?: string): Promise<Session[]> {
    const params = new URLSearchParams();
    if (tenantId) params.append('tenant_id', tenantId);
    if (userId) params.append('user_id', userId);
    return this.request<Session[]>(`/api/v1/sessions?${params.toString()}`);
  }

  async getAccessRequests(tenantId: string): Promise<AccessRequest[]> {
    return this.request<AccessRequest[]>(`/api/v1/access-requests?tenant_id=${tenantId}`);
  }

  async getAuditEvents(tenantId: string, limit?: number, offset?: number): Promise<AuditEvent[]> {
    const params = new URLSearchParams({ tenant_id: tenantId });
    if (limit) params.append('limit', limit.toString());
    if (offset) params.append('offset', offset.toString());
    return this.request<AuditEvent[]>(`/api/v1/audit?${params.toString()}`);
  }

  async evaluatePolicy(data: {
    tenant_id: string;
    user_id: string;
    device_id: string;
    resource_id: string;
    user_trust_score: number;
    device_trust_score: number;
    resource_sensitivity: string;
    mfa_verified: boolean;
  }): Promise<{ allowed: boolean; reasons: string[] }> {
    return this.request('/api/v1/policies/evaluate', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }
}

export const api = new ApiClient();
