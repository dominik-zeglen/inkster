const breakpoints = {
  xs: 576,
  sm: 768,
  md: 992,
  lg: 1200
};

type Breakpoint = "xs" | "sm" | "md" | "lg" | "xl" | string;

function makeColor(main: string, dark: string) {
  return {
    main,
    dark
  }
}
const palette = {
  red: makeColor('#fc5c65', '#eb3b5a'),
  orange: makeColor('#fd9644', '#fa8231'),
  yellow: makeColor('#fed330', '#f7b731'),
  green: makeColor('#2bcbba', '#0fb9b1'),
  blue: makeColor('#45aaf2', '#2d98da'),
  purple: makeColor('#a55eea', '#8854d0'),
  gray: makeColor('#778ca3', '#4b6584'),
  lightGray: makeColor('#d1d8e0', '#a5b1c2'),
  black: makeColor('#2e2e2e', '#212121'),
  white: makeColor('#fdfdfd', '#f6f6f6')
}
const colors = {
  primary: palette.blue,
  secondary: palette.green,
  success: palette.green,
  error: palette.red,
  disabled: palette.lightGray.dark,
  ...palette
};
const baseTypography = {
  color: colors.black.main,
  lineHeight: "1.42857143",
  fontSize: "1rem",
  fontFamily: '"Open Sans", sans-serif',
  fontWeight: 400
};
export const theme = {
  colors,
  spacing: 10,
  breakpoints: {
    down: (bp: Breakpoint) => {
      return `@media (max-width: ${breakpoints[bp] - 0.02}px)`;
    },
    up: (bp: Breakpoint) => {
      return `@media (min-width: ${breakpoints[bp]}px)`;
    },
    width: (bp: Breakpoint) => breakpoints[bp]
  },
  typography: {
    body: {
      ...baseTypography
    },
    mainHeading: {
      ...baseTypography,
      fontSize: "1.953rem"
    },
    subHeading: {
      ...baseTypography,
      fontSize: "1.563rem"
    },
    button: {
      ...baseTypography,
      fontSize: "0.8rem",
      fontWeight: 600,
      textTransform: "uppercase"
    },
    caption: {
      ...baseTypography,
      fontSize: "0.8rem"
    }
  },
  transition: {
    time: "300ms"
  }
};
export default theme;
