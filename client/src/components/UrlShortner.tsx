"use client";

import { ChangeEvent, FormEvent, useEffect, useRef, useState } from "react";
import { Button } from "./ui/button";
import { Input } from "./ui/input";

type KeyHandler = (event: KeyboardEvent) => void;

export default function UrlShortner() {
  const [input, setInput] = useState<string>("");
  const inputRef = useRef<HTMLInputElement>(null);
  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    setInput("");
  };

  // handle shortcut keyword.
  useEffect(() => {
    const KeyJobs: KeyHandler = (event) => {
      switch (event.key) {
        case "/":
          inputRef.current?.focus();
          event.preventDefault();
          //   console.log("/ button pressed");
          break;
        case "Esc":
          inputRef.current?.blur();
        //   console.log("Esc button pressed");
        default:
      }
    };

    window.addEventListener("keydown", KeyJobs);

    return () => {
      window.removeEventListener("keydown", KeyJobs);
    };
  }, []);

  return (
    <form onSubmit={handleSubmit} className="w-1/2 flex gap-2">
      <Input
        type="text"
        ref={inputRef}
        placeholder="enter url , press ' / '"
        value={input}
        onChange={(e: ChangeEvent<HTMLInputElement>) =>
          setInput(e.target.value)
        }
      />
      <Button type="submit" onClick={handleSubmit} className="cursor-pointer">
        Generate Short Url
      </Button>
    </form>
  );
}
