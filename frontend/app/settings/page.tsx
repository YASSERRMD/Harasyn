'use client';

export default function SettingsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Settings</h1>
        <p className="mt-1 text-sm text-gray-500">Configure your Harasyn platform settings</p>
      </div>

      <div className="bg-white shadow rounded-lg p-6">
        <h3 className="text-lg font-medium text-gray-900">General Settings</h3>
        <div className="mt-4 space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">Organization Name</label>
            <input type="text" className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-primary-500 focus:border-primary-500" defaultValue="Harasyn Corp" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Default Session Duration (minutes)</label>
            <input type="number" className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-primary-500 focus:border-primary-500" defaultValue="30" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">High Risk Threshold</label>
            <input type="number" className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-primary-500 focus:border-primary-500" defaultValue="70" />
          </div>
        </div>
        <div className="mt-6">
          <button className="px-4 py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700">
            Save Settings
          </button>
        </div>
      </div>
    </div>
  );
}
