'use client';

import { useEffect, useState } from 'react';
import Table from '@/components/Table';
import StatusBadge from '@/components/StatusBadge';
import Loading from '@/components/Loading';
import ErrorDisplay from '@/components/ErrorDisplay';
import { Resource, api } from '@/lib/api';

const columns = [
  { key: 'name', header: 'Resource Name' },
  { key: 'resource_type', header: 'Type' },
  { key: 'sensitivity', header: 'Sensitivity', render: (item: Resource) => (
    <StatusBadge status={item.sensitivity} />
  )},
  { key: 'status', header: 'Status', render: (item: Resource) => <StatusBadge status={item.status} /> },
  { key: 'endpoint', header: 'Endpoint' },
];

export default function ResourcesPage() {
  const [resources, setResources] = useState<Resource[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchResources = async () => {
    try {
      setLoading(true);
      const data = await api.getResources('default');
      setResources(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch resources');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchResources();
  }, []);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Resources</h1>
        <p className="mt-1 text-sm text-gray-500">Manage protected resources and their configurations</p>
      </div>

      {loading && <Loading />}
      {error && <ErrorDisplay message={error} onRetry={fetchResources} />}
      {!loading && !error && <Table columns={columns} data={resources} />}
    </div>
  );
}
