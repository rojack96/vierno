import { useEffect, useState } from "react";
import "./Dashboard.css";

const Dashboard = () => {
    const [folders, setFolders] = useState<string[]>([]);

    useEffect(() => {
        fetch("http://localhost:4788/folders", {
            method: "GET",
            credentials: "include", // ✅ fondamentale per inviare il cookie di sessione
        })
            .then(res => {
                if (!res.ok) {
                    throw new Error(`Errore HTTP: ${res.status}`);
                }
                return res.json();
            })
            .then(data => setFolders(data.folders || []))
            .catch(err => console.error("Errore caricamento cartelle:", err));
    }, []);


    return (
        <main className="dashboard-container">
            <h1>Benvenuto nella Dashboard!</h1>
            <div className="card-grid">
                {folders.map((folder, idx) => (
                    <div key={idx} className="card">
                        <h3>{folder}</h3>
                    </div>
                ))}
            </div>
        </main>
    );
};

export default Dashboard;
