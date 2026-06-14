'use client';

import React, { useState, useEffect } from 'react';

interface AuditLogEntry {
  id: string;
  timestamp: string;
  user: string;
  action: string;
  resource: string;
  details: string;
  riskLevel: string;
}

interface RealTimeDashboardProps {
  tenantId: string;
}

export default function RealTimeDashboard({ tenantId }: RealTimeDashboardProps) {
  const [logs, setLogs] = useState<AuditLogEntry[]>([]);
  const [stats, setStats] = useState({
    activeSessions: 0,
    blockedRequests: 0,
    deviceCount: 0,
    pendingApprovals: 0,
  });

  useEffect(() => {
    setStats({
      activeSessions: 47,
      blockedRequests: 12,
      deviceCount: 234,
      pendingApprovals: 3,
    });

    const mockLogs: AuditLogEntry[] = [
      { id: '1', timestamp: new Date().toISOString(), user: 'alice@acme.com', action: 'access_granted', resource: 'api.production', details: 'Policy: dev-access', riskLevel: 'low' },
      { id: '2', timestamp: new Date(Date.now() - 30000).toISOString(), user: 'bob@acme.com', action: 'device_posture_failed', resource: 'laptop-bob', details: 'OS version outdated', riskLevel: 'high' },
      { id: '3', timestamp: new Date(Date.now() - 60000).toISOString(), user: 'charlie@acme.com', action: 'access_denied', resource: 'admin-console', details: 'Policy: admin-access', riskLevel: 'medium' },
    ];
    setLogs(mockLogs);
  }, [tenantId]);

  return (
    <div className="space-y-6">
      <h2 className="text-lg font-semibold text-gray-900">Real-Time Activity</h2>
      
      <div className="grid grid-cols-4 gap-4">
        <div className="bg-white rounded-lg shadow p-4">
          <div className="text-sm text-gray-500">Active Sessions</div>
          <div className="text-2xl font-bold text-gray-900">{stats.activeSessions}</div>
        </div>
        <div className="bg-white rounded-lg shadow p-4">
          <div className="text-sm text-gray-500">Blocked Requests</div>
          <div className="text-2xl font-bold text-red-600">{stats.blockedRequests}</div>
        </div>
        <div className="bg-white rounded-lg shadow p-4">
          <div className="text-sm text-gray-500">Registered Devices</div>
          <div className="text-2xl font-bold text-gray-900">{stats.deviceCount}</div>
        </div>
        <div className="bg-white rounded-lg shadow p-4">
          <div className="text-sm text-gray-500">Pending Approvals</div>
          <div className="text-2xl font-bold text-yellow-600">{stats.pendingApprovals}</div>
        </div>
      </div>

      <div className="bg-white rounded-lg shadow overflow-hidden">
        <div className="px-4 py-3 border-b border-gray-200">
          <h3 className="font-medium text-gray-900">Recent Activity</h3>
        </div>
        <div className="divide-y divide-gray-200">
          {logs.map((log) => (
            <div key={log.id} className="px-4 py-3 hover:bg-gray-50">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <div className={`w-2 h-2 rounded-full ${
                    log.riskLevel === 'high' ? 'bg-red-500' :
                    log.riskLevel === 'medium' ? 'bg-yellow-500' : 'bg-green-500'
                  }`} />
                  <div>
                    <div className="text-sm font-medium text-gray-900">{log.user}</div>
                    <div className="text-sm text-gray-500">{log.action} - {log.resource}</div>
                  </div>
                </div>
                <div className="text-xs text-gray-400">
                  {new Date(log.timestamp).toLocaleTimeString()}
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
