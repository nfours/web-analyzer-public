import { useState } from "react";

const dummyData = {
    htmlVersion: "HTML5",
    "pageTitle": "Lucytech",
    internalLinks: 18,
    externalLinks: 9,
    loginFormExists: true
} 

export default function AnalyzerForm({ onResult }) {

  const [url, setUrl] = useState("");

  const analyze =  (e) => {
    e.preventDefault();
    onResult(dummyData);
  };

  return (
    <form onSubmit={analyze} style={{ textAlign: "center", marginTop: "2rem" }}>
      <input
        type="text"
        placeholder="Enter webpage URL"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
        style={{
          padding: "8px 12px",
          width: "70%",
          borderRadius: "8px",
          border: "1px solid #7C3AED",
          backgroundColor: "#1F1F28",
          color: "#fff"
        }}
      />
      <button
        type="submit"
        style={{
          marginLeft: "1rem",
          backgroundColor: "#7C3AED",
          color: "#fff",
          border: "none",
          padding: "8px 16px",
          borderRadius: "8px",
          cursor: "pointer"
        }}
      >
        Analyze
      </button>
    </form>
  );
}
