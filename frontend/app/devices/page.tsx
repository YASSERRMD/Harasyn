'use client';

import { useEffect, useState } from 'react';
import Table from '@/components/Table';
import StatusBadge from '@/components/StatusBadge';
import Loading from '@/components/Loading';
import ErrorDisplay from '@/components/ErrorDisplay';
import { Device, api } from '@/lib/api';

const columns = [
  { key: 'name', header: 'Device Name' },
  { key: 'os', header: 'OS', render: (item: Device) => `${item.os} ${item.os_version}` },
  { key: 'trust_score', header: 'Trust Score', render: (item: Device) => (
    <span className={`font-medium ${item.trust_score >= 70 ? 'text-green-600' : item.trust_score >= 40 ? 'text-yellow-600' : 'text-red-600'}`}>
      {item.trust_score}
    </span>
  )},
  { key: 'status', header: 'Status', render: (item: Device) => <StatusBadge status={item.status} /> },
  { key: 'last_seen_at', header: 'Last Seen', render: (item: Device) => item.last_seen_at ? new Date(item.last_seen_at).toLocaleString() : 'Never' },
];

export default function DevicesPage() {
  const [devices, setDevices] = useState<Device[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchDevices = async () => {
    try {
      setLoading(true);
      const data = await api.getDevices();
      setDevices(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch devices');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDevices();
  }, []);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Devices</h1>
        <p className="mt-1 text-sm text-gray-500">Manage registered devices and their trust scores</p>
      </div>

      {loading && <Loading />}
      {error && <ErrorDisplay message={error} onRetry={fetchDevices} />}
      {!loading && !error && <Table columns={columns} data={devices} />}
    </div>
  );
}
