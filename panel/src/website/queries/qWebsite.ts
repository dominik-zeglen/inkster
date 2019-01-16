import gql from "graphql-tag";

import { TypedQuery } from "../../api";
import { Website } from "./types/Website";

const qWebsite = gql`
  query Website {
    website {
      name
      description
      domain
    }
  }
`;
export default TypedQuery<Website, {}>(qWebsite);
