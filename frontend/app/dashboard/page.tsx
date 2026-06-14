'use client';

import TrustScoreCard from '@/components/TrustScoreCard';

const mockStats = {
  totalDevices: 24,
  activeSessions: 8,
  pendingRequests: 3,
  policyViolations: 1,
  avgDeviceTrust: 78,
  avgUserTrust: 85,
};

export default function DashboardPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
        <p className="mt-1 text-sm text-gray-500">
          Overview of your Zero Trust Access Platform
        </p>
      </div>

      <div className="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
        <div className="bg-white overflow-hidden shadow rounded-lg p-5">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <svg className="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
            </div>
            <div className="ml-5 w-0 flex-1">
              <dl>
                <dt className="text-sm font-medium text-gray-500 truncate">Total Devices</dt>
                <dd className="text-lg font-semibold text-gray-900">{mockStats.totalDevices}</dd>
              </dl>
            </div>
          </div>
        </div>

        <div className="bg-white overflow-hidden shadow rounded-lg p-5">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <svg className="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
            </div>
            <div className="ml-5 w-0 flex-1">
              <dl>
                <dt className="text-sm font-medium text-gray-500 truncate">Active Sessions</dt>
                <dd className="text-lg font-semibold text-gray-900">{mockStats.activeSessions}</dd>
              </dl>
            </div>
          </div>
        </div>

        <div className="bg-white overflow-hidden shadow rounded-lg p-5">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <svg className="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
              </svg>
            </div>
            <div className="ml-5 w-0 flex-1">
              <dl>
                <dt className="text-sm font-medium text-gray-500 truncate">Pending Requests</dt>
                <dd className="text-lg font-semibold text-gray-900">{mockStats.pendingRequests}</dd>
              </dl>
            </div>
          </div>
        </div>

        <div className="bg-white overflow-hidden shadow rounded-lg p-5">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <svg className="h-6 w-6 text-red-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
            </div>
            <div className="ml-5 w-0 flex-1">
              <dl>
                <dt className="text-sm font-medium text-gray-500 truncate">Policy Violations</dt>
                <dd className="text-lg font-semibold text-red-600">{mockStats.policyViolations}</dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <div className="grid grid-cols-1 gap-5 sm:grid-cols-2">
        <TrustScoreCard title="Average Device Trust" score={mockStats.avgDeviceTrust} description="Across all registered devices" />
        <TrustScoreCard title="Average User Trust" score={mockStats.avgUserTrust} description="Across all active users" />
      </div>

      <div className="bg-white shadow rounded-lg">
        <div className="px-4 py-5 sm:p-6">
          <h3 className="text-lg leading-6 font-medium text-gray-900">Recent Activity</h3>
          <div className="mt-4 space-y-3">
            <div className="flex items-center text-sm text-gray-500">
              <span className="font-medium text-gray-900">Device registered:</span>
              <span className="ml-2">MacBook Pro - John Doe</span>
              <span className="ml-auto text-xs">2 minutes ago</span>
            </div>
            <div className="flex items-center text-sm text-gray-500">
              <span className="font-medium text-gray-900">Session created:</span>
              <span className="ml-2">Access to Database Server</span>
              <span className="ml-auto text-xs">5 minutes ago</span>
            </div>
            <div className="flex items-center text-sm text-gray-500">
              <span className="font-medium text-gray-900">Policy evaluated:</span>
              <span className="ml-2">Default access policy - DENIED</span>
              <span className="ml-auto text-xs">10 minutes ago</span>
            </div>
            <div className="flex items-center text-sm text-gray-500">
              <span className="font-medium text-gray-900">Access request approved:</span>
              <span className="ml-2">Emergency access to Production API</span>
              <span className="ml-auto text-xs">15 minutes ago</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
