'use client';

import { useEffect, useState } from 'react';
import Table from '@/components/Table';
import StatusBadge from '@/components/StatusBadge';
import Loading from '@/components/Loading';
import ErrorDisplay from '@/components/ErrorDisplay';
import { AuditEvent, api } from '@/lib/api';

const columns = [
  { key: 'event_type', header: 'Event Type' },
  { key: 'action', header: 'Action' },
  { key: 'actor_id', header: 'Actor' },
  { key: 'resource_type', header: 'Resource Type' },
  { key: 'status', header: 'Status', render: (item: AuditEvent) => <StatusBadge status={item.status} /> },
  { key: 'created_at', header: 'Timestamp', render: (item: AuditEvent) => new Date(item.created_at).toLocaleString() },
];

export default function AuditPage() {
  const [events, setEvents] = useState<AuditEvent[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchEvents = async () => {
    try {
      setLoading(true);
      const data = await api.getAuditEvents('default');
      setEvents(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch audit events');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchEvents();
  }, []);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Audit Log</h1>
        <p className="mt-1 text-sm text-gray-500">View security events and audit trail</p>
      </div>

      {loading && <Loading />}
      {error && <ErrorDisplay message={error} onRetry={fetchEvents} />}
      {!loading && !error && <Table columns={columns} data={events} />}
    </div>
  );
}
