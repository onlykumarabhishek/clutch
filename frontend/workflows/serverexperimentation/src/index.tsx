import type { DefaultWorkflowConfig } from "@clutch-sh/core";

import { StartExperiment } from "./start-experiment";

const register = (): DefaultWorkflowConfig => {
  return {
    developer: {
      name: "Lyft",
      contactUrl: "mailto:hello@clutch.sh",
    },
    path: "server-experimentation",
    group: "Chaos Experimentation",
    displayName: "Server Fault Injection",
    routes: {
      startExperiment: {
        path: "start",
        displayName: "Start Experiment",
        description: "Start Server Experiment.",
        component: StartExperiment,
      },
    },
  };
};

export default register;
