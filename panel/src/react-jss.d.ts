declare module "react-jss" {
  import * as React from "react";

  // Theme definition
  export type Breakpoint = "xs" | "sm" | "md" | "lg" | "xl" | string;
  export interface Color {
    dark: string;
    light: string;
    lightest: string;
    main: string;
  }

  export type Typography = {
    [T in keyof CSSProperties]?: CSSProperties[T]
  };

  export interface Theme {
    colors: {
      black: Color;
      disabled: string;
      error: Color;
      gray: Color;
      green: Color;
      lightGray: Color;
      orange: Color;
      primary: Color;
      purple: Color;
      red: Color;
      secondary: Color;
      success: Color;
      white: Color;
      yellow: Color;
    };
    spacing: number;
    breakpoints: {
      down: (bp: Breakpoint) => string;
      up: (bp: Breakpoint) => string;
      width: (bp: Breakpoint) => number;
    };
    typography: {
      anchor: Typography;
      body: Typography;
      mainHeading: Typography;
      subHeading: Typography;
      button: Typography;
      caption: Typography;
    };
    transition: {
      time: string;
    };
  }

  // Module definition
  export interface CSSProperties extends React.CSSProperties {
    [key: string]: React.CSSProperties | string | CSSProperties | any;
    composes?: string | string[];
  }

  export type StyleSheet<Props = {}> = Record<
    string,
    | CSSProperties
    | ((props: Props) => React.CSSProperties)
    | Record<string, CSSProperties>
  >;

  type StyleRules<ClassKey extends string = string, Props = {}> = Record<
    ClassKey,
    CSSProperties | ((props: Props) => React.CSSProperties)
  >;

  export type ClassNameMap<ClassKey extends string = string> = Record<
    ClassKey,
    string
  >;

  export interface WithStyles<ClassKey extends string = string> {
    classes: ClassNameMap<ClassKey>;
  }

  export interface StyledComponentProps<ClassKey extends string = string> {
    classes?: Partial<ClassNameMap<ClassKey>>;
  }

  function injectSheet<ClassKey extends string>(
    style: StyleRules<ClassKey> | ((theme: Theme) => StyleRules<ClassKey>),
    options?: any
  ): <P>(
    component: React.ComponentType<P & WithStyles<ClassKey>>
  ) => React.ComponentType<P & StyledComponentProps<ClassKey>>;

  export const jss: any;
  export const JssProvider: any;
  export const ThemeProvider: React.StatelessComponent<{
    theme: Theme;
  }>;

  export default injectSheet;
}
