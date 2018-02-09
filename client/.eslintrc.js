module.exports = {
	"parserOptions": {
    "ecmaVersion": 6,
		"sourceType": "module"
	},
	"env": {
		"commonjs": true,
		"es6": true,
		"browser": true,
		"node": true,
	},
	"extends": [
    "eslint:recommended",
    "plugin:vue/recommended"
  ],
  "plugins": [
    'vue'
  ],
	"rules": {
		"indent": ["error", 2],
		"quotes": ["error",	"single"],
		"semi": ["error",	"never"],
		"no-console": 0,
		"no-unused-vars": ["error"],
		"no-cond-assign": 0,
    "vue/max-attributes-per-line": [2, {
      "singleline": 4,
      "multiline": {
        "max": 1,
        "allowFirstLine": false
      }
    }],
    "vue/require-default-prop": 0,
	}
}
