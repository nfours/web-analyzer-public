import React from "react";

export default function ErrorMessage({ message }) {
  return (
    <div className="text-red-600 font-semibold my-2">
      {message}
    </div>
  );
}
