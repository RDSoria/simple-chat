import React from "react";

interface Props {
  language: string;
  setLanguage: (lang: string) => void;
}

function LanguageSelector({ language, setLanguage }: Props) {
  return (
    <div className="p-4 bg-gray-200">
      <label className="mr-2">Select Language:</label>
      <select
        value={language}
        onChange={(e) => setLanguage(e.target.value)}
        className="p-2 border rounded-lg"
      >
        <option value="English">English</option>
        <option value="Spanish">Spanish</option>
      </select>
    </div>
  );
}

export default LanguageSelector;