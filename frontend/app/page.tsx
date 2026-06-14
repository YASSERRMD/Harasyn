import Link from 'next/link';

export default function HomePage() {
  return (
    <main className="flex-1 flex flex-col items-center justify-center p-8">
      <div className="max-w-3xl text-center">
        <h1 className="text-5xl font-bold text-gray-900 mb-6">
          Harasyn
        </h1>
        <p className="text-xl text-gray-600 mb-8">
          Zero Trust Access Platform
        </p>
        <p className="text-gray-500 mb-10 max-w-2xl mx-auto">
          A production-grade alternative to traditional VPNs, bastion hosts,
          and static network segmentation. Built on device trust, user trust,
          continuous authorization, and context-aware access policies.
        </p>
        <div className="flex flex-col sm:flex-row gap-4 justify-center mb-16">
          <Link
            href="/dashboard"
            className="px-6 py-3 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors"
          >
            Admin Console
          </Link>
          <Link
            href="/docs"
            className="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
          >
            Documentation
          </Link>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 text-left">
          <div className="p-6 bg-white rounded-xl border border-gray-200 shadow-sm">
            <h3 className="text-lg font-semibold text-gray-900 mb-2">Device Trust</h3>
            <p className="text-gray-500">Continuous posture evaluation with OS, encryption, and jailbreak detection.</p>
          </div>
          <div className="p-6 bg-white rounded-xl border border-gray-200 shadow-sm">
            <h3 className="text-lg font-semibold text-gray-900 mb-2">User Trust</h3>
            <p className="text-gray-500">Context-aware identity with MFA, location, and risk scoring.</p>
          </div>
          <div className="p-6 bg-white rounded-xl border border-gray-200 shadow-sm">
            <h3 className="text-lg font-semibold text-gray-900 mb-2">Policy Engine</h3>
            <p className="text-gray-500">Flexible Zero Trust policies with real-time evaluation and audit logging.</p>
          </div>
        </div>
        <div className="mt-12 text-sm text-gray-400">
          <p>Built with OpenCode and Kimi K2.6</p>
        </div>
      </div>
    </main>
  );
}
