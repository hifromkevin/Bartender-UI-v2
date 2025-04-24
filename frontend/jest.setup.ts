import '@testing-library/jest-dom';

global.importMetaEnv = {
  REACT_API_BASE_URL: 'http://localhost:3003',
};

Object.defineProperty(global, 'import', {
  value: {
    meta: {
      env: global.importMetaEnv,
    },
  },
});
