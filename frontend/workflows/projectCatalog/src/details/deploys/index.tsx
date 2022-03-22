import React from "react";
import type { ClutchError } from "@clutch-sh/core";
import { styled, Table, TableRow, Toast, Tooltip, Typography } from "@clutch-sh/core";
import { Card } from "@clutch-sh/project-selector";
import SecurityIcon from "@material-ui/icons/Security";

import { DefaultSummaryTitle, EventTime, setMilliseconds } from "../helpers";

import CommitInformation from "./commitInformation";
import EnvironmentIcon from "./environmentIcon";
import StatusIcon from "./statusIcon";
import type { DeployInfo, DeployJobInformation } from "./types";

export interface ProjectDeploysProps {
  fetchDeploys: () => Promise<DeployInfo>;
  singleProject?: boolean;
}

const Icon = styled("div")({
  fontSize: "24px",
  display: "flex",
  justifyContent: "center",
});

const StyledCard = styled(Card)({
  width: "100%",
  height: "100%",
});

const ProjectDeploys = ({ fetchDeploys, singleProject = true }: ProjectDeploysProps) => {
  const [error, setError] = React.useState<ClutchError | undefined>(undefined);
  const [deploys, setDeploys] = React.useState<DeployInfo>(undefined);
  const [isLoading, setIsLoading] = React.useState<boolean>(false);

  React.useEffect(() => {
    setIsLoading(true);
    fetchDeploys()
      .then((res: DeployInfo) => setDeploys(res))
      .catch((err: ClutchError) => {
        setError(err);

        // eslint-disable-next-line no-console
        console.error("Failed to fetch deploys", err.message);
      })
      .finally(() => setIsLoading(false));
  }, []);

  return (
    <>
      {error && <Toast>Failed to fetch Deploys</Toast>}
      <StyledCard
        avatar={<SecurityIcon />}
        title="Deploys"
        error={error}
        isLoading={isLoading}
        summary={[
          {
            title: deploys?.lastDeploy ? (
              <Typography variant="subtitle2">
                <EventTime date={setMilliseconds(deploys.lastDeploy)} />
              </Typography>
            ) : (
              <DefaultSummaryTitle />
            ),
            subheader: "Last deploy",
          },
          {
            title: deploys?.inProgress ? (
              <Typography variant="subtitle2" color="#3548D4">
                {deploys.inProgress}
              </Typography>
            ) : (
              <DefaultSummaryTitle />
            ),
            subheader: "In progress",
          },
          {
            title: deploys?.failures ? (
              <Typography variant="subtitle2" color="#DB3615">
                {deploys.failures}
              </Typography>
            ) : (
              <DefaultSummaryTitle />
            ),
            subheader: "Failed",
          },
        ]}
      >
        <Table columns={["", "", "", "", ""]} responsive>
          {deploys?.jobs?.length ? (
            deploys?.jobs?.map((job: DeployJobInformation) => {
              return (
                <TableRow key={job.name}>
                  {!singleProject && <div>{job.name}</div>}
                  <CommitInformation {...job.commitMetadata} />
                  <EventTime date={job.timestamp} />
                  <Tooltip title={job.status} placement="top">
                    <Icon>{StatusIcon(job.status)}</Icon>
                  </Tooltip>
                  <Tooltip title={job.environment} placement="top">
                    <Icon>{EnvironmentIcon(job.environment)}</Icon>
                  </Tooltip>
                </TableRow>
              );
            })
          ) : (
            <TableRow>
              <div>No deploys found for selected project(s)</div>
            </TableRow>
          )}
        </Table>
      </StyledCard>
    </>
  );
};

export default ProjectDeploys;
