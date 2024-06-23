// https://eslint.org/docs/latest/use/configure/configuration-files
// JavaScript (ESM) - use .eslintrc.cjs when running ESLint in JavaScript packages that specify 'type':'module' in their package.json. Note that ESLint does not support ESM configuration at this time.
module.exports = {
    env: {
        node: true,
    },
    extends: [
        'plugin:sonarjs/recommended-legacy',
    ],
    overrides: [{
        files: ['*.ts', '*.mts', '*.cts', '*.tsx'],
        rules: {
            '@typescript-eslint/explicit-function-return-type': 'error',
            '@typescript-eslint/explicit-member-accessibility': [
                'error',
                {
                    accessibility: 'explicit',
                    overrides: {
                        constructors: 'no-public',
                    },
                },
            ],
        },
    }],
    parser: '@typescript-eslint/parser',
    plugins: [
        '@typescript-eslint',
        'typescript-sort-keys',
        'import',
        'prettier',
        'sonarjs',
        'unicorn',
    ],
    rules: {
        '@typescript-eslint/explicit-function-return-type': 'error',
        '@typescript-eslint/explicit-member-accessibility': 'error',
        // https://github.com/typescript-eslint/typescript-eslint/issues/2483#issuecomment-687095358
        '@typescript-eslint/no-shadow': 'error',
        '@typescript-eslint/no-unused-vars': [
            'error',
            {
                argsIgnorePattern: '^_',
                varsIgnorePattern: '^_',
            },
        ],
        'eqeqeq': 'error',
        'import/order': [
            'error',
            {
                'alphabetize': { order: 'asc' },
                'newlines-between': 'always',
            },
        ],
        'no-console': 'error',
        'no-shadow': 'off',
        'padding-line-between-statements': [
            'error',
            { blankLine: 'always', next: 'return', prev: '*' },
        ],
        'prettier/prettier': 'error',
        'semi': 'off', // Otherwise it conflicts with prettier
        'sort-imports': ['error', { ignoreDeclarationSort: true }],
        'sort-keys': 'error',
        'typescript-sort-keys/interface': 'error',
        'typescript-sort-keys/string-enum': 'error',
        'unicorn/numeric-separators-style': 'error',
    },
};
