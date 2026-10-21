export type Theme = "system" | "light" | "dark";
export const themeStorageKey = "go-koans:theme:v1";

export function parseTheme(value: string | null): Theme {
  return value === "light" || value === "dark" ? value : "system";
}

// Run before the body paints so a saved choice overrides the system colors immediately.
export const themeInitScript = `try{const theme=localStorage.getItem(${JSON.stringify(themeStorageKey)});if(theme==="light"||theme==="dark")document.documentElement.dataset.theme=theme;}catch{}`;
