const breakpoints = {
  xs: 576,
  sm: 768,
  md: 992,
  lg: 1200
};

type Breakpoint = "xs" | "sm" | "md" | "lg" | "xl" | string;

export const theme = {
  colors: {
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
    disabled: "#bdbdbd"
  },
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
      fontSize: "1rem"
    },
    mainHeading: {
      fontSize: "1.75rem",
      fontWeight: 500
    }
  },
  transition: {
    time: "300ms"
  }
};
export default theme;
