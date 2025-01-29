import type { Theme } from "./lib/Theme";

export const themes: Theme[] = [
  {
    name: "Fringo Light",
    backgroundColor: "#aaaaaa",
    textColor: "#000000",
    grid: {
      color: "#000000",
    },
    field: {
      default: {
        color: "#ffffff",
        hoverColor: "#d5d5d5",
      },
      checked: {
        color: "#a1ff84",
        hoverColor: "#72cf56",
      },
      gay: {
        color:
          "conic-gradient(red 0%, orange 14%, yellow 28%, green 42%, blue 57%, indigo 71%, violet 85%, red 100%) 50% 50%",
        hoverColor:
          "conic-gradient(red 0%, orange 14%, yellow 28%, green 42%, blue 57%, indigo 71%, violet 85%, red 100%) 50% 50%",
        textOutlineColor: "#ff77ed",
      },
    },
    button: {
      color: "#ffffff",
      hoverColor: "#eeeeee",
    },
    sidebar: {
      color: "#7e7e7eaa",
      backgroundBlur: "5px",
    },
  },

  {
    name: "Fringo Dark",
    backgroundColor: "#151515",
    textColor: "#ffffff",
    grid: {
      color: "#333333",
    },

    field: {
      default: {
        color: "#111111",
        hoverColor: "#222222",
      },
      checked: {
        color: "#08740c",
        hoverColor: "#0d9912",
      },
      gay: {
        color:
          "conic-gradient(darkred 0%, darkorange 14%, rgb(142, 142, 0) 28%, darkgreen 42%, darkblue 57%, indigo 71%, darkviolet 85%, darkred 100%) 50% 50%",
        hoverColor:
          "conic-gradient(darkred 0%, darkorange 14%, rgb(142, 142, 0) 28%, darkgreen 42%, darkblue 57%, indigo 71%, darkviolet 85%, darkred 100%) 50% 50%",
        textOutlineColor: "#96378a",
      },
    },
    button: {
      color: "#444444",
      hoverColor: "#555555",
    },
    sidebar: {
      color: "#212121aa",
      backgroundBlur: "5px",
    },
  },

  {
    name: "Unicorn Poop",
    backgroundColor:
      "linear-gradient(-45deg,rgb(136, 50, 255),rgb(255, 73, 143))",
    textColor: "#ffffff",
    grid: {
      color: "#7322ff",
    },
    field: {
      default: {
        color: "#ff87e4",
        hoverColor: "#ff7ce2",
      },
      checked: {
        color: "#ff00f5",
        hoverColor: "#ff5af9",
      },
      gay: {
        color:
          "conic-gradient(red 0%, orange 14%, yellow 28%, green 42%, blue 57%, indigo 71%, violet 85%, red 100%) 50% 50%",
        hoverColor:
          "conic-gradient(red 0%, orange 14%, yellow 28%, green 42%, blue 57%, indigo 71%, violet 85%, red 100%) 50% 50%",
        textOutlineColor: "#ca00ff",
      },
    },
    button: {
      color: "#ff75e0",
      hoverColor: "#ff8ee6",
    },
    sidebar: {
      color: "#9b277faa",
      backgroundBlur: "5px",
    },
  },
];
