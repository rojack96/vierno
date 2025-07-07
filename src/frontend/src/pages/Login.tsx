import { useState } from "react";
import { useNavigate } from "react-router-dom";

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();

        const credentials = btoa(`${username}:${password}`);

        const res = await fetch("/api/status", {
            headers: {
                "Authorization": `Basic ${credentials}`,
            },
        });

        if (res.status === 200) {
            sessionStorage.setItem("basicAuth", credentials);
            navigate("/dashboard");
        } else if (res.status === 401) {
            setError("Credenziali non valide");
        } else {
            setError("Errore di connessione");
        }
    };

    return (
        <div style={{ maxWidth: 400, margin: "auto", padding: 20 }}>
            <h2>Login</h2>
            <form onSubmit={handleLogin}>
                <div>
                    <label>Username</label>
                    <input value={username} onChange={e => setUsername(e.target.value)} required />
                </div>
                <div>
                    <label>Password</label>
                    <input type="password" value={password} onChange={e => setPassword(e.target.value)} required />
                </div>
                <button type="submit">Login</button>
                {error && <p style={{ color: "red" }}>{error}</p>}
            </form>
        </div>
    );
};

export default Login;
