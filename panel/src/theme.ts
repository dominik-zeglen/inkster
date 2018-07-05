const breakpoints = {
  xs: 576,
  sm: 768,
  md: 992,
  lg: 1200
};

type Breakpoint = "xs" | "sm" | "md" | "lg" | "xl" | string;

const colors = {
  primary: {
    main: "#4285f4",
    dark: "#0d47a1"
  },
  secondary: {
    main: "#aa66cc",
    dark: "#9933cc"
  },
  success: {
    main: "#00C851",
    dark: "#007E33"
  },
  error: {
    main: "#ff4444",
    dark: "#CC0000"
  },
  black: {
    main: "#2e2e2e",
    dark: "#212121"
  },
  white: {
    main: "#FDFDFD",
    dark: "#F6F6F6"
  },
  disabled: "#bdbdbd"
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
