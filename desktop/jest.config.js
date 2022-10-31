module.exports = {
  transformIgnorePatterns: [
    "/node_modules/.*", "/dist_electron/.*"
],
  testEnvironment: "jsdom",
  moduleNameMapper: {
    "@/(.+)$": "<rootDir>/src/$1",
  },
  collectCoverage: true,
  collectCoverageFrom: [
    'src/**/*.js',
    'src/**/**/*.vue',
  ],

  transform: {
    "\\.[jt]sx?$": "babel-jest",
    "^.+\\.vue$": "@vue/vue3-jest",
  },
};
