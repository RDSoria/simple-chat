import React, { useState } from "react";
import ChatRoom from "./components/ChatRoom";
import LanguageSelector from "./components/LanguageSelector";

function App() {
  const [language, setLanguage] = useState<string>("EN");

  return (
    <div className="App">
      <h1 className="text-center text-xl font-bold mb-4">Chat Application</h1>
      <div className="flex">
        {/* First ChatRoom */}
        <div className="w-1/2 border-r border-gray-300">
          <ChatRoom user="A"/>
        </div>
        {/* Divider */}
        <div className="h-full w-[3px] bg-gray-500 rounded-lg"></div>
        {/* Second ChatRoom */}
        <div className="w-1/2">
          <ChatRoom user="B" />
        </div>
      </div>
    </div>
  );
}

export default App;