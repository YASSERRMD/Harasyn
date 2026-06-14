'use client';

import React, { useState, useEffect, useRef } from 'react';

interface CommandPaletteProps {
  isOpen: boolean;
  onClose: () => void;
  commands: Command[];
}

interface Command {
  id: string;
  name: string;
  description: string;
  action: () => void;
  category: string;
}

export default function CommandPalette({ isOpen, onClose, commands }: CommandPaletteProps) {
  const [query, setQuery] = useState('');
  const [selectedIndex, setSelectedIndex] = useState(0);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (isOpen) {
      inputRef.current?.focus();
    }
  }, [isOpen]);

  const filteredCommands = commands.filter(cmd =>
    cmd.name.toLowerCase().includes(query.toLowerCase()) ||
    cmd.description.toLowerCase().includes(query.toLowerCase())
  );

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Escape') {
      onClose();
    } else if (e.key === 'ArrowDown') {
      e.preventDefault();
      setSelectedIndex(prev => Math.min(prev + 1, filteredCommands.length - 1));
    } else if (e.key === 'ArrowUp') {
      e.preventDefault();
      setSelectedIndex(prev => Math.max(prev - 1, 0));
    } else if (e.key === 'Enter') {
      filteredCommands[selectedIndex]?.action();
      onClose();
    }
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 overflow-y-auto">
      <div className="flex min-h-screen items-start justify-center px-4 pt-4 pb-20 text-center sm:block sm:p-0">
        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" onClick={onClose} />
        <div className="inline-block w-full max-w-2xl transform overflow-hidden rounded-lg bg-white text-left align-bottom shadow-xl transition-all sm:my-8 sm:align-middle">
          <div className="p-4">
            <input
              ref={inputRef}
              type="text"
              className="w-full border-0 border-b border-gray-200 px-0 py-3 text-lg focus:ring-0 focus:outline-none"
              placeholder="Type a command..."
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              onKeyDown={handleKeyDown}
            />
            <div className="mt-4 max-h-96 overflow-y-auto">
              {filteredCommands.map((cmd, idx) => (
                <div
                  key={cmd.id}
                  className={`flex items-center justify-between px-4 py-3 cursor-pointer ${
                    idx === selectedIndex ? 'bg-primary-50' : 'hover:bg-gray-50'
                  }`}
                  onClick={() => {
                    cmd.action();
                    onClose();
                  }}
                >
                  <div>
                    <div className="font-medium text-gray-900">{cmd.name}</div>
                    <div className="text-sm text-gray-500">{cmd.description}</div>
                  </div>
                  <span className="text-xs text-gray-400">{cmd.category}</span>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
