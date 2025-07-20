import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./style.css";

interface LoginProps {
    setIsAuth: (value: boolean) => void;
}

const Login = ({ setIsAuth }: LoginProps) => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();

        const res = await fetch("http://localhost:4788/auth", {
            method: "POST",
            headers: { "Content-Type": "application/x-www-form-urlencoded" },
            credentials: "include",
            body: new URLSearchParams({ username, password }),
        });

        if (res.status === 200) {
            setIsAuth(true); // ✅ autentica
            navigate("/dashboard");
        } else if (res.status === 401) {
            setError("Credenziali non valide");
        } else {
            setError("Errore di connessione");
        }
    };

    return (
        <div className="login-form">
            <h2>Vierno</h2>
            <form onSubmit={handleLogin}>
                <div>
                    <input value={username} placeholder="Username" onChange={e => setUsername(e.target.value)} required />
                </div>
                <div>
                    <input type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} required />
                </div>
                <button type="submit">Login</button>
                {error && <p style={{ color: "red" }}>{error}</p>}
            </form>
        </div>
    );
};

export default Login;
