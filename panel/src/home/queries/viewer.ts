import gql from "graphql-tag";

import { TypedQuery } from "../../api";
import { Viewer } from "./types/Viewer";

const viewer = gql`
  query Viewer {
    viewer {
      id
      email
      pages {
        id
        name
        slug
        isPublished
      }
    }
  }
`;
export const ViewerQuery = TypedQuery<Viewer, {}>(viewer);
