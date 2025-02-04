import React from "react";
import { Grid, Paper, TextField, Toast, Typography, useNavigate } from "@clutch-sh/core";
import { Box, CircularProgress } from "@material-ui/core";
import SearchIcon from "@material-ui/icons/Search";

import type { WorkflowProps } from ".";
import catalogReducer from "./catalog-reducer";
import ProjectCard from "./project-card";
import { addProject, getProjects, removeProject } from "./storage";
import type { CatalogState } from "./types";

const initialState: CatalogState = {
  projects: [],
  search: "",
  isLoading: false,
  isSearching: false,
  error: undefined,
};

const Catalog: React.FC<WorkflowProps> = ({ heading }) => {
  const navigate = useNavigate();
  const [state, dispatch] = React.useReducer(catalogReducer, initialState);

  const navigateToProject = project => {
    navigate(`/catalog/${project.name}`);
  };

  const setError = err => dispatch({ type: "HYDRATE_ERROR", payload: { result: err.message } });

  React.useEffect(() => {
    dispatch({ type: "HYDRATE_START" });
    getProjects(
      projects => dispatch({ type: "HYDRATE_END", payload: { result: projects } }),
      setError
    );
  }, []);

  // TODO: Decouple some of the logic in the storage functions and migrate it to the reducer
  const triggerProjectAdd = () => {
    dispatch({ type: "SEARCH_START" });
    addProject(
      state.search,
      projects => {
        dispatch({ type: "ADD_PROJECT", payload: { projects } });
        dispatch({ type: "SEARCH_END" });
      },
      e => {
        dispatch({ type: "SEARCH_END" });
        setError(e);
      }
    );
  };

  const triggerProjectRemove = project => {
    removeProject(
      project.name,
      projects => dispatch({ type: "REMOVE_PROJECT", payload: { projects } }),
      setError
    );
  };

  return (
    <Box style={{ padding: "32px" }}>
      <div style={{ marginBottom: "8px" }}>
        <Typography variant="caption2">{heading} /</Typography>
      </div>
      <div style={{ marginBottom: "32px" }}>
        <Typography variant="h2">Project Catalog</Typography>
      </div>
      <Paper>
        <div style={{ margin: "16px" }}>
          <TextField
            placeholder="Search"
            value={state.search}
            onChange={e => dispatch({ type: "SEARCH", payload: { search: e.target.value } })}
            onKeyDown={e => e.key === "Enter" && triggerProjectAdd()}
            endAdornment={
              state.isSearching ? (
                <CircularProgress size="24px" />
              ) : (
                <SearchIcon onClick={triggerProjectAdd} />
              )
            }
          />
        </div>
      </Paper>
      <div style={{ marginBottom: "16px", marginTop: "32px" }}>
        <Typography variant="h3">My Projects</Typography>
      </div>
      <Grid container direction="row" spacing={5}>
        {state.projects.map(p => (
          <Grid item onClick={() => navigateToProject(p)}>
            <ProjectCard project={p} onRemove={() => triggerProjectRemove(p)} />
          </Grid>
        ))}
      </Grid>
      {state.error && (
        <Toast severity="error" onClose={() => dispatch({ type: "CLEAR_ERROR" })}>
          {state.error}
        </Toast>
      )}
    </Box>
  );
};

export default Catalog;
