#!/usr/bin/env node

// Esegue un comando usando pm dedotto dalla shell (npm/pnpm/yarn/bun).
// Eventualmente legge da .env.PACKAGE_MANAGER se impostato.
const { spawnSync } = require('node:child_process');
const { resolve } = require('node:path');

const detectPackageManager = () => {
  const override = (process.env.PACKAGE_MANAGER || '').trim();
  if (override) return override;

  const ua = process.env.npm_config_user_agent || '';
  if (ua.startsWith('pnpm')) return 'pnpm';
  if (ua.startsWith('yarn')) return 'yarn';
  if (ua.startsWith('bun')) return 'bun';
  if (ua.startsWith('npm')) return 'npm';
  return 'npm';
};

const splitCwd = (args) => {
  const index = args.indexOf('--cwd');
  if (index === -1) return { args, cwd: process.cwd() };

  const dir = args[index + 1];
  const cleaned = args.slice(0, index).concat(args.slice(index + 2));
  return { args: cleaned, cwd: dir ? resolve(process.cwd(), dir) : process.cwd() };
};

const [, , ...argv] = process.argv;
if (argv.length === 0) {
  console.error('Usage: node scripts/pm.js <command> [...args] [--cwd path]');
  process.exit(1);
}

const { args, cwd } = splitCwd(argv);
const pm = detectPackageManager();
const result = spawnSync(pm, args, { stdio: 'inherit', cwd, env: process.env });

if (result.error) {
  console.error(`[pm] Failed to run ${pm}: ${result.error.message}`);
  process.exit(1);
}

process.exit(typeof result.status === 'number' ? result.status : 1);
