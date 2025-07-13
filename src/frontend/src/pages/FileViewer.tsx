import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const FileViewer = () => {
    const { folder, file } = useParams();
    const [content, setContent] = useState<string>("");

    useEffect(() => {
        fetch(`http://localhost:4788/${folder}.${file}`, {
            method: "GET",
            credentials: "include",
        })
            .then(res => res.text())
            .then(setContent)
            .catch(err => console.error("Errore caricamento contenuto:", err));
    }, [folder, file]);

    console.log("content", content)

    return (
        <div style={{ padding: "2rem" }}>
            <h2>Contenuto di: {file}</h2>
            <pre style={{
                background: "#f4f4f4",
                padding: "1rem",
                borderRadius: "8px",
                whiteSpace: "pre-wrap",
                wordWrap: "break-word"
            }}>
                {content}
            </pre>
        </div>
    );
};

export default FileViewer;
