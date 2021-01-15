import React, { Component } from 'react';

import { ProjectType, ClusterType } from '../shared/types';

type PropsType = {
}

type StateType = {
}

const Context = React.createContext({});

const { Provider } = Context;
const ContextConsumer = Context.Consumer;

/**
 * Component managing a universal (application-wide) data store.
 *
 * Important Usage Notes:
 * 1) Each field must have an accompanying setter
 * 2) No function calls are allowed from within Context (not counting
 *    initialization)
 * 3) Context should be used as a last-resort (changes will re-render ALL
 *    components consuming Context)
 * 4) As a rule of thumb, Context should not be used for UI-related state
 */
class ContextProvider extends Component {
  state = {
    currentModal: null as string | null,
    currentModalData: null as any,
    setCurrentModal: (currentModal: string, currentModalData?: any) => {
      this.setState({ currentModal, currentModalData });
    },
    currentError: null as string | null,
    setCurrentError: (currentError: string) => {
      this.setState({ currentError });
    },
    currentCluster: null as ClusterType | null,
    setCurrentCluster: (currentCluster: ClusterType) => {
      localStorage.setItem('currentCluster', currentCluster.toString());
      this.setState({ currentCluster });
    },
    currentProject: null as ProjectType | null,
    setCurrentProject: (currentProject: ProjectType) => {
      this.setState({ currentProject });
    },
    projects: [] as ProjectType[],
    setProjects: (projects: ProjectType[]) => {
      this.setState({ projects });
    },
    user: null as any,
    setUser: (userId: number, email: string) => {
      this.setState({ user: {userId, email} });
    },
    devOpsMode: true,
    setDevOpsMode: (devOpsMode: boolean) => {
      this.setState({ devOpsMode });
    },
    clearContext: () => {
      this.setState({
        currentModal: null,
        currentModalData: null,
        currentError: null,
        currentCluster: null,
        currentProject: null,
        projects: [],
        user: null,
        devOpsMode: true,
      });
    }
  };
  
  render() {
    return (
      <Provider value={this.state}>{this.props.children}</Provider>
    );
  }
}

export { Context, ContextProvider, ContextConsumer };
