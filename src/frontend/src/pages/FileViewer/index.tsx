import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Editor from "@monaco-editor/react";

const FileViewer = () => {
    const { folder, file } = useParams();
    const [content, setContent] = useState<string>("");
    const [isEditing, setIsEditing] = useState<boolean>(false);
    const [originalContent, setOriginalContent] = useState<string>("");

    useEffect(() => {
        fetch(`http://localhost:4788/${folder}.${file}`, {
            method: "GET",
            credentials: "include",
        })
            .then(res => res.text())
            .then(text => {
                setContent(text);
                setOriginalContent(text);
            })
            .catch(err => console.error("Errore caricamento contenuto:", err));
    }, [folder, file]);

    const handleSave = async () => {
        const response = await fetch(`http://localhost:4788/${folder}.${file}`, {
            method: "PUT",
            headers: { "Content-Type": "text/plain" },
            body: content,
            credentials: "include",
        });

        if (response.ok) {
            setOriginalContent(content);
            setIsEditing(false);
        } else {
            alert("Errore durante il salvataggio");
        }
    };

    return (
        <div style={{ padding: "2rem" }}>
            <div style={{ display: "flex", justifyContent: "space-between", marginBottom: "0.5rem" }}>
                <h2>{folder}/{file}</h2>
                {isEditing ? (
                    <div>
                        <button onClick={handleSave}>💾 Salva</button>
                        <button onClick={() => {
                            setContent(originalContent);
                            setIsEditing(false);
                        }} style={{ marginLeft: "0.5rem" }}>❌ Annulla</button>
                    </div>
                ) : (
                    <button onClick={() => setIsEditing(true)}>✏️ Modifica</button>
                )}
            </div>

            <Editor
                height="70vh"
                language={(file?.split(".").pop() || "plaintext")}
                value={content}
                onChange={(value) => setContent(value || "")}
                onMount={(editor, monaco) => {
                    setTimeout(() => editor.layout(), 0); // forza layout al mount
                }}
                options={{
                    readOnly: !isEditing,
                    fontSize: 14,
                    minimap: { enabled: false },
                    theme: "vs-dark",
                }}
            />
        </div>
    );
};

export default FileViewer;
