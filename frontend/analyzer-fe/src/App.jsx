import  { useState } from "react";
import AnalyzerForm from "./components/AnalyzerForm";
import ResultCard from "./components/ResultCard";
import ProgressBar from "./components/ProgressBar";
import ErrorMessage from "./components/ErrorMessage";
import { analyzePage } from "./api/analyzeApi";

export default function App() {
  const [loading, setLoading] = useState(false);
  const [progress, setProgress] = useState(0);
  const [error, setError] = useState("");
  const [result, setResult] = useState(null);

  const handleAnalyze = async (url) => {
    setLoading(true);
    setError("");
    setResult(null);
    setProgress(0);

    try {
      realTimeProgress();
      const data = await analyzePage(url);
      setResult(data);
      setProgress(100);
    } catch (err) {
      setError("Request Failed. Please try again.");
    } 
  };

  const realTimeProgress = () => {
    let progress = 0;
    const interval = setInterval(() => {
      progress += 10;
      if (progress > 100) progress = 100;
      setProgress(progress);

    }, 300);

    setTimeout(() => clearInterval(interval), 3300);

  };


  console.log(result);
  return (
    <div className="max-w-3xl mx-auto p-4" >
      <div style={{ display: 'flex', justifyContent: "center", alignItems: "center" }}><h1 className="text-2xl font-bold mb-4" >Web Analyzer</h1></div>
      <AnalyzerForm onSubmitClick={handleAnalyze} AnalizingStatus={loading} />
      <ProgressBar progress={progress} />


      <div style={{justifyItems:"center", alignItems: "center", color:"red", width:"100%"}}><p style={{display:"flex"}}>{error && <ErrorMessage message={error} />}</p></div>
      <div>
        <div style={{justifyItems:"center", alignItems: "center"}}>
          {result && <ResultCard data={result} />}
        </div>
      </div>
    </div>
  );
}
