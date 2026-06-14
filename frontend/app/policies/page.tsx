'use client';

import StatusBadge from '@/components/StatusBadge';

const policies = [
  {
    id: '1',
    name: 'Default Access Policy',
    description: 'Standard access policy for internal resources',
    effect: 'allow',
    enabled: true,
    priority: 100,
    conditions: [
      { type: 'device_trust', operator: '>=', value: '60' },
      { type: 'user_trust', operator: '>=', value: '50' },
    ],
  },
  {
    id: '2',
    name: 'Critical Resource Policy',
    description: 'Strict policy for critical resources',
    effect: 'allow',
    enabled: true,
    priority: 50,
    conditions: [
      { type: 'device_trust', operator: '>=', value: '80' },
      { type: 'user_trust', operator: '>=', value: '70' },
      { type: 'mfa_status', operator: '==', value: 'true' },
    ],
  },
  {
    id: '3',
    name: 'Block Untrusted Devices',
    description: 'Block devices with low trust scores',
    effect: 'deny',
    enabled: true,
    priority: 200,
    conditions: [
      { type: 'device_trust', operator: '<', value: '40' },
    ],
  },
];

export default function PoliciesPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Policies</h1>
        <p className="mt-1 text-sm text-gray-500">Manage Zero Trust access policies</p>
      </div>

      <div className="bg-white shadow rounded-lg divide-y divide-gray-200">
        {policies.map((policy) => (
          <div key={policy.id} className="p-6">
            <div className="flex items-center justify-between">
              <div>
                <h3 className="text-lg font-medium text-gray-900">{policy.name}</h3>
                <p className="mt-1 text-sm text-gray-500">{policy.description}</p>
              </div>
              <div className="flex items-center space-x-3">
                <StatusBadge status={policy.effect} />
                <StatusBadge status={policy.enabled ? 'Enabled' : 'Disabled'} />
              </div>
            </div>
            <div className="mt-4">
              <h4 className="text-sm font-medium text-gray-700">Conditions</h4>
              <div className="mt-2 flex flex-wrap gap-2">
                {policy.conditions.map((cond, idx) => (
                  <span key={idx} className="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-gray-100 text-gray-800">
                    {cond.type} {cond.operator} {cond.value}
                  </span>
                ))}
              </div>
            </div>
            <div className="mt-4 text-xs text-gray-400">
              Priority: {policy.priority}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
