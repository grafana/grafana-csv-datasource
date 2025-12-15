import { defineConfig } from 'eslint/config';
import baseConfig from './.config/eslint.config.mjs';
import prettier from 'eslint-plugin-prettier';

export default defineConfig([
  {
    ignores: [
      '**/logs',
      '**/*.log',
      '**/npm-debug.log*',
      '**/yarn-debug.log*',
      '**/yarn-error.log*',
      '**/.eslintcache',
      '**/node_modules/',
      '**/pids',
      '**/*.pid',
      '**/*.seed',
      '**/*.pid.lock',
      '**/lib-cov',
      '**/coverage',
      '**/dist/',
      '**/artifacts/',
      '**/work/',
      '**/ci/',
      '**/e2e-results/',
      '**/test-results/',
      '**/playwright-report/',
      '**/blob-report/',
      'playwright/.cache/',
      'playwright/.auth/',
      '**/.idea',
      '**/mage_output_file.go',
      '**/.pnp.*',
      '.yarn/*',
      '!.yarn/patches',
      '!.yarn/plugins',
      '!.yarn/releases',
      '!.yarn/sdks',
      '!.yarn/versions',
    ],
  },
  ...baseConfig,
  {
    plugins: {
      prettier: prettier,
    },

    rules: {
      'prettier/prettier': 'error',
    },
  },
]);
