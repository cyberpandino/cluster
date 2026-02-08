import js from '@eslint/js';
import globals from 'globals';
import { defineConfig } from 'eslint/config';
import stylistic from '@stylistic/eslint-plugin';

export default defineConfig([
  { 
    files: ['**/*.{js,mjs,cjs}'], 
    plugins: { 
      js, 
      '@stylistic': stylistic 
    }, 
    extends: ['js/recommended'], 
    languageOptions: { 
      globals: globals.node 
    }, 
    rules: {
      'camelcase': ['warn'],      
      'no-unused-vars': ['off'],
      'camelcase': [
        'error', { 
          'properties': 'always' 
        }
      ],
      'new-cap': [
        'error', { 
          'newIsCap': true, 
          'capIsNew': false 
        }
      ],
      'id-match': [
        'error',
        '^([a-z][a-zA-Z0-9]*|[A-Z][A-Z0-9_]*)$',
        {
          'onlyDeclarations': true,
          'properties': false
        }
      ],

      '@stylistic/indent': ['warn', 2],
      '@stylistic/quotes': ['warn', 'single'],
    }
  },
]);