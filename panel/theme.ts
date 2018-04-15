import "../assets/fonts/OpenSans-Bold.ttf";

interface TypographyType {
  color?: string;
  fontSize?: number | string;
  fontWeight?: number;
  fontFamily?: string;
  "-webkit-font-smoothing"?: string;
}

interface BorderType {
  borderRadius?: number | string;
  borderStyle?: string;
  borderWidth?: number | string;
}

interface ThemeType {
  spacing?: number;
  typography?: {
    default?: TypographyType;
    button?: TypographyType;
    disabled?: TypographyType;
  };
  colors?: {
    default?: string;
    accent?: string;
    success?: string;
    error?: string;
    disabled?: string;
    gradient?: string;
    gradientSecondary?: string;
  };
  borders: BorderType;
  components?: {
    [key: string]: {
      [key: string]: TypographyType | BorderType | any;
    };
  };
}

const colors = {
  accent: "#17c5ff",
  default: "#544aff",
  disabled: "#555555",
  gradient: "linear-gradient(4deg, #7800ff, #00f4ff)"
};
const defaultTypography: TypographyType = {
  "-webkit-font-smoothing": "antialiased",
  color: "#222222",
  fontFamily: "Open Sans, sans-serif",
  fontSize: 14,
  fontWeight: 400
};
const typography = {
  default: defaultTypography,
  disabled: {
    ...defaultTypography,
    color: colors.disabled
  }
};
const borders = {
  borderRadius: 2,
  borderStyle: "none",
  borderWidth: 0
};
export const theme: ThemeType = {
  borders,
  colors,
  components: {
    button: {
      "&:not(\\20)": {
        ...borders
      }
    },
    iconButton: {
      "&:not(\\20)": {
        borderRadius: borders.borderRadius
      }
    },
    listGroupItem: {
      "&:focus": {
        outline: "none"
      }
    },
    panel: {
      "&:not(\\20)": {
        "& .panel-footer": {
          backgroundColor: "#ffffff",
          borderTop: "none",
          display: "flex"
        },
        "& .panel-heading": {
          alignItems: "center",
          display: "flex"
        },
        borderRadius: borders.borderRadius
      }
    }
  },
  spacing: 10,
  typography
};
export default theme;
