import { useEffect, useState } from "react";
import "./style.css";
import Card from "../../components/Card";

const Dashboard = () => {
    const [folders, setFolders] = useState<string[]>([]);

    useEffect(() => {
        fetch("http://localhost:4788/folders/lookup", {
            method: "GET",
            credentials: "include",
        })
            .then(res => {
                if (!res.ok) {
                    throw new Error(`Errore HTTP: ${res.status}`);
                }
                return res.json();
            })
            .then(data => setFolders(data.folders || []))
            .catch(err => console.error("Error open folders", err));
    }, []);


    return (
        <div style={{ padding: "2rem" }}>
            <h2>Cartelle disponibili</h2>
            <Card elements={folders} />
        </div>
    );
};

export default Dashboard;
