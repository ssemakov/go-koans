"use client";

import { useLayoutEffect, useState } from "react";
import { parseTheme, type Theme, themeStorageKey } from "@/lib/theme";

export default function ThemeSelect() {
  const [theme, setTheme] = useState<Theme>("system");

  useLayoutEffect(() => {
    let saved: Theme = "system";
    try { saved = parseTheme(localStorage.getItem(themeStorageKey)); }
    catch { /* System colors remain available when storage is blocked. */ }
    document.documentElement.dataset.theme = saved;
    setTheme(saved);
  }, []);

  function changeTheme(value: string) {
    const next = parseTheme(value);
    setTheme(next);
    document.documentElement.dataset.theme = next;
    try { localStorage.setItem(themeStorageKey, next); }
    catch { /* The selected theme still works for this visit. */ }
  }

  return (
    <label className="theme-select">Theme
      <select value={theme} onChange={(event) => changeTheme(event.target.value)}>
        <option value="system">System</option>
        <option value="light">Light</option>
        <option value="dark">Dark</option>
      </select>
    </label>
  );
}
