import React, { useState } from "react";

interface Props {
  onSend: (text: string) => void;
}

function MessageInput({ onSend }: Props) {
  const [message, setMessage] = useState<string>("");

  const handleSend = () => {
    if (message.trim()) {
      onSend(message);
      setMessage("");
    }
  };

  return (
    <div className="flex items-center bg-gray-200 p-4">
      <input
        type="text"
        placeholder="Type a message..."
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        className="flex-1 border border-gray-300 rounded-lg p-2"
      />
      <button
        onClick={handleSend}
        className="bg-teal-600 text-white px-4 py-2 rounded-lg ml-2"
      >
        Send
      </button>
    </div>
  );
}

export default MessageInput;