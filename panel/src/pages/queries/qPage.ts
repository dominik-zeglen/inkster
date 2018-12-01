import gql from "graphql-tag";

import { TypedQuery } from "../../api";
import { Page, PageVariables } from "./types/Page";

const qPage = gql`
  query Page($id: ID!) {
    page(id: $id) {
      id
      createdAt
      updatedAt
      name
      slug
      isPublished
      fields {
        id
        name
        type
        value
      }
      parent {
        id
      }
    }
  }
`;
export default TypedQuery<Page, PageVariables>(qPage);
