import React, { useState, useEffect } from "react";
import MessageInput from "./MessageInput";
import LanguageSelector from "./LanguageSelector";

type Message = {
  original_text: string;
  translated_text: string;
  lang: string;
  user: string;
};

interface Props {
  user: string;
}

function ChatRoom({ user }: Props) {
  const [messages, setMessages] = useState<Message[]>([]);
  const [language, setLanguage] = useState<string>("English");

  // Update language on the backend when changed
  useEffect(() => {
    const updateLanguage = async () => {
      try {
        await fetch(`${process.env.REACT_APP_CHAT_API}/set-language`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ user, lang: language }),
        });
      } catch (error) {
        console.error("Error updating language:", error);
      }
    };

    updateLanguage();
  }, [language, user]);

  // Poll messages from the server
  useEffect(() => {
    const fetchMessages = async () => {
      try {
        const response = await fetch(`${process.env.REACT_APP_CHAT_API}/messages`);
        if (response.ok) {
          const data: Message[] = await response.json();
          setMessages(data);
        } else {
          console.error("Failed to fetch messages");
        }
      } catch (error) {
        console.error("Error fetching messages:", error);
      }
    };

    fetchMessages();
    const interval = setInterval(fetchMessages, 5000);

    return () => clearInterval(interval);
  }, []);

  // Send a message to the server
  const sendMessage = async (text: string) => {
    const message = { original_text: text, lang: language, user };
    try {
      const response = await fetch(`${process.env.REACT_APP_CHAT_API}/messages`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(message),
      });

      if (!response.ok) {
        console.error("Failed to send message");
      }
    } catch (error) {
      console.error("Error sending message:", error);
    }
  };

  return (
    <div className="flex flex-col h-[90vh] bg-gray-100">
      <header className="bg-teal-600 text-white p-4 shadow-md">
        <h1 className="text-xl font-bold">{`Chat Room for User ${user}`}</h1>
      </header>
      <LanguageSelector language={language} setLanguage={setLanguage} />
      <div className="flex-1 overflow-y-auto p-4 space-y-2">
        {(messages || []).map((msg, index) => {
          const isOwnMessage = msg.user === user;
          const displayedText = isOwnMessage ? msg.original_text : msg.translated_text;

          return (
            <div
              key={index}
              className={`flex ${
                isOwnMessage ? "justify-end" : "justify-start"
              }`}
            >
              <div
                className={`max-w-xs p-3 rounded-lg shadow-md text-white ${
                  isOwnMessage ? "bg-blue-500" : "bg-gray-500"
                }`}
              >
                <span className="block">{displayedText}</span>
              </div>
            </div>
          );
        })}
      </div>
      <MessageInput onSend={sendMessage} />
    </div>
  );
}

export default ChatRoom;