import { useNavigate, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import Card from "../../components/Card";

const FilePage = () => {
    const { folder } = useParams<{ folder: string }>();
    const navigate = useNavigate();
    const [files, setFiles] = useState<string[]>([]);

    useEffect(() => {
        fetch(`http://localhost:4788/folders/${folder}`, {
            method: "GET",
            credentials: "include",
        })
            .then(res => res.json())
            .then(data => setFiles(data.files || []))
            .catch(err => console.error("Errore caricamento file:", err));
    }, [folder]);

    const handleFileClick = (fileName: string) => {
        navigate(`/folders/${folder}/${fileName}`);
    };

    return (
        <div style={{ padding: "2rem" }}>
            <h2>File nella cartella: {folder}</h2>
            <Card elements={files} onClick={handleFileClick} />
        </div>
    );
};

export default FilePage;
