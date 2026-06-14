'use client';

import React, { useState } from 'react';

interface PolicyRule {
  id: string;
  resource: string;
  action: string;
  conditions: Condition[];
  effect: 'allow' | 'deny';
}

interface Condition {
  type: string;
  operator: string;
  value: string;
}

interface PolicyBuilderProps {
  onSave: (policy: PolicyRule) => void;
}

export default function PolicyBuilder({ onSave }: PolicyBuilderProps) {
  const [policy, setPolicy] = useState<PolicyRule>({
    id: '',
    resource: '',
    action: '',
    conditions: [],
    effect: 'allow',
  });

  const addCondition = () => {
    setPolicy({
      ...policy,
      conditions: [
        ...policy.conditions,
        { id: Date.now().toString(), type: 'identity', operator: 'equals', value: '' },
      ],
    });
  };

  const updateCondition = (id: string, field: keyof Condition, value: string) => {
    setPolicy({
      ...policy,
      conditions: policy.conditions.map(c =>
        c.id === id ? { ...c, [field]: value } : c
      ),
    });
  };

  const removeCondition = (id: string) => {
    setPolicy({
      ...policy,
      conditions: policy.conditions.filter(c => c.id !== id),
    });
  };

  return (
    <div className="bg-white rounded-lg shadow p-6">
      <h3 className="text-lg font-semibold text-gray-900 mb-4">Policy Builder</h3>
      
      <div className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700">Resource</label>
          <input
            type="text"
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
            value={policy.resource}
            onChange={(e) => setPolicy({ ...policy, resource: e.target.value })}
            placeholder="e.g., api.production"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700">Action</label>
          <select
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
            value={policy.action}
            onChange={(e) => setPolicy({ ...policy, action: e.target.value })}
          >
            <option value="">Select action</option>
            <option value="read">Read</option>
            <option value="write">Write</option>
            <option value="delete">Delete</option>
            <option value="execute">Execute</option>
          </select>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700">Effect</label>
          <div className="mt-1 flex space-x-4">
            <label className="flex items-center">
              <input
                type="radio"
                className="focus:ring-primary-500 h-4 w-4 text-primary-600 border-gray-300"
                checked={policy.effect === 'allow'}
                onChange={() => setPolicy({ ...policy, effect: 'allow' })}
              />
              <span className="ml-2 text-sm text-gray-700">Allow</span>
            </label>
            <label className="flex items-center">
              <input
                type="radio"
                className="focus:ring-primary-500 h-4 w-4 text-primary-600 border-gray-300"
                checked={policy.effect === 'deny'}
                onChange={() => setPolicy({ ...policy, effect: 'deny' })}
              />
              <span className="ml-2 text-sm text-gray-700">Deny</span>
            </label>
          </div>
        </div>

        <div>
          <div className="flex items-center justify-between mb-2">
            <label className="block text-sm font-medium text-gray-700">Conditions</label>
            <button
              type="button"
              onClick={addCondition}
              className="text-sm text-primary-600 hover:text-primary-500"
            >
              + Add Condition
            </button>
          </div>
          
          <div className="space-y-3">
            {policy.conditions.map((condition) => (
              <div key={condition.id} className="flex items-center space-x-3 p-3 bg-gray-50 rounded-md">
                <select
                  className="rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
                  value={condition.type}
                  onChange={(e) => updateCondition(condition.id, 'type', e.target.value)}
                >
                  <option value="identity">Identity</option>
                  <option value="device">Device</option>
                  <option value="location">Location</option>
                  <option value="time">Time</option>
                  <option value="risk">Risk Score</option>
                </select>
                <select
                  className="rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
                  value={condition.operator}
                  onChange={(e) => updateCondition(condition.id, 'operator', e.target.value)}
                >
                  <option value="equals">Equals</option>
                  <option value="not_equals">Not Equals</option>
                  <option value="contains">Contains</option>
                  <option value="greater_than">Greater Than</option>
                  <option value="less_than">Less Than</option>
                </select>
                <input
                  type="text"
                  className="flex-1 rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm"
                  value={condition.value}
                  onChange={(e) => updateCondition(condition.id, 'value', e.target.value)}
                  placeholder="Value"
                />
                <button
                  type="button"
                  onClick={() => removeCondition(condition.id)}
                  className="text-red-600 hover:text-red-500"
                >
                  Remove
                </button>
              </div>
            ))}
          </div>
        </div>

        <div className="flex justify-end">
          <button
            type="button"
            onClick={() => onSave(policy)}
            className="ml-3 inline-flex justify-center rounded-md border border-transparent bg-primary-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2"
          >
            Save Policy
          </button>
        </div>
      </div>
    </div>
  );
}
