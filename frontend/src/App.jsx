import React, { useState } from "react";
import { Container, TextField, Button, Typography, Card, CardContent, Alert } from "@mui/material";

const API_URL = import.meta.env.VITE_API_URL || "http://backend:8080";

function App() {
  const [value, setValue] = useState("");
  const [response, setResponse] = useState(null);
  const [error, setError] = useState(null);

  const handleSearch = async () => {
    if (!value) {
      setError("Please enter a valid number.");
      return;
    }

    setResponse(null);
    setError(null);

    try {
      const res = await fetch(`${API_URL}/search/${value}`);
      const data = await res.json();

      if (!res.ok) {
        throw new Error(data.message || `Error ${res.status}: ${res.statusText}`);
      }

      setResponse(data);
    } catch (err) {
      console.error("Fetch error:", err.message);
      setError(err.message);
    }
  };

  return (
    <Container maxWidth="sm" style={{ marginTop: "50px", textAlign: "center" }}>
      <Typography variant="h4" gutterBottom>
        Search Closest Value
      </Typography>

      <TextField
        type="number"
        label="Enter a number"
        variant="outlined"
        fullWidth
        value={value}
        onChange={(e) => setValue(e.target.value)}
        style={{ marginBottom: "20px" }}
      />

      <Button variant="contained" color="primary" onClick={handleSearch} style={{ marginBottom: "20px" }}>
        Search
      </Button>

      {error && <Alert severity="error">{error}</Alert>}

      {response && (
        <Card style={{ marginTop: "20px" }}>
          <CardContent>
            <Typography variant="h6">Result:</Typography>
            <Typography>
              <strong>Index:</strong> {response.index !== -1 ? response.index : "Not Found"}
            </Typography>
            <Typography>
              <strong>Value:</strong> {response.value !== -1 ? response.value : "Not Found"}
            </Typography>
            <Typography>
              <strong>Message:</strong> {response.message}
            </Typography>
          </CardContent>
        </Card>
      )}
    </Container>
  );
}

export default App;
