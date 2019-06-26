import * as ColorFunc from "color";
import baseAuroraTheme from "aurora-ui-kit/dist/theme";
import { IColor } from "aurora-ui-kit/dist/theme/palette";

function makeColor(
  main: string,
  dark?: string,
  light?: string,
  lightest?: string,
): IColor {
  return {
    main,
    dark:
      dark ||
      ColorFunc(main)
        .darken(0.2)
        .rgb()
        .string(),
    darkest:
      dark ||
      ColorFunc(main)
        .darken(0.2)
        .blacken(0.85)
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
  secondary: palette.purple,
  success: palette.green,
  error: palette.red,
  disabled: palette.lightGray.dark,
};

export const theme: typeof baseAuroraTheme = {
  ...baseAuroraTheme,
  colors: {
    ...baseAuroraTheme.colors,
    ...colors,
    background: {
      ...baseAuroraTheme.colors.background,
      default: ColorFunc(colors.primary.main)
        .alpha(0.02)
        .rgb()
        .string(),
    },
  },
};

export default theme;
