import * as React from "react";

import { ViewProps, PaginatedListProps } from "../src/";

interface Props {
  [key: string]: any;
}
interface Defaultable<T extends Props = {}> {
  default: T;
}
interface Loadable<T extends Props = {}> extends Defaultable<T> {
  loading: T;
  disabled: T;
  loadingDisabled: T;
}
interface Paginable<T extends Props = {}> extends Defaultable<T> {
  hasNextPage: T;
  hasPreviousPage: T;
  hasNoPages: T;
}
interface Transactable<T extends Props = {}> extends Defaultable<T> {
  loading: T;
  success: T;
  error: T;
}

export const loadable: Loadable<ViewProps> = {
  default: {
    disabled: false,
    loading: false
  },
  loading: {
    disabled: false,
    loading: true
  },
  disabled: {
    disabled: true,
    loading: false
  },
  loadingDisabled: {
    disabled: true,
    loading: true
  }
};
export const paginable: Paginable<PaginatedListProps> = {
  default: {
    pageInfo: {
      hasNextPage: true,
      hasPreviousPage: true
    },
    onNextPage: () => {},
    onPreviousPage: () => {},
    onRowClick: () => () => {}
  },
  hasNextPage: {
    pageInfo: {
      hasNextPage: true,
      hasPreviousPage: false
    },
    onNextPage: () => {},
    onPreviousPage: () => {},
    onRowClick: () => () => {}
  },
  hasPreviousPage: {
    pageInfo: {
      hasNextPage: false,
      hasPreviousPage: true
    },
    onNextPage: () => {},
    onPreviousPage: () => {},
    onRowClick: () => () => {}
  },
  hasNoPages: {
    pageInfo: {
      hasNextPage: false,
      hasPreviousPage: false
    },
    onNextPage: () => {},
    onPreviousPage: () => {},
    onRowClick: () => () => {}
  }
};

export const transactable: Transactable<{}> = {
  default: {
    transaction: "default" as "default"
  },
  loading: {
    transaction: "loading" as "loading"
  },
  success: {
    transaction: "success" as "success"
  },
  error: {
    transaction: "error" as "error"
  }
};

export function makeProps<T, P>(a: T, b: P): T & P {
  const output = {};
  const interfaces = [a, b];
  interfaces.forEach((i, index) => {
    Object.keys(i).forEach(k => {
      output[k] = {
        ...interfaces[(index + 1) % 2]["default"],
        ...interfaces[index][k]
      };
    });
  });
  return output as T & P;
}

export function makeStories(stories: any, props: any, Component: any) {
  Object.keys(props).forEach(propKey => {
    stories.add(propKey, () => <Component {...props[propKey]} />);
  });
}
export const listViewProps = makeProps(makeProps(loadable, paginable), {
  default: { onAdd: () => {} }
});
export const formViewProps = makeProps(makeProps(loadable, transactable), {
  default: { onBack: () => {}, onSubmit: () => {} }
});
