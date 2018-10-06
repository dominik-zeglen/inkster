import * as ColorFunc from "color";
import { Breakpoint, Color, Theme, Typography } from "react-jss";

const breakpoints = {
  xs: 576,
  sm: 768,
  md: 992,
  lg: 1280,
};

function makeColor(
  main: string,
  dark?: string,
  light?: string,
  lightest?: string,
): Color {
  return {
    main,
    dark:
      dark ||
      ColorFunc(main)
        .darken(0.2)
        .rgb()
        .string(),
    light:
      light ||
      ColorFunc(main)
        .lighten(0.2)
        .rgb()
        .string(),
    lightest:
      lightest ||
      ColorFunc(main)
        .lighten(0.2)
        .whiten(0.85)
        .rgb()
        .string(),
  };
}
const palette = {
  red: makeColor("#fc5c65", "#eb3b5a", undefined, "#fff1f1"),
  orange: makeColor("#fd9644", "#fa8231"),
  yellow: makeColor("#fed330", "#f7b731"),
  green: makeColor("#2bcbba", "#0fb9b1"),
  blue: makeColor("#45aaf2", "#2d98da"),
  purple: makeColor("#a55eea", "#8854d0"),
  gray: makeColor("#778ca3", "#4b6584"),
  lightGray: makeColor("#d1d8e0", "#a5b1c2"),
  black: makeColor("#2e2e2e", "#212121"),
  white: makeColor("#fdfdfd", "#f6f6f6"),
};
const colors = {
  primary: palette.blue,
  secondary: palette.green,
  success: palette.green,
  error: palette.red,
  disabled: palette.lightGray.dark,
  ...palette,
};
const baseTypography: Typography = {
  color: colors.black.main,
  lineHeight: "1.42857143",
  fontSize: "1rem",
  fontFamily: '"Open Sans", sans-serif',
  fontWeight: 400,
};
export const theme: Theme = {
  colors,
  spacing: 10,
  breakpoints: {
    down: (bp: Breakpoint) => {
      return `@media (max-width: ${breakpoints[bp] - 0.02}px)`;
    },
    up: (bp: Breakpoint) => {
      return `@media (min-width: ${breakpoints[bp]}px)`;
    },
    width: (bp: Breakpoint) => breakpoints[bp],
  },
  typography: {
    anchor: {
      ...baseTypography,
      color: colors.primary.main,
      display: "inline-block" as "inline-block",
      fontSize: "0.8rem",
      transitionDuration: "500ms",
      "&:hover, &:focus": {
        color: colors.secondary.dark,
      },
      "&:after": {
        background: colors.primary.main,
        content: "''",
        display: "block" as "block",
        height: 1,
        maxWidth: 20,
        width: "100%",
        transitionDuration: "500ms",
      },
      "&:hover:after, &:focus:after": {
        background: colors.secondary.dark,
        maxWidth: 40,
      },
    },
    body: {
      ...baseTypography,
    },
    mainHeading: {
      ...baseTypography,
      fontSize: "1.953rem",
    },
    subHeading: {
      ...baseTypography,
      fontSize: "1.563rem",
    },
    button: {
      ...baseTypography,
      fontSize: "0.8rem",
      fontWeight: 600 as 600,
      textTransform: "uppercase" as "uppercase",
    },
    caption: {
      ...baseTypography,
      fontSize: "0.8rem",
    },
  },
  transition: {
    time: "500ms",
  },
};
export default theme;
