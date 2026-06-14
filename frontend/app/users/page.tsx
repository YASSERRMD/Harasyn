'use client';

import { useEffect, useState } from 'react';
import Table from '@/components/Table';
import StatusBadge from '@/components/StatusBadge';
import Loading from '@/components/Loading';
import ErrorDisplay from '@/components/ErrorDisplay';
import TrustScoreCard from '@/components/TrustScoreCard';
import { User, api } from '@/lib/api';

const columns = [
  { key: 'email', header: 'Email' },
  { key: 'display_name', header: 'Name' },
  { key: 'mfa_enabled', header: 'MFA', render: (item: User) => (
    <StatusBadge status={item.mfa_enabled ? 'Enabled' : 'Disabled'} />
  )},
  { key: 'status', header: 'Status', render: (item: User) => <StatusBadge status={item.status} /> },
  { key: 'last_login_at', header: 'Last Login', render: (item: User) => item.last_login_at ? new Date(item.last_login_at).toLocaleString() : 'Never' },
];

export default function UsersPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchUsers = async () => {
    try {
      setLoading(true);
      const data = await api.getDevices();
      setUsers([]);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch users');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Users</h1>
        <p className="mt-1 text-sm text-gray-500">Manage user identities and trust scores</p>
      </div>

      <div className="grid grid-cols-1 gap-5 sm:grid-cols-2">
        <TrustScoreCard title="Average User Trust" score={85} description="Across all active users" />
        <TrustScoreCard title="MFA Adoption" score={72} description="Users with MFA enabled" />
      </div>

      {loading && <Loading />}
      {error && <ErrorDisplay message={error} onRetry={fetchUsers} />}
      {!loading && !error && <Table columns={columns} data={users} />}
    </div>
  );
}
