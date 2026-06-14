'use client';

import { useEffect, useState } from 'react';
import Table from '@/components/Table';
import StatusBadge from '@/components/StatusBadge';
import Loading from '@/components/Loading';
import ErrorDisplay from '@/components/ErrorDisplay';
import { AccessRequest, api } from '@/lib/api';

const columns = [
  { key: 'user_id', header: 'User' },
  { key: 'resource_id', header: 'Resource' },
  { key: 'request_type', header: 'Type', render: (item: AccessRequest) => (
    <StatusBadge status={item.request_type} variant={item.request_type === 'emergency' ? 'error' : 'info'} />
  )},
  { key: 'status', header: 'Status', render: (item: AccessRequest) => <StatusBadge status={item.status} /> },
  { key: 'duration_minutes', header: 'Duration', render: (item: AccessRequest) => `${item.duration_minutes} min` },
  { key: 'requested_at', header: 'Requested', render: (item: AccessRequest) => new Date(item.requested_at).toLocaleString() },
];

export default function AccessRequestsPage() {
  const [requests, setRequests] = useState<AccessRequest[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchRequests = async () => {
    try {
      setLoading(true);
      const data = await api.getAccessRequests('default');
      setRequests(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch access requests');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchRequests();
  }, []);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Access Requests</h1>
        <p className="mt-1 text-sm text-gray-500">Review and manage access requests</p>
      </div>

      {loading && <Loading />}
      {error && <ErrorDisplay message={error} onRetry={fetchRequests} />}
      {!loading && !error && <Table columns={columns} data={requests} />}
    </div>
  );
}
