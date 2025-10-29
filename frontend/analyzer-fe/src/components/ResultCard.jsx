export default function ResultCard({ data }) {
  return (
    <div style={{
      backgroundColor: "#1F1F28",
      padding: "20px",
      borderRadius: "12px",
      marginTop: "2rem"
    }}>
      <h3 style={{ color: "#A78BFA" }}>Analysis Result</h3>
      <p><b>HTML Version:</b> {data.htmlVersion}</p>
      <p><b>Title:</b> {data.pageTitle}</p>
      <p><b>Internal Links:</b> {data.internalLinks}</p>
      <p><b>External Links:</b> {data.externalLinks}</p>
      <p><b>Inaccessible Links</b> {data.inaccessibleLinks}</p>
      <p><b>h1:</b> {data.headings.h1 || 0}</p>
      <p><b>h2:</b> {data.headings.h2 || 0}</p>
      <p><b>h3:</b> {data.headings.h3 || 0}</p>
      <p><b>h4:</b> {data.headings.h4 || 0}</p>
      <p><b>Login Form:</b> {data.loginFormExists ? "Yes" : "No"}</p>
    </div>
  );
}
