import React from 'react';

interface StatusBadgeProps {
  status: string;
  variant?: 'success' | 'warning' | 'error' | 'info';
}

const variantStyles: Record<string, string> = {
  success: 'bg-green-100 text-green-800',
  warning: 'bg-yellow-100 text-yellow-800',
  error: 'bg-red-100 text-red-800',
  info: 'bg-blue-100 text-blue-800',
};

function getVariant(status: string): string {
  switch (status.toLowerCase()) {
    case 'active':
    case 'approved':
    case 'compliant':
    case 'success':
      return 'success';
    case 'pending':
    case 'partial':
      return 'warning';
    case 'revoked':
    case 'rejected':
    case 'expired':
    case 'non_compliant':
    case 'denied':
    case 'failure':
      return 'error';
    default:
      return 'info';
  }
}

export default function StatusBadge({ status, variant }: StatusBadgeProps) {
  const resolvedVariant = variant || getVariant(status);

  return (
    <span
      className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
        variantStyles[resolvedVariant]
      }`}
    >
      {status}
    </span>
  );
}
