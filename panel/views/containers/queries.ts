import gql from "graphql-tag";

export const getRootContainersQuery = gql`
  query GetRootContainers {
    getRootContainers {
      id
      name
    }
  }
`;

export const getContainerChildrenQuery = gql`
  query GetContainerChildren($id: Int!) {
    getContainer(id: $id) {
      id
      children {
        id
        name
      }
    }
  }
`;

export const getContainerDetailsQuery = gql`
  query GetContainerDetails($id: Int!) {
    getContainer(id: $id) {
      id
      name
    }
  }
`;

export const removeContainerMutation = gql`
  mutation RemoveContainer($id: Int!) {
    removeContainer(id: $id)
  }
`;
