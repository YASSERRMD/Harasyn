'use client';

import { useEffect, useState } from 'react';
import Table from '@/components/Table';
import StatusBadge from '@/components/StatusBadge';
import Loading from '@/components/Loading';
import ErrorDisplay from '@/components/ErrorDisplay';
import { Session, api } from '@/lib/api';

const columns = [
  { key: 'user_id', header: 'User' },
  { key: 'resource_id', header: 'Resource' },
  { key: 'status', header: 'Status', render: (item: Session) => <StatusBadge status={item.status} /> },
  { key: 'risk_score', header: 'Risk Score', render: (item: Session) => (
    <span className={`font-medium ${item.risk_score <= 30 ? 'text-green-600' : item.risk_score <= 60 ? 'text-yellow-600' : 'text-red-600'}`}>
      {item.risk_score}
    </span>
  )},
  { key: 'granted_at', header: 'Granted', render: (item: Session) => new Date(item.granted_at).toLocaleString() },
  { key: 'expires_at', header: 'Expires', render: (item: Session) => new Date(item.expires_at).toLocaleString() },
];

export default function SessionsPage() {
  const [sessions, setSessions] = useState<Session[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchSessions = async () => {
    try {
      setLoading(true);
      const data = await api.getSessions('default');
      setSessions(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch sessions');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchSessions();
  }, []);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Sessions</h1>
        <p className="mt-1 text-sm text-gray-500">View and manage active access sessions</p>
      </div>

      {loading && <Loading />}
      {error && <ErrorDisplay message={error} onRetry={fetchSessions} />}
      {!loading && !error && <Table columns={columns} data={sessions} />}
    </div>
  );
}
