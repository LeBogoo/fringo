export interface ButtonColors {
  color: string;
  hoverColor: string;
  textColor?: string;
  textOutlineColor?: string;
}

export interface Theme {
  id: string;
  author: string;
  name: string;
  backgroundColor: string;
  textColor: string;
  grid: {
    color: string;
  };
  field: {
    default: ButtonColors;
    checked: ButtonColors;
    gay: ButtonColors;
  };
  button: ButtonColors;
  sidebar: {
    color: string;
    backgroundBlur: string;
  };
  input: {
    color: string;
    indicatorColor: string;
    activeColor: string;
  };
}
