import  { useState } from "react";
import AnalyzerForm from "./components/AnalyzerForm";
import ResultCard from "./components/ResultCard";

export default function App() {
  const [result, setResult] = useState(null);
  console.log(result);
  return (
    <div className="p-8 max-w-3xl mx-auto">
      <h1 style={{ color: "#7C3AED", textAlign: "center" }}>Web Analyzer</h1>
      <AnalyzerForm onResult={setResult} />
      {result && <ResultCard data={result} />}
    </div>
  );
}
