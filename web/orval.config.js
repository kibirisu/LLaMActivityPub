import { defineConfig } from "orval";

export default defineConfig({
  api: {
    input: {
      target: "../doc/openapi.yaml",
    },

    output: {
      mode: "tags-split",
      target: "src/api/api.gen.ts",
      schemas: "src/model",
      client: "react-query",
    },
  },
});
