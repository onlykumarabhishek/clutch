import type { BaseWorkflowProps, DefaultWorkflowConfig } from "@clutch-sh/core";

import HelloWorld from "./hello-world";

export interface WorkflowProps extends BaseWorkflowProps {}

const register = (): DefaultWorkflowConfig => {
  return {
    developer: {
      name: "{{ .DeveloperName }}",
      contactUrl: "mailto:{{ .DeveloperEmail }}",
    },
    path: "{{ .URLRoot }}",
    group: "{{ .Name }}",
    displayName: "{{ .Name }}",
    routes: {
      landing: {
        path: "{{ .URLPath }}",
        description: "{{ .Description }}.",
        component: HelloWorld,
      },
    },
  };
};

export default register;
