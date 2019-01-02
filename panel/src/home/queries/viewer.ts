import gql from "graphql-tag";

import { TypedQuery, pageInfoFragment } from "../../api";
import { Viewer } from "./types/Viewer";

const viewer = gql`
  ${pageInfoFragment}
  query Viewer {
    viewer {
      id
      email
      pages(paginate: { first: 5 }, sort: { field: UPDATED_AT, order: DESC }) {
        edges {
          node {
            id
            name
            slug
            isPublished
          }
        }
        pageInfo {
          ...PageInfoFragment
        }
      }
    }
  }
`;
export const ViewerQuery = TypedQuery<Viewer, {}>(viewer);
