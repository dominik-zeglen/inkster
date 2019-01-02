import { Viewer_viewer } from "../../src/home/queries/types/Viewer";

export const viewer: Viewer_viewer = {
  __typename: "User",
  id: "dXNlcjox",
  email: "admin@example.com",
  pages: {
    __typename: "PageConnection",
    edges: [
      {
        __typename: "PageConnectionEdge",
        node: {
          __typename: "Page",
          id: "cGFnZTox",
          name: "How hairless apes evolved to build civilisation",
          slug: "how-apes-build-civilisation",
          isPublished: true,
        },
      },
    ],
    pageInfo: {
      __typename: "PageInfo",
      endCursor: "",
      hasNextPage: true,
      hasPreviousPage: false,
      startCursor: "",
    },
  },
};
