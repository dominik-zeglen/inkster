declare module 'react-jss' {

  import * as React from 'react';

  export interface CSSProperties extends React.CSSProperties {
    [key: string]: React.CSSProperties | string | CSSProperties | any;
    composes?: string | string[];
  }

  export type StyleSheet<Props = {}>
    = Record<
        string, 
        CSSProperties 
        | ((props: Props) => React.CSSProperties)
        | Record<string, CSSProperties>
      >;

  type StyleRules<ClassKey extends string = string, Props = {}> 
    = Record<ClassKey, CSSProperties 
    | ((props: Props) => React.CSSProperties)>;

  export type ClassNameMap<ClassKey extends string = string> = Record<ClassKey, string>;

  export interface WithStyles<ClassKey extends string = string> {
    classes: ClassNameMap<ClassKey>
  }

  export interface StyledComponentProps<ClassKey extends string = string> {
    classes?: Partial<ClassNameMap<ClassKey>>
  }

  function injectSheet<ClassKey extends string>(
    style: StyleRules<ClassKey> | ((theme: any) => StyleRules<ClassKey>),
    options?: any,
  ): <P>(
      component: React.ComponentType<P & WithStyles<ClassKey>>,
    ) => React.ComponentType<P & StyledComponentProps<ClassKey>>;

  export default injectSheet;

  export const jss: any

  export const JssProvider: any

  export const ThemeProvider: any
}
