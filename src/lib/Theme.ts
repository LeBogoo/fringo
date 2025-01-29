export interface ButtonColors {
  color: string;
  hoverColor: string;
  textColor?: string;
  textOutlineColor?: string;
}

export interface Theme {
  name: string;
  backgroundColor: string;
  textColor: string;
  grid: {
    color: string;
  };
  field: {
    default: ButtonColors;
    checked: ButtonColors;
    disabled: ButtonColors;
    gay: ButtonColors;
  };
  button: ButtonColors;
  sidebar: {
    color: string;
    backgroundBlur: string;
  };
}
